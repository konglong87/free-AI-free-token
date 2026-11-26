package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"
)

const fakeKey = "sk-test-FAKEKEY-1234567890"

// smokeRoot 建一个只含 demo 供应商的临时仓库，base_url 指向测试服务器。
func smokeRoot(t *testing.T, baseURL string) string {
	t.Helper()
	root := t.TempDir()
	must := func(err error) {
		t.Helper()
		if err != nil {
			t.Fatal(err)
		}
	}
	must(os.MkdirAll(filepath.Join(root, "docs"), 0o755))
	must(os.MkdirAll(filepath.Join(root, "providers", "demo"), 0o755))
	must(os.WriteFile(filepath.Join(root, "docs", "model-capabilities.md"),
		[]byte("## 统一标签\n\n| 标签 | 说明 |\n|---|---|\n| 📝 文本生成 | a |\n| 👁 图片理解 | b |\n| 🖼 文生图 | c |\n"), 0o644))
	y := strings.Replace(validYAML, "https://api.demo.test/v1\n", baseURL+"/v1/\n", 1)
	must(os.WriteFile(filepath.Join(root, "providers", "demo", "provider.yaml"), []byte(y), 0o644))
	return root
}

func fixedNow() time.Time { return time.Date(2026, 11, 26, 9, 0, 0, 0, time.UTC) }

func smokeOptions(env map[string]string) smokeOpts {
	return smokeOpts{Timeout: 2 * time.Second, Getenv: func(k string) string { return env[k] }, Now: fixedNow}
}

func verifiedAt(t *testing.T, root string) string {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(root, "providers", "demo", "provider.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	return mustParse(t, string(data)).VerifiedAt
}

func TestSmokeEnvVar(t *testing.T) {
	if got := SmokeEnvVar("nvidia-nim"); got != "NVIDIA_NIM_API_KEY" {
		t.Fatal(got)
	}
	if got := SmokeEnvVar("groq"); got != "GROQ_API_KEY" {
		t.Fatal(got)
	}
}

func TestSmokeModelSelection(t *testing.T) {
	p := mustParse(t, validYAML)
	if id, ok := SmokeModel(p); !ok || id != "m1" {
		t.Fatalf("默认应取第一个上首页的文本模型，得到 %q", id)
	}
	p.Models[1].Smoke = true
	if id, _ := SmokeModel(p); id != "m2" {
		t.Fatalf("smoke: true 应优先，得到 %q", id)
	}
	if _, ok := SmokeModel(Provider{}); ok {
		t.Fatal("没有模型时不应选出模型")
	}
}

func TestValidateSmoke(t *testing.T) {
	p := mustParse(t, validYAML)
	p.Models[0].Smoke, p.Models[1].Smoke = true, true
	if errs := Validate(p, testCaps, t.TempDir()); len(errs) != 1 || !strings.Contains(errs[0], "最多一个") {
		t.Fatalf("两个 smoke 应报错: %v", errs)
	}
	p = mustParse(t, validYAML)
	p.Models[3].Smoke = true // 文生图
	if errs := Validate(p, testCaps, t.TempDir()); len(errs) != 1 || !strings.Contains(errs[0], "smoke") {
		t.Fatalf("非文本模型的 smoke 应报错: %v", errs)
	}
}

func TestSmokeSuccess(t *testing.T) {
	var chatBody map[string]any
	var auth []string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth = append(auth, r.Header.Get("Authorization"))
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/v1/models":
			io.WriteString(w, `{"data":[]}`)
		case r.Method == http.MethodPost && r.URL.Path == "/v1/chat/completions":
			json.NewDecoder(r.Body).Decode(&chatBody)
			io.WriteString(w, `{"choices":[]}`)
		default:
			http.NotFound(w, r)
		}
	}))
	defer srv.Close()

	root := smokeRoot(t, srv.URL)
	var out bytes.Buffer
	if code := runSmoke(root, smokeOptions(map[string]string{"DEMO_API_KEY": fakeKey}), &out); code != 0 {
		t.Fatalf("退出码 %d\n%s", code, out.String())
	}
	if chatBody["model"] != "m1" || chatBody["max_tokens"] != float64(1) {
		t.Fatalf("chat 请求体不对: %v", chatBody)
	}
	if len(auth) != 2 || auth[0] != "Bearer "+fakeKey {
		t.Fatalf("Authorization 不对: %v", auth)
	}
	for _, want := range []string{"| Demo | 200 | 200 |", "✅ 通过", "可考虑把 home.status 改为 ok"} {
		if !strings.Contains(out.String(), want) {
			t.Fatalf("输出缺少 %q:\n%s", want, out.String())
		}
	}
	if strings.Contains(out.String(), fakeKey) {
		t.Fatal("输出里出现了 key")
	}
	if got := verifiedAt(t, root); got != "2026-01-02" {
		t.Fatalf("没有 -write 时不应改 verified_at: %s", got)
	}
}

