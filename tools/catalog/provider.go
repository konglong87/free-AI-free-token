package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

const templateSlug = "_template"

// 首页标识 → 文件名。
var pageFiles = map[string]string{
	"readme":   "README.md",
	"readme-b": "readme-B.md",
}

var (
	openAICompat = []string{"supported", "unverified", "unsupported"}
	freeTypes    = []string{"permanent", "daily", "one_time", "promo", "unknown"}
	reqValues    = []string{"required", "none", "unknown"}
	statuses     = []string{"ok", "pending"}
	slugRe       = regexp.MustCompile(`^[a-z0-9]+(-[a-z0-9]+)*$`)
)

type Link struct {
	Label string   `yaml:"label"`
	URL   string   `yaml:"url"`
	Pages []string `yaml:"pages"` // 省略 = 供应商出现的所有首页
}

type Free struct {
	Type string `yaml:"type"`
	Note string `yaml:"note"`
}

type Requirements struct {
	CreditCard string `yaml:"credit_card"`
	RealName   string `yaml:"real_name"`
	Region     string `yaml:"region"`
	Other      string `yaml:"other"`
}

type Home struct {
	Order       int               `yaml:"order"`
	Pages       []string          `yaml:"pages"`
	Tags        []string          `yaml:"tags"`
	PendingTags []string          `yaml:"pending_tags"`
	Status      string            `yaml:"status"`
	StatusNote  string            `yaml:"status_note"`
	Signup      Link              `yaml:"signup"`
	PageAPIURL  map[string]string `yaml:"page_api_url"` // 个别首页的 API 链接与 base_url 不同时使用
	Extras      []Link            `yaml:"extras"`
}

type ModelHome struct {
	Pages     []string `yaml:"pages"`
	GroupNote string   `yaml:"group_note"` // 追加在分组标题后
	Label     string   `yaml:"label"`      // 模型 ID 上方的加粗名称
	Bold      bool     `yaml:"bold"`       // 模型 ID 加粗
	Suffix    string   `yaml:"suffix"`     // 模型 ID 后的说明文字
}

type Model struct {
	ID                  string    `yaml:"id"`
	Capabilities        []string  `yaml:"capabilities"`
	PendingCapabilities []string  `yaml:"pending_capabilities"`
	Context             string    `yaml:"context"`
	RPM                 int       `yaml:"rpm"`
	RPD                 int       `yaml:"rpd"`
	TPM                 int       `yaml:"tpm"`
	TPD                 int       `yaml:"tpd"`
	Note                string    `yaml:"note"`
	Home                ModelHome `yaml:"home"`
}

type Provider struct {
	Slug              string       `yaml:"slug"`
	Name              string       `yaml:"name"`
	VerifiedAt        string       `yaml:"verified_at"`
	BaseURL           string       `yaml:"base_url"`
	OpenAICompatible  string       `yaml:"openai_compatible"`
	AnthropicEndpoint string       `yaml:"anthropic_endpoint"`
	Free              Free         `yaml:"free"`
	Requirements      Requirements `yaml:"requirements"`
	Sources           []Link       `yaml:"sources"`
	Home              Home         `yaml:"home"`
	Models            []Model      `yaml:"models"`
}

func (p Provider) onPage(page string) bool { return contains(p.Home.Pages, page) }

func (m Model) onPage(page string) bool { return contains(m.Home.Pages, page) }

// ParseProvider 解析单个 provider.yaml；未知字段视为错误，避免拼写错误被悄悄忽略。
func ParseProvider(data []byte) (Provider, error) {
	var p Provider
	dec := yaml.NewDecoder(bytes.NewReader(data))
	dec.KnownFields(true)
	if err := dec.Decode(&p); err != nil {
		return p, err
	}
	return p, nil
}

