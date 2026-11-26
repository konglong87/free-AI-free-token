package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"time"
)

// StaleItem 是一条过期或缺失核验日期的供应商记录。
type StaleItem struct {
	Slug, Name, VerifiedAt string
	Days                   int // 距今天数；日期缺失或无效时为 -1
	Reason                 string
}

// FindStale 返回 verified_at 为空、无效或早于 days 天前的供应商（跳过模板）。
// 最需要处理的排前面：先是缺失 / 无效，再按过期天数从多到少。
func FindStale(providers []Provider, today time.Time, days int) []StaleItem {
	var items []StaleItem
	for _, p := range providers {
		if p.Slug == templateSlug {
			continue
		}
		it := StaleItem{Slug: p.Slug, Name: p.Name, VerifiedAt: p.VerifiedAt, Days: -1}
		if p.VerifiedAt == "" {
			it.Reason = "未填写核验日期"
			items = append(items, it)
			continue
		}
		t, err := time.Parse("2006-01-02", p.VerifiedAt)
		if err != nil {
			it.Reason = "日期格式无效"
			items = append(items, it)
			continue
		}
		it.Days = int(today.Sub(t).Hours() / 24)
		if it.Days > days {
			it.Reason = fmt.Sprintf("超过 %d 天未核验", days)
			items = append(items, it)
		}
	}
	sort.SliceStable(items, func(i, j int) bool {
		if (items[i].Days < 0) != (items[j].Days < 0) {
			return items[i].Days < 0
		}
		return items[i].Days > items[j].Days
	})
	return items
}

// RenderStale 输出 Markdown，可直接追加到 GitHub Actions 的 job summary。
func RenderStale(items []StaleItem, total, days int, today time.Time) string {
	var b strings.Builder
	fmt.Fprintf(&b, "## 供应商核验过期检查\n\n截至 %s，共 %d 家供应商，阈值 %d 天。\n\n", today.Format("2006-01-02"), total, days)
	if len(items) == 0 {
		b.WriteString("✅ 没有过期的供应商。\n")
		return b.String()
	}
	fmt.Fprintf(&b, "⚠️ %d 家需要重新核验：\n\n| 供应商 | verified_at | 已过天数 | 原因 |\n|---|---|---|---|\n", len(items))
	for _, it := range items {
		va, d := it.VerifiedAt, fmt.Sprint(it.Days)
		if va == "" {
			va = "-"
		}
		if it.Days < 0 {
			d = "-"
		}
		fmt.Fprintf(&b, "| %s (`%s`) | %s | %s | %s |\n", it.Name, it.Slug, va, d, it.Reason)
	}
	return b.String()
}

func runStale(root string, days int, today time.Time, out io.Writer) int {
	providers, errs := LoadProviders(root)
	if len(errs) > 0 {
		for _, e := range errs {
			fmt.Fprintln(os.Stderr, "错误:", e)
		}
		return 1
	}
	total := 0
	for _, p := range providers {
		if p.Slug != templateSlug {
			total++
		}
	}
	fmt.Fprint(out, RenderStale(FindStale(providers, today, days), total, days, today))
	return 0
}

// staleCommand 只报告，过期不会让命令失败（退出码 0）。
func staleCommand(args []string) int {
	fs := flag.NewFlagSet("stale", flag.ExitOnError)
	root := fs.String("root", "", "仓库根目录（默认从当前目录向上查找）")
	days := fs.Int("days", 45, "verified_at 早于多少天算过期")
	fs.Parse(args)
	r, err := resolveRoot(*root)
	if err != nil {
		fatal(err)
	}
	return runStale(r, *days, time.Now(), os.Stdout)
}