func TestSmokeWriteOnlyWhenBothOK(t *testing.T) {
	chatStatus := http.StatusOK
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/chat/completions") {
			w.WriteHeader(chatStatus)
		}
	}))
	defer srv.Close()
	root := smokeRoot(t, srv.URL)
	o := smokeOptions(map[string]string{"DEMO_API_KEY": fakeKey})
	o.Write = true

	chatStatus = http.StatusInternalServerError
	if code := runSmoke(root, o, io.Discard); code != 1 {
		t.Fatalf("chat 失败应返回 1，得到 %d", code)
	}
	if got := verifiedAt(t, root); got != "2026-01-02" {
		t.Fatalf("失败时不应改 verified_at: %s", got)
	}

	chatStatus = http.StatusOK
	var out bytes.Buffer
	if code := runSmoke(root, o, &out); code != 0 {
		t.Fatalf("退出码 %d\n%s", code, out.String())
	}
	if got := verifiedAt(t, root); got != "2026-11-26" {
		t.Fatalf("两项成功后应改成今天: %s", got)
	}
	data, _ := os.ReadFile(filepath.Join(root, "providers", "demo", "provider.yaml"))
	if !strings.Contains(string(data), `verified_at: "2026-11-26"`) || !strings.Contains(string(data), "status: pending") {
		t.Fatalf("应只改 verified_at，home.status 保持不变:\n%s", data)
	}
}

func TestSmokeErrorStatuses(t *testing.T) {
	cases := []struct {
		name string
		code int
		body string
		want string
	}{
		{"401", 401, `{"error":{"message":"Incorrect API key provided: ` + fakeKey + `"}}`, "Incorrect API key"},
		{"429", 429, `{"error":"rate limit exceeded for ` + fakeKey + `"}`, "rate limit"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(c.code)
				io.WriteString(w, c.body)
			}))
			defer srv.Close()
			root := smokeRoot(t, srv.URL)
			o := smokeOptions(map[string]string{"DEMO_API_KEY": fakeKey})
			o.Write = true
			var out bytes.Buffer
			if code := runSmoke(root, o, &out); code != 1 {
				t.Fatalf("应返回 1，得到 %d", code)
			}
			s := out.String()
			if !strings.Contains(s, "| Demo | "+itoa(c.code)+" | "+itoa(c.code)+" |") || !strings.Contains(s, c.want) {
				t.Fatalf("输出不对:\n%s", s)
			}
			if strings.Contains(s, fakeKey) || !strings.Contains(s, "***") {
				t.Fatalf("key 必须被打码:\n%s", s)
			}
			if verifiedAt(t, root) != "2026-01-02" {
				t.Fatal("失败时不应改 verified_at")
			}
		})
	}
}

func itoa(n int) string { return strconv.Itoa(n) }

