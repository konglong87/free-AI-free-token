package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

// SmokeResult 是单个供应商的实测结果。Skipped 表示没有设置 key，未发请求。
type SmokeResult struct {
	Slug        string
	Name        string
	EnvVar      string
	Skipped     bool
	Placeholder bool // 跳过原因：base_url 含 {占位符}，需手动实测
	Model       string
	ModelsCode  int // 0 = 没有拿到 HTTP 响应（网络错误或超时）
	ChatCode    int
	Elapsed     time.Duration
	Err         string // 已打码的简短错误
	Status      string // home.status，用于「可考虑改为 ok」提示
}

func (r SmokeResult) OK() bool { return !r.Skipped && is2xx(r.ModelsCode) && is2xx(r.ChatCode) }

func is2xx(code int) bool { return code >= 200 && code < 300 }

// SmokeEnvVar 返回供应商 key 的环境变量名：slug 转大写、- 换成 _，再加 _API_KEY。
func SmokeEnvVar(slug string) string {
	return strings.ToUpper(strings.ReplaceAll(slug, "-", "_")) + "_API_KEY"
}

// SmokeModel 选择实测用的模型：优先 smoke: true，否则第一个上首页的文本模型。
// ID 含 * 的是模型集合（如 "*:free"），不是真实模型，一律跳过。
func SmokeModel(p Provider) (string, bool) {
	for _, m := range p.Models {
		if m.Smoke && !strings.Contains(m.ID, "*") {
			return m.ID, true
		}
	}
	for _, m := range p.Models {
		if len(m.Home.Pages) > 0 && contains(m.Capabilities, textCapability) && !strings.Contains(m.ID, "*") {
			return m.ID, true
		}
	}
	return "", false
}

const textCapability = "文本生成"

var bearerRe = regexp.MustCompile(`(?i)bearer\s+\S+`)

// maskKey 把错误文本里出现的 key 以及 Bearer 凭据替换成 ***。
func maskKey(s, key string) string {
	if key != "" {
		s = strings.ReplaceAll(s, key, "***")
	}
	return bearerRe.ReplaceAllString(s, "Bearer ***")
}

// errSnippet 从错误响应里取一段简短、单行、已打码的说明。
func errSnippet(body []byte, key string) string {
	var parsed struct {
		Error json.RawMessage `json:"error"`
	}
	text := strings.TrimSpace(string(body))
	if json.Unmarshal(body, &parsed) == nil && len(parsed.Error) > 0 {
		var obj struct {
			Message string `json:"message"`
		}
		var str string
		switch {
		case json.Unmarshal(parsed.Error, &obj) == nil && obj.Message != "":
			text = obj.Message
		case json.Unmarshal(parsed.Error, &str) == nil && str != "":
			text = str
		}
	}
	text = strings.Join(strings.Fields(text), " ")
	text = strings.ReplaceAll(text, "|", "/")
	text = maskKey(text, key)
	if r := []rune(text); len(r) > 80 {
		text = string(r[:80]) + "…"
	}
	return text
}

