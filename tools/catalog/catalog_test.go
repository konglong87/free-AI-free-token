package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var testCaps = Capabilities{"文本生成": "📝", "图片理解": "👁", "文生图": "🖼"}

const validYAML = `
slug: demo
name: Demo
verified_at: "2026-01-02"
base_url: https://api.demo.test/v1
openai_compatible: supported
anthropic_endpoint: null
free: {type: unknown, note: null}
requirements: {credit_card: unknown, real_name: unknown, region: unknown, other: null}
sources:
  - {label: 官网, url: "https://demo.test/"}
home:
  order: 1
  pages: [readme, readme-b]
  tags: [文本生成, 图片理解]
  pending_tags: [图片理解]
  status: pending
  status_note: 待实测
  signup: {label: 注册, url: "https://demo.test/signup"}
  page_api_url: {readme-b: "https://api.demo.test/v1/"}
  extras:
    - {label: 定价, url: "https://demo.test/pricing"}
    - {label: 仅A, url: "https://demo.test/a", pages: [readme]}
models:
  - id: m1
    capabilities: [文本生成]
    home: {pages: [readme, readme-b]}
  - id: m2
    capabilities: [文本生成]
    home: {pages: [readme, readme-b], bold: true}
  - id: m3
    capabilities: [文本生成, 图片理解]
    pending_capabilities: [图片理解]
    home: {pages: [readme, readme-b], label: Demo V, suffix: 说明, group_note: 限时}
  - id: m4
    capabilities: [文生图]
    rpm: 10
    home: {pages: [readme]}
  - id: hidden
    capabilities: [文本生成]
`

func mustParse(t *testing.T, y string) Provider {
	t.Helper()
	p, err := ParseProvider([]byte(y))
	if err != nil {
		t.Fatal(err)
	}
	return p
}

func TestParseProvider(t *testing.T) {
	p := mustParse(t, validYAML)
	if p.Slug != "demo" || len(p.Models) != 5 || p.Models[3].RPM != 10 || !p.Models[1].Home.Bold {
		t.Fatalf("解析结果不对: %+v", p)
	}
	if p.AnthropicEndpoint != "" || p.Free.Note != "" {
		t.Fatalf("null 应解析为空值")
	}
	if _, err := ParseProvider([]byte(validYAML + "typo_field: 1\n")); err == nil {
		t.Fatal("未知字段应报错")
	}
}

func TestValidateOK(t *testing.T) {
	if errs := Validate(mustParse(t, validYAML), testCaps, t.TempDir()); len(errs) != 0 {
		t.Fatalf("合法数据不应报错: %v", errs)
	}
}

func TestValidateErrors(t *testing.T) {
	cases := []struct{ name, from, to, want string }{
		{"枚举", "openai_compatible: supported", "openai_compatible: maybe", "openai_compatible"},
		{"免费类型", "free: {type: unknown", "free: {type: forever", "free.type"},
		{"必填 base_url", "base_url: https://api.demo.test/v1", "base_url: null", "base_url 必填"},
		{"日期格式", `verified_at: "2026-01-02"`, `verified_at: "2026/01/02"`, "YYYY-MM-DD"},
		{"非统一标签", "tags: [文本生成, 图片理解]", "tags: [文本生成, 图片理解, 语音合成]", "非统一标签"},
		{"模型能力非统一标签", "capabilities: [文生图]", "capabilities: [绘画]", "非统一标签"},
		{"待核验必须是子集", "pending_tags: [图片理解]", "pending_tags: [文生图]", "pending_tags"},
		{"模型 ID 重复", "id: hidden", "id: m1", "重复"},
		{"状态枚举", "status: pending", "status: done", "home.status"},
		{"未知首页", "pages: [readme, readme-b]\n  tags", "pages: [readme, readme-c]\n  tags", "未知首页"},
		{"额外链接不是链接", `url: "https://demo.test/pricing"`, "url: ftp://x", "http(s)"},
		{"额外链接文件不存在", `url: "https://demo.test/pricing"`, "url: providers/demo/nope.png", "不存在"},
		{"首页没有模型", "home: {pages: [readme, readme-b]}\n  - id: m2", "home: {pages: [readme]}\n  - id: m2", ""},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			y := strings.Replace(validYAML, c.from, c.to, 1)
			if y == validYAML {
				t.Fatal("测试用例没有改动数据")
			}
			errs := Validate(mustParse(t, y), testCaps, t.TempDir())
			if c.want == "" {
				return // 仅确保不 panic
			}
			if !strings.Contains(strings.Join(errs, "\n"), c.want) {
				t.Fatalf("期望错误包含 %q，实际 %v", c.want, errs)
			}
		})
	}
}

