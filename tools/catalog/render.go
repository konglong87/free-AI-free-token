package main

import (
	"fmt"
	"html"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Capabilities 是 docs/model-capabilities.md「统一标签」表里的标签（名称 → 图标）。
type Capabilities map[string]string

func (c Capabilities) Has(name string) bool { _, ok := c[name]; return ok }

// LoadCapabilities 解析「## 统一标签」下表格的第一列，形如「📝 文本生成」。
func LoadCapabilities(root string) (Capabilities, error) {
	data, err := os.ReadFile(filepath.Join(root, "docs", "model-capabilities.md"))
	if err != nil {
		return nil, err
	}
	return ParseCapabilities(string(data))
}

func ParseCapabilities(md string) (Capabilities, error) {
	caps := Capabilities{}
	in := false
	for _, line := range strings.Split(md, "\n") {
		if strings.HasPrefix(line, "## ") {
			in = strings.TrimSpace(strings.TrimPrefix(line, "## ")) == "统一标签"
			continue
		}
		if !in || !strings.HasPrefix(line, "|") {
			continue
		}
		cell := strings.TrimSpace(strings.Split(line, "|")[1])
		parts := strings.SplitN(cell, " ", 2)
		if len(parts) == 2 && parts[0] != "" && parts[0][0] >= 0x80 {
			caps[strings.TrimSpace(parts[1])] = parts[0]
		}
	}
	if len(caps) == 0 {
		return nil, fmt.Errorf("docs/model-capabilities.md 中没有解析到统一标签")
	}
	return caps, nil
}

const (
	tableHead = `<table width="100%">
<thead>
<tr><th width="30%">供应商 · 免费能力</th><th width="70%">代表性模型 · 接入</th></tr>
</thead>
<tbody>`
	tableFoot = `</tbody>
</table>`
)

var statusIcons = map[string]string{"ok": "✅", "pending": "🟡"}

// 首页左列的免费方式与门槛文字；unknown 不显示。
var (
	freeLabels = map[string]string{"permanent": "长期免费", "daily": "每日额度", "one_time": "一次性赠金", "promo": "限时活动"}
	reqLabels  = []struct {
		get            func(Requirements) string
		required, none string
	}{
		{func(r Requirements) string { return r.CreditCard }, "需绑卡", "无需绑卡"},
		{func(r Requirements) string { return r.RealName }, "需实名", "无需实名"},
		{func(r Requirements) string { return r.Region }, "有地区限制", "无地区限制"},
	}
)

// accessSummary 返回左列「免费方式 / 门槛」两行；都未知时为空。
func accessSummary(p Provider) string {
	var lines []string
	if l, ok := freeLabels[p.Free.Type]; ok {
		lines = append(lines, "免费方式："+l)
	}
	var reqs []string
	for _, r := range reqLabels {
		switch r.get(p.Requirements) {
		case "required":
			reqs = append(reqs, r.required)
		case "none":
			reqs = append(reqs, r.none)
		}
	}
	if len(reqs) > 0 {
		lines = append(lines, "门槛："+strings.Join(reqs, " · "))
	}
	if len(lines) == 0 {
		return ""
	}
	return "<br><br><sub>" + strings.Join(lines, "<br>") + "</sub>"
}

// LatestVerified 返回所有供应商（不含模板）里最新的 verified_at，作为首页的更新日期。
func LatestVerified(providers []Provider) string {
	latest := ""
	for _, p := range providers {
		if p.Slug != templateSlug && p.VerifiedAt > latest {
			latest = p.VerifiedAt
		}
	}
	return latest
}

// ProvidersOn 返回出现在某份首页上的供应商（已按 home.order 排序），不含模板。
func ProvidersOn(providers []Provider, page string) (list []Provider, count int) {
	for _, p := range providers {
		if p.Slug == templateSlug {
			continue
		}
		if p.onPage(page) {
			list = append(list, p)
			count++
		}
	}
	return
}

func esc(s string) string { return html.EscapeString(s) }

// RenderTable 生成某份首页的供应商表（不含标记注释）。
func RenderTable(providers []Provider, caps Capabilities, page string) string {
	list, _ := ProvidersOn(providers, page)
	var b strings.Builder
	b.WriteString(tableHead + "\n")
	for _, p := range list {
		b.WriteString(renderRow(p, caps, page) + "\n")
	}
	b.WriteString(tableFoot)
	return b.String()
}

func renderRow(p Provider, caps Capabilities, page string) string {
	var tags []string
	for _, t := range p.Home.Tags {
		tags = append(tags, caps[t]+" "+t)
	}
	if len(p.Home.PendingTags) > 0 {
		tags = append(tags, "（待核验）")
	}

	api := p.BaseURL
	if u, ok := p.Home.PageAPIURL[page]; ok {
		api = u
	}
	links := fmt.Sprintf(`<a href="%s">API</a> · <a href="%s">%s</a> · <a href="providers/%s/models.md">完整模型</a>`,
		esc(api), esc(p.Home.Signup.URL), esc(p.Home.Signup.Label), p.Slug)

	status := "状态：" + statusIcons[p.Home.Status]
	if p.Home.StatusNote != "" {
		status += " " + esc(p.Home.StatusNote)
	}
	for _, e := range p.Home.Extras {
		if len(e.Pages) == 0 || contains(e.Pages, page) {
			status += fmt.Sprintf(` · <a href="%s">%s</a>`, esc(e.URL), esc(e.Label))
		}
	}

	return fmt.Sprintf(`<tr>
<td valign="top"><a href="providers/%s/"><strong>%s</strong></a><br><sub>%s</sub>%s</td>
<td valign="top">%s<br><br>%s<br><sub>%s</sub></td>
</tr>`, p.Slug, esc(p.Name), strings.Join(tags, "<br>"), accessSummary(p), renderGroups(p, page), links, status)
}

// renderGroups 把该首页上的模型按列表顺序分组：相邻且能力、待核验、分组说明都相同的合并成一组。
func renderGroups(p Provider, page string) string {
	type group struct {
		title string
		items []string
	}
	var groups []group
	lastKey := ""
	for _, m := range p.Models {
		if !m.onPage(page) {
			continue
		}
		title := strings.Join(m.Capabilities, " · ")
		if m.Home.GroupNote != "" {
			title += " · " + m.Home.GroupNote
		}
		if len(m.PendingCapabilities) > 0 {
			title += "（待核验）"
		}
		item := "<code>" + esc(m.ID) + "</code>"
		if m.Home.Bold {
			item = "<strong>" + item + "</strong>"
		}
		if m.Home.Label != "" {
			item = "<strong>" + esc(m.Home.Label) + "</strong><br>" + item
		}
		if m.Home.Suffix != "" {
			item += " " + esc(m.Home.Suffix)
		}
		if len(groups) == 0 || title != lastKey {
			groups = append(groups, group{title: title})
		}
		lastKey = title
		g := &groups[len(groups)-1]
		g.items = append(g.items, item)
	}
	var out []string
	for _, g := range groups {
		out = append(out, "<strong>"+esc(g.title)+"</strong><br>"+strings.Join(g.items, "<br>"))
	}
	return strings.Join(out, "<br><br>")
}

const (
	tableStart = "<!-- providers:start -->"
	tableEnd   = "<!-- providers:end -->"
	countStart = "<!-- providers-count:start -->"
	countEnd   = "<!-- providers-count:end -->"
	dateStart  = "<!-- updated:start -->"
	dateEnd    = "<!-- updated:end -->"
)

var (
	badgeRe     = regexp.MustCompile(`(badge/providers-)\d+`)
	strongRe    = regexp.MustCompile(`(<strong>)\d+(</strong>)`)
	badgeDateRe = regexp.MustCompile(`\d{4}--\d{2}--\d{2}`) // shields.io 徽章里的 - 要写成 --
	dateRe      = regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
)

// Splice 把生成内容写进首页文本的标记区域：供应商表、所有供应商数量标记区域，以及更新日期标记区域（date 为空时不改）。
func Splice(doc, table string, count int, date string) (string, error) {
	doc, err := replaceBetween(doc, tableStart, tableEnd, func(string) string { return "\n" + table + "\n" })
	if err != nil {
		return "", err
	}
	n := fmt.Sprint(count)
	doc, err = replaceAllBetween(doc, countStart, countEnd, func(region string) string {
		if strings.TrimSpace(region) != "" && strings.Trim(region, "0123456789") == "" {
			return n
		}
		region = badgeRe.ReplaceAllString(region, "${1}"+n)
		return strongRe.ReplaceAllString(region, "${1}"+n+"${2}")
	})
	if err != nil || date == "" {
		return doc, err
	}
	return replaceAllBetween(doc, dateStart, dateEnd, func(region string) string {
		region = badgeDateRe.ReplaceAllString(region, strings.ReplaceAll(date, "-", "--"))
		return dateRe.ReplaceAllString(region, date)
	})
}

func replaceBetween(doc, start, end string, fn func(string) string) (string, error) {
	if strings.Count(doc, start) != 1 || strings.Count(doc, end) != 1 {
		return "", fmt.Errorf("需要恰好一对标记 %s … %s", start, end)
	}
	i := strings.Index(doc, start) + len(start)
	j := strings.Index(doc, end)
	if j < i {
		return "", fmt.Errorf("标记顺序错误：%s 应在 %s 之前", start, end)
	}
	return doc[:i] + fn(doc[i:j]) + doc[j:], nil
}

func replaceAllBetween(doc, start, end string, fn func(string) string) (string, error) {
	if strings.Count(doc, start) != strings.Count(doc, end) {
		return "", fmt.Errorf("标记 %s 与 %s 数量不一致", start, end)
	}
	var out strings.Builder
	for {
		i := strings.Index(doc, start)
		if i < 0 {
			out.WriteString(doc)
			return out.String(), nil
		}
		i += len(start)
		j := strings.Index(doc[i:], end)
		if j < 0 {
			return "", fmt.Errorf("标记 %s 缺少结束标记", start)
		}
		out.WriteString(doc[:i] + fn(doc[i:i+j]))
		doc = doc[i+j:]
	}
}
