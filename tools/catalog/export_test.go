package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
	"time"
)

var exportNow = time.Date(2026, 11, 26, 8, 0, 0, 0, time.UTC)

func TestExportProvidersOnlyAndRecommended(t *testing.T) {
	a := mustParse(t, validYAML)
	a.Models[1].Recommended, a.Models[1].ToolCalling = true, "supported"
	a.UsageNotes = []string{"免费层输入可能用于训练"}
	b := mustParse(t, validYAML)
	b.Slug, b.Home.Order = "second", 2
	tpl := Provider{Slug: templateSlug}

	cat, err := ExportProviders([]Provider{tpl, a, b}, []string{"second", "demo"}, exportNow)
	if err != nil {
		t.Fatal(err)
	}
	if cat.SchemaVersion != 1 || cat.GeneratedAt != "2026-11-26T08:00:00Z" {
		t.Fatalf("头部不对: %+v", cat)
	}
	if len(cat.Providers) != 2 || cat.Providers[0].Slug != "second" || cat.Providers[1].Slug != "demo" {
		t.Fatalf("应按 -only 顺序只导出两家，得到 %+v", cat.Providers)
	}
	demo := cat.Providers[1]
	if demo.RecommendedModel == nil || demo.RecommendedModel.ID != "m2" || demo.RecommendedModel.ToolCalling != "supported" {
		t.Fatalf("recommended_model 不对: %+v", demo.RecommendedModel)
	}
	if cat.Providers[0].RecommendedModel != nil {
		t.Fatal("没有 recommended 的供应商应导出 null")
	}
	if demo.Models[0].ToolCalling != "pending" {
		t.Fatalf("未填写 tool_calling 应导出 pending，得到 %q", demo.Models[0].ToolCalling)
	}
	if demo.SignupURL != "https://demo.test/signup" || demo.Status != "pending" || len(demo.UsageNotes) != 1 {
		t.Fatalf("字段映射不对: %+v", demo)
	}

	all, _ := ExportProviders([]Provider{tpl, a, b}, nil, exportNow)
	if len(all.Providers) != 2 {
		t.Fatalf("不带 -only 应导出除模板外全部，得到 %d", len(all.Providers))
	}
	if _, err := ExportProviders([]Provider{a}, []string{"nope"}, exportNow); err == nil {
		t.Fatal("未知 slug 应报错")
	}
}

func TestExportJSONFieldNames(t *testing.T) {
	p := mustParse(t, validYAML)
	p.Models[0].Recommended = true
	cat, _ := ExportProviders([]Provider{p}, nil, exportNow)
	data, err := json.Marshal(cat)
	if err != nil {
		t.Fatal(err)
	}
	var raw struct {
		Providers []map[string]json.RawMessage `json:"providers"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		t.Fatal(err)
	}
	for _, k := range []string{"slug", "name", "base_url", "openai_compatible", "free", "requirements", "signup_url",
		"status", "status_note", "verified_at", "sources", "recommended_model", "models", "usage_notes"} {
		if _, ok := raw.Providers[0][k]; !ok {
			t.Errorf("缺少字段 %q", k)
		}
	}
	if _, ok := raw.Providers[0]["anthropic_endpoint"]; ok {
		t.Error("anthropic_endpoint 为空时应省略")
	}
	if !bytes.Contains(data, []byte(`"requirements":{"credit_card":"unknown","real_name":"unknown","region":"unknown"}`)) {
		t.Errorf("requirements 应使用 snake_case 且省略空的 other: %s", data)
	}
	if !bytes.Contains(data, []byte(`"usage_notes":[]`)) || !bytes.Contains(data, []byte(`"pending_capabilities":[]`)) {
		t.Errorf("空列表应导出 []，不是 null: %s", data)
	}
}

func TestValidateRecommendedAndToolCalling(t *testing.T) {
	cases := []struct{ name, old, new, want string }{
		{"recommended 只能一个", "  - id: m2\n    capabilities: [文本生成]\n", "  - id: m2\n    capabilities: [文本生成]\n    recommended: true\n", ""},
		{"tool_calling 枚举", "  - id: m1\n    capabilities: [文本生成]\n", "  - id: m1\n    capabilities: [文本生成]\n    tool_calling: maybe\n", "tool_calling"},
		{"recommended 须为文本模型", "    capabilities: [文生图]\n", "    capabilities: [文生图]\n    recommended: true\n", "recommended"},
		{"usage_notes 不能为空", "sources:", "usage_notes: [\"\"]\nsources:", "usage_notes"},
	}
	for _, c := range cases {
		y := strings.Replace(validYAML, c.old, c.new, 1)
		if c.want == "" {
			// 再给 m1 加一个 recommended，构成两个
			y = strings.Replace(y, "  - id: m1\n    capabilities: [文本生成]\n", "  - id: m1\n    capabilities: [文本生成]\n    recommended: true\n", 1)
			c.want = "recommended: true"
		}
		errs := Validate(mustParse(t, y), testCaps, t.TempDir())
		if !strings.Contains(strings.Join(errs, "\n"), c.want) {
			t.Errorf("%s: 想要包含 %q 的错误，得到 %v", c.name, c.want, errs)
		}
	}
}