func TestSmokeTimeout(t *testing.T) {
	done := make(chan struct{})
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-done:
		case <-r.Context().Done():
		}
	}))
	defer srv.Close()
	defer close(done)

	root := smokeRoot(t, srv.URL)
	o := smokeOptions(map[string]string{"DEMO_API_KEY": fakeKey})
	o.Timeout = 100 * time.Millisecond
	var out bytes.Buffer
	if code := runSmoke(root, o, &out); code != 1 {
		t.Fatalf("超时应返回 1，得到 %d", code)
	}
	s := out.String()
	if !strings.Contains(s, "超时") || !strings.Contains(s, "| Demo | - | - |") {
		t.Fatalf("输出不对:\n%s", s)
	}
	if strings.Contains(s, fakeKey) {
		t.Fatal("输出里出现了 key")
	}
}

func TestSmokeNetworkErrorMasksKey(t *testing.T) {
	srv := httptest.NewServer(http.NotFoundHandler())
	url := srv.URL
	srv.Close() // 连接被拒绝
	root := smokeRoot(t, url)
	var out bytes.Buffer
	if code := runSmoke(root, smokeOptions(map[string]string{"DEMO_API_KEY": fakeKey}), &out); code != 1 {
		t.Fatalf("应返回 1，得到 %d", code)
	}
	if strings.Contains(out.String(), fakeKey) {
		t.Fatal("输出里出现了 key")
	}
}

func TestSmokeNoKeySkipsAll(t *testing.T) {
	hit := false
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { hit = true }))
	defer srv.Close()
	root := smokeRoot(t, srv.URL)
	var out bytes.Buffer
	if code := runSmoke(root, smokeOptions(nil), &out); code != 0 {
		t.Fatalf("全部跳过应返回 0，得到 %d", code)
	}
	if hit || !strings.Contains(out.String(), "全部跳过") {
		t.Fatalf("没有 key 时不应发请求:\n%s", out.String())
	}
}

func TestSmokeProviderFilter(t *testing.T) {
	root := smokeRoot(t, "http://127.0.0.1:1")
	if code := runSmoke(root, smokeOpts{Provider: "nope", Timeout: time.Second, Getenv: func(string) string { return "" }, Now: fixedNow}, io.Discard); code != 1 {
		t.Fatalf("不存在的供应商应返回 1，得到 %d", code)
	}
}

func TestMaskKey(t *testing.T) {
	got := maskKey("bad "+fakeKey+" and Authorization: Bearer abc.def", fakeKey)
	if strings.Contains(got, fakeKey) || strings.Contains(got, "abc.def") {
		t.Fatal(got)
	}
}

func TestSmokeModelSkipsWildcard(t *testing.T) {
	p := mustParse(t, validYAML)
	p.Models[0].ID = "*:free"
	if id, _ := SmokeModel(p); id != "m2" {
		t.Fatalf("应跳过 ID 含 * 的模型集合，得到 %q", id)
	}
	p.Models[1].ID, p.Models[1].Smoke = "x/*", true
	if id, _ := SmokeModel(p); id != "m3" {
		t.Fatalf("smoke: true 也不能选中含 * 的 ID，得到 %q", id)
	}
}

func TestSmokePlaceholderBaseURLSkipped(t *testing.T) {
	hit := false
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { hit = true }))
	defer srv.Close()
	root := smokeRoot(t, srv.URL+"/accounts/{account_id}")
	var out bytes.Buffer
	// 即使设置了 key 也要跳过，且不需要额外环境变量
	if code := runSmoke(root, smokeOptions(map[string]string{"DEMO_API_KEY": fakeKey}), &out); code != 0 {
		t.Fatalf("跳过应返回 0，得到 %d\n%s", code, out.String())
	}
	if hit {
		t.Fatal("base_url 含占位符时不应发请求")
	}
	if !strings.Contains(out.String(), "base_url 含占位符，需手动实测") || !strings.Contains(out.String(), "demo") {
		t.Fatalf("输出缺少占位符提示:\n%s", out.String())
	}
}
