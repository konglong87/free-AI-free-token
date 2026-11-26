// catalog 读取 providers/*/provider.yaml，校验后生成 README.md 与 readme-B.md 的供应商表。
//
//	go run .          生成并写回两份首页
//	go run . -check   只检查：校验失败或生成结果与已提交内容不一致时以非零状态退出
//	go run . smoke    用环境变量里的 key 实测各家 /models 与 chat（-write 回写 verified_at）
//	go run . stale    列出 verified_at 为空或过期的供应商（Markdown）
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "smoke":
			os.Exit(smokeCommand(os.Args[2:]))
		case "stale":
			os.Exit(staleCommand(os.Args[2:]))
		}
	}
	root := flag.String("root", "", "仓库根目录（默认从当前目录向上查找）")
	check := flag.Bool("check", false, "只检查，不写文件")
	flag.Parse()

	r, err := resolveRoot(*root)
	if err != nil {
		fatal(err)
	}
	*root = r
	if code := run(*root, *check); code != 0 {
		os.Exit(code)
	}
}

func run(root string, check bool) int {
	providers, errs := LoadProviders(root)
	if len(errs) > 0 {
		for _, e := range errs {
			fmt.Fprintln(os.Stderr, "错误:", e)
		}
		return 1
	}
	caps, err := LoadCapabilities(root)
	if err != nil {
		fatal(err)
	}

	stale := false
	for page, name := range pageFiles {
		path := filepath.Join(root, name)
		old, err := os.ReadFile(path)
		if err != nil {
			fatal(err)
		}
		_, count := ProvidersOn(providers, page)
		next, err := Splice(string(old), RenderTable(providers, caps, page), count)
		if err != nil {
			fatal(fmt.Errorf("%s: %w", name, err))
		}
		if next == string(old) {
			continue
		}
		if check {
			fmt.Fprintf(os.Stderr, "%s 与 provider.yaml 生成结果不一致，请运行: cd tools/catalog && go run .\n", name)
			stale = true
			continue
		}
		if err := os.WriteFile(path, []byte(next), 0o644); err != nil {
			fatal(err)
		}
		fmt.Println("已更新", name)
	}
	if stale {
		return 1
	}
	if check {
		fmt.Println("catalog OK")
	}
	return 0
}

func resolveRoot(root string) (string, error) {
	if root != "" {
		return root, nil
	}
	return findRoot()
}

func findRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "docs", "model-capabilities.md")); err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("找不到仓库根目录（需要 docs/model-capabilities.md），请用 -root 指定")
		}
		dir = parent
	}
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, "错误:", err)
	os.Exit(1)
}