// LoadProviders 读取 providers/*/provider.yaml，并校验目录名与 slug 一致。
func LoadProviders(root string) ([]Provider, []error) {
	dirs, err := filepath.Glob(filepath.Join(root, "providers", "*"))
	if err != nil {
		return nil, []error{err}
	}
	caps, err := LoadCapabilities(root)
	if err != nil {
		return nil, []error{err}
	}
	var providers []Provider
	var errs []error
	for _, dir := range dirs {
		if fi, err := os.Stat(dir); err != nil || !fi.IsDir() {
			continue
		}
		dirName := filepath.Base(dir)
		data, err := os.ReadFile(filepath.Join(dir, "provider.yaml"))
		if err != nil {
			errs = append(errs, fmt.Errorf("%s: 缺少 provider.yaml", dirName))
			continue
		}
		p, err := ParseProvider(data)
		if err != nil {
			errs = append(errs, fmt.Errorf("%s: 解析失败: %v", dirName, err))
			continue
		}
		if p.Slug != dirName {
			errs = append(errs, fmt.Errorf("%s: slug %q 与目录名不一致", dirName, p.Slug))
			continue
		}
		providers = append(providers, p)
	}
	errs = append(errs, ValidateAll(providers, caps, root)...)
	sortByOrder(providers)
	return providers, errs
}

func sortByOrder(ps []Provider) {
	sort.Slice(ps, func(i, j int) bool { return ps[i].Home.Order < ps[j].Home.Order })
}

// ValidateAll 校验每家供应商，并检查首页排序不重复。
func ValidateAll(providers []Provider, caps Capabilities, root string) []error {
	var errs []error
	orders := map[int]string{}
	for _, p := range providers {
		for _, msg := range Validate(p, caps, root) {
			errs = append(errs, fmt.Errorf("%s: %s", p.Slug, msg))
		}
		if other, ok := orders[p.Home.Order]; ok {
			errs = append(errs, fmt.Errorf("%s: home.order %d 与 %s 重复", p.Slug, p.Home.Order, other))
		}
		orders[p.Home.Order] = p.Slug
	}
	return errs
}

