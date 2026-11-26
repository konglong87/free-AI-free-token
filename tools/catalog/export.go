package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

// ExportSchemaVersion 在 JSON 结构发生不兼容变化时加一。
const ExportSchemaVersion = 1

// 导出 JSON 的结构。数值字段为 0、字符串为空时省略，表示未公布或未记录。
type (
	ExportCatalog struct {
		SchemaVersion int              `json:"schema_version"`
		GeneratedAt   string           `json:"generated_at"`
		Providers     []ExportProvider `json:"providers"`
	}
	ExportProvider struct {
		Slug              string        `json:"slug"`
		Name              string        `json:"name"`
		BaseURL           string        `json:"base_url"`
		OpenAICompatible  string        `json:"openai_compatible"`
		AnthropicEndpoint string        `json:"anthropic_endpoint,omitempty"`
		Free              ExportFree    `json:"free"`
		Requirements      Requirements  `json:"requirements"`
		SignupURL         string        `json:"signup_url"`
		Status            string        `json:"status"`
		StatusNote        string        `json:"status_note,omitempty"`
		VerifiedAt        string        `json:"verified_at"`
		Sources           []ExportLink  `json:"sources"`
		RecommendedModel  *ExportModel  `json:"recommended_model"`
		Models            []ExportModel `json:"models"`
		UsageNotes        []string      `json:"usage_notes"`
	}
	ExportFree struct {
		Type string `json:"type"`
		Note string `json:"note,omitempty"`
	}
	ExportLink struct {
		Label string `json:"label"`
		URL   string `json:"url"`
	}
	ExportModel struct {
		ID                  string   `json:"id"`
		Capabilities        []string `json:"capabilities"`
		PendingCapabilities []string `json:"pending_capabilities"`
		ToolCalling         string   `json:"tool_calling"`
		Context             string   `json:"context,omitempty"`
		RPM                 int      `json:"rpm,omitempty"`
		RPD                 int      `json:"rpd,omitempty"`
		TPM                 int      `json:"tpm,omitempty"`
		TPD                 int      `json:"tpd,omitempty"`
	}
)

func (r Requirements) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		CreditCard string `json:"credit_card"`
		RealName   string `json:"real_name"`
		Region     string `json:"region"`
		Other      string `json:"other,omitempty"`
	}{r.CreditCard, r.RealName, r.Region, r.Other})
}

func exportModel(m Model) ExportModel {
	tc := m.ToolCalling
	if tc == "" {
		tc = "pending"
	}
	return ExportModel{
		ID: m.ID, Capabilities: nonNil(m.Capabilities), PendingCapabilities: nonNil(m.PendingCapabilities),
		ToolCalling: tc, Context: m.Context, RPM: m.RPM, RPD: m.RPD, TPM: m.TPM, TPD: m.TPD,
	}
}

func nonNil(s []string) []string {
	if s == nil {
		return []string{}
	}
	return s
}

// ExportProviders 把供应商转成导出结构（跳过模板）。only 非空时只导出其中的 slug，并按 only 的顺序输出；
// 找不到的 slug 返回错误。
func ExportProviders(providers []Provider, only []string, now time.Time) (ExportCatalog, error) {
	bySlug := map[string]Provider{}
	var selected []Provider
	for _, p := range providers {
		if p.Slug == templateSlug {
			continue
		}
		bySlug[p.Slug] = p
		if len(only) == 0 {
			selected = append(selected, p)
		}
	}
	for _, s := range only {
		p, ok := bySlug[s]
		if !ok {
			return ExportCatalog{}, fmt.Errorf("找不到供应商 %q", s)
		}
		selected = append(selected, p)
	}

	out := ExportCatalog{SchemaVersion: ExportSchemaVersion, GeneratedAt: now.UTC().Format(time.RFC3339), Providers: []ExportProvider{}}
	for _, p := range selected {
		ep := ExportProvider{
			Slug: p.Slug, Name: p.Name, BaseURL: p.BaseURL, OpenAICompatible: p.OpenAICompatible,
			AnthropicEndpoint: p.AnthropicEndpoint, Free: ExportFree{p.Free.Type, p.Free.Note},
			Requirements: p.Requirements, SignupURL: p.Home.Signup.URL, Status: p.Home.Status,
			StatusNote: p.Home.StatusNote, VerifiedAt: p.VerifiedAt, Sources: []ExportLink{},
			Models: []ExportModel{}, UsageNotes: nonNil(p.UsageNotes),
		}
		for _, s := range p.Sources {
			ep.Sources = append(ep.Sources, ExportLink{s.Label, s.URL})
		}
		for _, m := range p.Models {
			em := exportModel(m)
			ep.Models = append(ep.Models, em)
			if m.Recommended {
				rm := em
				ep.RecommendedModel = &rm
			}
		}
		out.Providers = append(out.Providers, ep)
	}
	return out, nil
}

func runExport(root string, only []string, now time.Time, out io.Writer) int {
	providers, errs := LoadProviders(root)
	if len(errs) > 0 {
		for _, e := range errs {
			fmt.Fprintln(os.Stderr, "错误:", e)
		}
		return 1
	}
	cat, err := ExportProviders(providers, only, now)
	if err != nil {
		fmt.Fprintln(os.Stderr, "错误:", err)
		return 1
	}
	for _, p := range cat.Providers {
		if p.RecommendedModel == nil {
			fmt.Fprintf(os.Stderr, "提示: %s 没有设置 recommended: true 的模型，recommended_model 为 null\n", p.Slug)
		}
	}
	enc := json.NewEncoder(out)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	if err := enc.Encode(cat); err != nil {
		fmt.Fprintln(os.Stderr, "错误:", err)
		return 1
	}
	return 0
}

func exportCommand(args []string) int {
	fs := flag.NewFlagSet("export", flag.ExitOnError)
	root := fs.String("root", "", "仓库根目录（默认从当前目录向上查找）")
	only := fs.String("only", "", "只导出这些供应商，逗号分隔的 slug，按给定顺序输出")
	outPath := fs.String("o", "", "写入这个文件（默认输出到标准输出）")
	fs.Parse(args)
	r, err := resolveRoot(*root)
	if err != nil {
		fatal(err)
	}
	var slugs []string
	for _, s := range strings.Split(*only, ",") {
		if s = strings.TrimSpace(s); s != "" {
			slugs = append(slugs, s)
		}
	}
	var out io.Writer = os.Stdout
	if *outPath != "" {
		f, err := os.Create(*outPath)
		if err != nil {
			fatal(err)
		}
		defer f.Close()
		out = f
	}
	return runExport(r, slugs, time.Now(), out)
}
