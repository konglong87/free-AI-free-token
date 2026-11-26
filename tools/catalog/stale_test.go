package main

import (
	"bytes"
	"strings"
	"testing"
	"time"
)

func TestFindStale(t *testing.T) {
	today := time.Date(2026, 11, 26, 0, 0, 0, 0, time.UTC)
	ps := []Provider{
		{Slug: "fresh", Name: "Fresh", VerifiedAt: "2026-11-01"},
		{Slug: "edge", Name: "Edge", VerifiedAt: "2026-10-12"}, // 恰好 45 天，不算过期
		{Slug: "old", Name: "Old", VerifiedAt: "2026-06-01"},
		{Slug: "older", Name: "Older", VerifiedAt: "2026-01-01"},
		{Slug: "empty", Name: "Empty"},
		{Slug: "bad", Name: "Bad", VerifiedAt: "yesterday"},
		{Slug: templateSlug, Name: "模板"},
	}
	items := FindStale(ps, today, 45)
	var got []string
	for _, it := range items {
		got = append(got, it.Slug)
	}
	if strings.Join(got, ",") != "empty,bad,older,old" {
		t.Fatalf("结果或顺序不对: %v", got)
	}
	if items[3].Days != 178 {
		t.Fatalf("天数不对: %d", items[3].Days)
	}
}

func TestRenderStale(t *testing.T) {
	today := time.Date(2026, 11, 26, 0, 0, 0, 0, time.UTC)
	md := RenderStale([]StaleItem{
		{Slug: "old", Name: "Old", VerifiedAt: "2026-06-01", Days: 178, Reason: "超过 45 天未核验"},
		{Slug: "empty", Name: "Empty", Days: -1, Reason: "未填写核验日期"},
	}, 5, 45, today)
	for _, want := range []string{"## 供应商核验过期检查", "共 5 家", "2 家需要重新核验", "| Old (`old`) | 2026-06-01 | 178 |", "| Empty (`empty`) | - | - | 未填写核验日期 |"} {
		if !strings.Contains(md, want) {
			t.Fatalf("缺少 %q:\n%s", want, md)
		}
	}
	if !strings.Contains(RenderStale(nil, 5, 45, today), "✅") {
		t.Fatal("没有过期时应输出 ✅")
	}
}

func TestRunStaleNeverFails(t *testing.T) {
	root := smokeRoot(t, "http://127.0.0.1:1") // demo 的 verified_at 是 2026-01-02
	var out bytes.Buffer
	if code := runStale(root, 45, time.Date(2026, 11, 26, 0, 0, 0, 0, time.UTC), &out); code != 0 {
		t.Fatalf("stale 只报告，应返回 0，得到 %d", code)
	}
	if !strings.Contains(out.String(), "Demo (`demo`)") {
		t.Fatalf("输出不对:\n%s", out.String())
	}
}