func TestValidatePageWithoutModels(t *testing.T) {
	y := strings.ReplaceAll(validYAML, "home: {pages: [readme, readme-b]", "home: {pages: [readme]")
	y = strings.Replace(y, "home: {pages: [readme, readme-b], bold: true}", "home: {pages: [readme]}", 1)
	errs := Validate(mustParse(t, y), testCaps, t.TempDir())
	if !strings.Contains(strings.Join(errs, "\n"), `首页 "readme-b" 上没有任何模型`) {
		t.Fatalf("应提示 readme-b 无模型: %v", errs)
	}
}

func TestValidateTemplateIsLenient(t *testing.T) {
	p := Provider{Slug: templateSlug, Home: Home{Pages: []string{"readme"}, Status: "pending"}, OpenAICompatible: "unverified"}
	if errs := Validate(p, testCaps, t.TempDir()); len(errs) != 0 {
		t.Fatalf("模板不应要求必填: %v", errs)
	}
	p.Home.Tags = []string{"不存在"}
	if errs := Validate(p, testCaps, t.TempDir()); len(errs) == 0 {
		t.Fatal("模板里填了非法标签也应报错")
	}
}

func TestLoadProvidersChecksDirAndSlug(t *testing.T) {
	root := t.TempDir()
	write := func(rel, content string) {
		path := filepath.Join(root, rel)
		os.MkdirAll(filepath.Dir(path), 0o755)
		if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	write("docs/model-capabilities.md", "## 统一标签\n\n| 标签 | 输入 |\n| --- | --- |\n| 📝 文本生成 | a |\n| 👁 图片理解 | b |\n| 🖼 文生图 | c |\n")
	write("providers/demo/provider.yaml", validYAML)
	write("providers/other/provider.yaml", validYAML) // slug 仍是 demo，与目录名不一致
	write("providers/empty/README.md", "x")           // 缺少 provider.yaml
	_, errs := LoadProviders(root)
	msg := ""
	for _, e := range errs {
		msg += e.Error() + "\n"
	}
	for _, want := range []string{"other: slug", "empty: 缺少 provider.yaml"} {
		if !strings.Contains(msg, want) {
			t.Errorf("缺少错误 %q，实际:\n%s", want, msg)
		}
	}
}

func TestValidateAllDuplicateOrder(t *testing.T) {
	a, b := mustParse(t, validYAML), mustParse(t, validYAML)
	b.Slug = "demo2"
	if errs := ValidateAll([]Provider{a, b}, testCaps, t.TempDir()); len(errs) == 0 {
		t.Fatal("home.order 重复应报错")
	}
}

func TestParseCapabilities(t *testing.T) {
	md := "## 统一标签\n\n| 标签 | 输入 → 输出 |\n| --- | --- |\n| 📝 文本生成 | 文本 → 文本 |\n| 🖼 图像编辑 | x |\n\n## 其他\n\n| 🎬 不该解析 | x |\n"
	caps, err := ParseCapabilities(md)
	if err != nil {
		t.Fatal(err)
	}
	if len(caps) != 2 || caps["文本生成"] != "📝" || caps["图像编辑"] != "🖼" || caps.Has("不该解析") {
		t.Fatalf("标签解析不对: %v", caps)
	}
	if _, err := ParseCapabilities("# 空"); err == nil {
		t.Fatal("没有标签应报错")
	}
}

const wantRowA = `<tr>
<td valign="top"><a href="providers/demo/"><strong>Demo</strong></a><br><sub>📝 文本生成<br>👁 图片理解<br>（待核验）</sub></td>
<td valign="top"><strong>文本生成</strong><br><code>m1</code><br><strong><code>m2</code></strong><br><br><strong>文本生成 · 图片理解 · 限时（待核验）</strong><br><strong>Demo V</strong><br><code>m3</code> 说明<br><br><strong>文生图</strong><br><code>m4</code><br><br><a href="https://api.demo.test/v1">API</a> · <a href="https://demo.test/signup">注册</a> · <a href="providers/demo/models.md">完整模型</a><br><sub>状态：🟡 待实测 · <a href="https://demo.test/pricing">定价</a> · <a href="https://demo.test/a">仅A</a></sub></td>
</tr>`

func TestRenderRowPerPage(t *testing.T) {
	p := mustParse(t, validYAML)
	if got := renderRow(p, testCaps, "readme"); got != wantRowA {
		t.Fatalf("readme 行不一致:\n%s", got)
	}
	b := renderRow(p, testCaps, "readme-b")
	for _, want := range []string{`href="https://api.demo.test/v1/">API`, "定价"} {
		if !strings.Contains(b, want) {
			t.Errorf("readme-b 缺少 %q", want)
		}
	}
	for _, bad := range []string{"仅A", "<code>m4</code>", "hidden"} {
		if strings.Contains(b, bad) {
			t.Errorf("readme-b 不应包含 %q", bad)
		}
	}
}

func TestRenderTableOrderAndTemplate(t *testing.T) {
	a := mustParse(t, validYAML)
	b := mustParse(t, validYAML)
	b.Slug, b.Name, b.Home.Order = "second", "Second", 0
	b.Home.Order = 2
	tpl := Provider{Slug: templateSlug, Home: Home{Order: 999, Pages: []string{"readme", "readme-b"}}}
	all := []Provider{tpl, b, a}
	sortByOrder(all)

	readme := RenderTable(all, testCaps, "readme")
	if !(strings.Index(readme, "providers/demo/") < strings.Index(readme, "providers/second/") &&
		strings.Index(readme, "providers/second/") < strings.Index(readme, "providers/_template/")) {
		t.Fatal("README 的行应按 home.order 排序且模板在最后")
	}
	if strings.Contains(RenderTable(all, testCaps, "readme-b"), "_template") {
		t.Fatal("模板行只应出现在 README")
	}
	if _, n := ProvidersOn(all, "readme"); n != 2 {
		t.Fatalf("数量不应包含模板，得到 %d", n)
	}
}

func TestSplice(t *testing.T) {
	doc := "head\n<!-- providers-count:start --><img src=\"x/badge/providers-3-blue\"><!-- providers-count:end -->\n" +
		"<td><strong><!-- providers-count:start -->3<!-- providers-count:end --></strong></td>\n" +
		"<!-- providers:start -->\nOLD\n<!-- providers:end -->\ntail\n"
	got, err := Splice(doc, "NEW", 7)
	if err != nil {
		t.Fatal(err)
	}
	want := "head\n<!-- providers-count:start --><img src=\"x/badge/providers-7-blue\"><!-- providers-count:end -->\n" +
		"<td><strong><!-- providers-count:start -->7<!-- providers-count:end --></strong></td>\n" +
		"<!-- providers:start -->\nNEW\n<!-- providers:end -->\ntail\n"
	if got != want {
		t.Fatalf("Splice 结果不对:\n%s", got)
	}
	again, _ := Splice(got, "NEW", 7)
	if again != got {
		t.Fatal("Splice 应幂等")
	}
	for _, bad := range []string{"no markers", "<!-- providers:end -->\n<!-- providers:start -->", "<!-- providers:start -->\n<!-- providers:start -->\n<!-- providers:end -->"} {
		if _, err := Splice(bad, "x", 1); err == nil {
			t.Errorf("标记不合法应报错: %q", bad)
		}
	}
}

// 仓库里提交的首页必须与 provider.yaml 的生成结果一致（等价于 -check）。
func TestRepoInSync(t *testing.T) {
	root, err := filepath.Abs(filepath.Join("..", ".."))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(root, "docs", "model-capabilities.md")); err != nil {
		t.Skip("不在仓库内")
	}
	if code := run(root, true); code != 0 {
		t.Fatal("catalog -check 未通过")
	}
}