// Validate 返回单个供应商的所有校验问题。_template 只检查取值合法性，不要求必填。
func Validate(p Provider, caps Capabilities, root string) []string {
	var errs []string
	add := func(format string, a ...any) { errs = append(errs, fmt.Sprintf(format, a...)) }
	tpl := p.Slug == templateSlug

	if !tpl && !slugRe.MatchString(p.Slug) {
		add("slug %q 必须是小写 kebab-case", p.Slug)
	}
	required := func(field, value string) {
		if !tpl && strings.TrimSpace(value) == "" {
			add("%s 必填", field)
		}
	}
	oneOf := func(field, value string, allowed []string) {
		if value == "" && tpl {
			return
		}
		if !contains(allowed, value) {
			add("%s=%q 不在枚举 %v 内", field, value, allowed)
		}
	}
	checkURL := func(field, value string) {
		if value == "" || isHTTP(value) {
			return
		}
		if strings.Contains(value, "://") {
			add("%s=%q 必须是 http(s) 链接或仓库内相对路径", field, value)
			return
		}
		if _, err := os.Stat(filepath.Join(root, filepath.FromSlash(value))); err != nil {
			add("%s=%q 指向的仓库文件不存在", field, value)
		}
	}
	checkPages := func(field string, pages, within []string) {
		for _, pg := range pages {
			if _, ok := pageFiles[pg]; !ok {
				add("%s 含未知首页 %q", field, pg)
			} else if within != nil && !contains(within, pg) {
				add("%s 的首页 %q 不在所属供应商 home.pages 内", field, pg)
			}
		}
	}

	required("name", p.Name)
	if p.VerifiedAt != "" {
		if _, err := time.Parse("2006-01-02", p.VerifiedAt); err != nil {
			add("verified_at=%q 必须是 YYYY-MM-DD", p.VerifiedAt)
		}
	} else {
		required("verified_at", p.VerifiedAt)
	}
	required("base_url", p.BaseURL)
	if p.BaseURL != "" && !isHTTP(p.BaseURL) {
		add("base_url=%q 必须是 http(s) 链接", p.BaseURL)
	}
	oneOf("openai_compatible", p.OpenAICompatible, openAICompat)
	if p.AnthropicEndpoint != "" && !isHTTP(p.AnthropicEndpoint) {
		add("anthropic_endpoint=%q 必须是 http(s) 链接", p.AnthropicEndpoint)
	}
	oneOf("free.type", p.Free.Type, freeTypes)
	oneOf("requirements.credit_card", p.Requirements.CreditCard, reqValues)
	oneOf("requirements.real_name", p.Requirements.RealName, reqValues)
	oneOf("requirements.region", p.Requirements.Region, reqValues)

	if !tpl && len(p.Sources) == 0 {
		add("sources 至少一条官方来源")
	}
	for i, s := range p.Sources {
		if s.Label == "" || !isHTTP(s.URL) {
			add("sources[%d] 需要 label 和 http(s) url", i)
		}
	}

	h := p.Home
	if !tpl && h.Order <= 0 {
		add("home.order 必须为正整数")
	}
	if !tpl && len(h.Pages) == 0 {
		add("home.pages 必填")
	}
	checkPages("home.pages", h.Pages, nil)
	if !tpl && len(h.Tags) == 0 {
		add("home.tags 至少一个能力标签")
	}
	checkTags := func(field string, tags []string) {
		seen := map[string]bool{}
		for _, t := range tags {
			if !caps.Has(t) {
				add("%s 含非统一标签 %q（见 docs/model-capabilities.md）", field, t)
			}
			if seen[t] {
				add("%s 中 %q 重复", field, t)
			}
			seen[t] = true
		}
	}
	checkTags("home.tags", h.Tags)
	for _, t := range h.PendingTags {
		if !contains(h.Tags, t) {
			add("home.pending_tags 的 %q 不在 home.tags 内", t)
		}
	}
	oneOf("home.status", h.Status, statuses)
	required("home.signup.label", h.Signup.Label)
	if !tpl && !isHTTP(h.Signup.URL) {
		add("home.signup.url 必须是 http(s) 链接")
	}
	for pg, u := range h.PageAPIURL {
		if !contains(h.Pages, pg) {
			add("home.page_api_url 的首页 %q 不在 home.pages 内", pg)
		}
		if !isHTTP(u) {
			add("home.page_api_url[%s]=%q 必须是 http(s) 链接", pg, u)
		}
	}
	for i, e := range h.Extras {
		if e.Label == "" || e.URL == "" {
			add("home.extras[%d] 需要 label 和 url", i)
			continue
		}
		checkURL(fmt.Sprintf("home.extras[%d].url", i), e.URL)
		checkPages(fmt.Sprintf("home.extras[%d].pages", i), e.Pages, h.Pages)
	}

	if !tpl && len(p.Models) == 0 {
		add("models 至少一个模型")
	}
	ids := map[string]bool{}
	for i, m := range p.Models {
		f := fmt.Sprintf("models[%d]", i)
		if m.ID == "" {
			add("%s.id 必填", f)
		} else if ids[m.ID] {
			add("模型 ID %q 重复", m.ID)
		}
		ids[m.ID] = true
		if len(m.Capabilities) == 0 {
			add("%s(%s).capabilities 至少一个能力标签", f, m.ID)
		}
		checkTags(f+".capabilities", m.Capabilities)
		for _, t := range m.PendingCapabilities {
			if !contains(m.Capabilities, t) {
				add("%s(%s).pending_capabilities 的 %q 不在 capabilities 内", f, m.ID, t)
			}
		}
		if m.RPM < 0 || m.RPD < 0 || m.TPM < 0 || m.TPD < 0 {
			add("%s(%s) 的 rpm/rpd/tpm/tpd 不能为负数", f, m.ID)
		}
		checkPages(f+".home.pages", m.Home.Pages, h.Pages)
	}
	if !tpl {
		for _, pg := range h.Pages {
			onPage := false
			for _, m := range p.Models {
				onPage = onPage || m.onPage(pg)
			}
			if !onPage {
				add("首页 %q 上没有任何模型（models[].home.pages）", pg)
			}
		}
	}
	return errs
}

func contains(list []string, v string) bool {
	for _, x := range list {
		if x == v {
			return true
		}
	}
	return false
}

func isHTTP(s string) bool {
	return strings.HasPrefix(s, "https://") || strings.HasPrefix(s, "http://")
}