// doSmoke 发一次请求，返回状态码（0 = 无响应）和已打码的错误说明（成功为空）。
func doSmoke(client *http.Client, method, target, key string, payload []byte) (int, string) {
	var body io.Reader
	if payload != nil {
		body = bytes.NewReader(payload)
	}
	req, err := http.NewRequest(method, target, body)
	if err != nil {
		return 0, maskKey(err.Error(), key)
	}
	req.Header.Set("Authorization", "Bearer "+key)
	req.Header.Set("Accept", "application/json")
	if payload != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	resp, err := client.Do(req)
	if err != nil {
		var ne net.Error
		if errors.Is(err, context.DeadlineExceeded) || (errors.As(err, &ne) && ne.Timeout()) {
			return 0, fmt.Sprintf("超时(>%s)", client.Timeout)
		}
		msg := err.Error()
		var ue *url.Error
		if errors.As(err, &ue) {
			msg = ue.Err.Error() // 去掉 Get "<url>": 前缀
		}
		return 0, maskKey(msg, key)
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(io.LimitReader(resp.Body, 4096))
	if is2xx(resp.StatusCode) {
		return resp.StatusCode, ""
	}
	return resp.StatusCode, errSnippet(data, key)
}

// SmokeProvider 对一家供应商发两次请求：GET /models，再发一次 max_tokens=1 的 chat completion。
// 每个请求只发一次，不重试。
func SmokeProvider(client *http.Client, p Provider, key string) SmokeResult {
	res := SmokeResult{Slug: p.Slug, Name: p.Name, EnvVar: SmokeEnvVar(p.Slug), Status: p.Home.Status}
	model, ok := SmokeModel(p)
	if !ok {
		res.Err = "没有可用的文本模型（models[]）"
		return res
	}
	res.Model = model
	base := strings.TrimRight(p.BaseURL, "/")
	start := time.Now()

	var errs []string
	var msg string
	res.ModelsCode, msg = doSmoke(client, http.MethodGet, base+"/models", key, nil)
	if msg != "" {
		errs = append(errs, "models: "+msg)
	}
	payload, _ := json.Marshal(map[string]any{
		"model":      model,
		"messages":   []map[string]string{{"role": "user", "content": "hi"}},
		"max_tokens": 1,
	})
	res.ChatCode, msg = doSmoke(client, http.MethodPost, base+"/chat/completions", key, payload)
	if msg != "" {
		errs = append(errs, "chat: "+msg)
	}
	res.Elapsed = time.Since(start)
	res.Err = strings.Join(errs, "; ")
	return res
}

func codeCell(code int) string {
	if code == 0 {
		return "-"
	}
	return fmt.Sprint(code)
}

// RenderSmoke 输出 Markdown 结果表；未设置 key 的供应商只在表后列出。
func RenderSmoke(results []SmokeResult) string {
	var tested, skipped, manual []SmokeResult
	for _, r := range results {
		switch {
		case r.Placeholder:
			manual = append(manual, r)
		case r.Skipped:
			skipped = append(skipped, r)
		default:
			tested = append(tested, r)
		}
	}
	var b strings.Builder
	if len(tested) == 0 {
		b.WriteString("全部跳过：没有检测到任何供应商的 API key 环境变量，或需手动实测。\n")
	} else {
		b.WriteString("| 供应商 | /models | chat | 耗时 | 备注 |\n|---|---|---|---|---|\n")
		for _, r := range tested {
			note := r.Err
			if r.OK() {
				note = "✅ 通过"
			}
			fmt.Fprintf(&b, "| %s | %s | %s | %s | %s |\n", r.Name, codeCell(r.ModelsCode), codeCell(r.ChatCode),
				r.Elapsed.Round(time.Millisecond), note)
		}
	}
	if len(skipped) > 0 {
		names := make([]string, len(skipped))
		for i, r := range skipped {
			names[i] = r.Slug
		}
		fmt.Fprintf(&b, "\n跳过（未设置对应 API key）%d 家：%s\n", len(skipped), strings.Join(names, "、"))
	}
	if len(manual) > 0 {
		names := make([]string, len(manual))
		for i, r := range manual {
			names[i] = r.Slug
		}
		fmt.Fprintf(&b, "\n跳过（base_url 含占位符，需手动实测）%d 家：%s\n", len(manual), strings.Join(names, "、"))
	}
	return b.String()
}

// SetVerifiedAt 只替换 provider.yaml 里的 verified_at 一行，保留其余内容和注释。
func SetVerifiedAt(path, date string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	re := regexp.MustCompile(`(?m)^verified_at:.*$`)
	if !re.Match(data) {
		return fmt.Errorf("%s 中找不到 verified_at", path)
	}
	next := re.ReplaceAll(data, []byte(fmt.Sprintf("verified_at: %q", date)))
	return os.WriteFile(path, next, 0o644)
}

type smokeOpts struct {
	Provider string
	Write    bool
	Timeout  time.Duration
	Getenv   func(string) string
	Now      func() time.Time
}

// runSmoke 执行实测并把结果写到 out。任何一家失败返回 1；全部跳过返回 0。
func runSmoke(root string, o smokeOpts, out io.Writer) int {
	providers, errs := LoadProviders(root)
	if len(errs) > 0 {
		for _, e := range errs {
			fmt.Fprintln(os.Stderr, "错误:", e)
		}
		return 1
	}
	var selected []Provider
	for _, p := range providers {
		if p.Slug == templateSlug || (o.Provider != "" && p.Slug != o.Provider) {
			continue
		}
		selected = append(selected, p)
	}
	if o.Provider != "" && len(selected) == 0 {
		fmt.Fprintf(os.Stderr, "错误: 找不到供应商 %q\n", o.Provider)
		return 1
	}

	client := &http.Client{Timeout: o.Timeout}
	var results []SmokeResult
	for _, p := range selected {
		if strings.Contains(p.BaseURL, "{") {
			results = append(results, SmokeResult{Slug: p.Slug, Name: p.Name, EnvVar: SmokeEnvVar(p.Slug), Skipped: true, Placeholder: true})
			continue
		}
		key := o.Getenv(SmokeEnvVar(p.Slug))
		if key == "" {
			results = append(results, SmokeResult{Slug: p.Slug, Name: p.Name, EnvVar: SmokeEnvVar(p.Slug), Skipped: true})
			continue
		}
		results = append(results, SmokeProvider(client, p, key))
	}
	fmt.Fprint(out, RenderSmoke(results))

	today := o.Now().Format("2006-01-02")
	failed := false
	for _, r := range results {
		if r.Skipped {
			continue
		}
		if !r.OK() {
			failed = true
			continue
		}
		if r.Status != "ok" {
			fmt.Fprintf(out, "\n💡 %s：两项实测均成功，可考虑把 home.status 改为 ok（不会自动修改）\n", r.Slug)
		}
		if o.Write {
			if err := SetVerifiedAt(filepath.Join(root, "providers", r.Slug, "provider.yaml"), today); err != nil {
				fmt.Fprintln(os.Stderr, "错误:", err)
				return 1
			}
			fmt.Fprintf(out, "\n已更新 %s 的 verified_at: %s\n", r.Slug, today)
		}
	}
	if failed {
		return 1
	}
	return 0
}

func smokeCommand(args []string) int {
	fs := flag.NewFlagSet("smoke", flag.ExitOnError)
	root := fs.String("root", "", "仓库根目录（默认从当前目录向上查找）")
	provider := fs.String("provider", "", "只测这一家（slug）")
	write := fs.Bool("write", false, "两项均成功时，把该供应商的 verified_at 改成今天")
	timeout := fs.Duration("timeout", 20*time.Second, "单个请求的超时时间")
	fs.Parse(args)
	r, err := resolveRoot(*root)
	if err != nil {
		fatal(err)
	}
	return runSmoke(r, smokeOpts{Provider: *provider, Write: *write, Timeout: *timeout, Getenv: os.Getenv, Now: time.Now}, os.Stdout)
}
