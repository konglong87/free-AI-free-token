# Contribution Guide

## 新增供应商

```bash
cp -R providers/_template providers/<provider-slug>
```

然后按顺序填写：

1. `provider.yaml`：供应商的唯一结构化数据来源（见下文）
2. `official-links.md`
3. `models.md`
4. `api.md`
5. `integrations/*`
6. `README.md`

不要手改首页的供应商表。改完 yaml 后运行生成，再运行校验：

```bash
cd tools/catalog && go run .   # 重新生成 README.md / readme-B.md 的供应商表和数量
./scripts/validate-layout.sh   # 结构校验；已安装 go 时同时运行 go run . -check
```

## provider.yaml

字段说明见 `providers/_template/provider.yaml`。要点：

- 数据只来自该供应商的官方来源或本目录已有资料；查不到的写 `null` 或「待核验」，不要猜测。
- `home.tags`、模型的 `capabilities` 必须使用[统一能力标签](model-capabilities.md)；`home.order` 决定首页顺序，不可重复；国内供应商排在前面。
- `free.type`、`requirements` 核验后会显示在首页左列的「免费方式」「门槛」里；查不到就保持 `unknown`，首页不显示。
- 模型按列表顺序展示，`home.pages` 决定出现在哪份首页；相邻且能力相同的模型会合并成一组。
- 生成器会校验必填字段、枚举值、能力标签、slug 与目录名一致、同一家里模型 ID 不重复。
- `recommended: true` 标出推荐给代码智能体的默认模型（每家最多一个），`tool_calling` 记录工具调用是否经官方确认；`usage_notes` 写用户必须知道的限制（输入用于训练、禁止商用、需实名等）。这三项会进入导出的 JSON。

## 导出目录 JSON

给其他程序（例如桌面端内置的免费模型引导）使用的供应商目录：

```bash
cd tools/catalog
go run . export -o /tmp/catalog.json                                  # 导出全部供应商（不含模板）
go run . export -only zhipu-glm,siliconflow -o /tmp/catalog.json      # 只导出指定供应商，按给定顺序
```

结构：`{schema_version, generated_at, providers: [...]}`。每家包含 `slug`、`name`、`base_url`、`openai_compatible`、`anthropic_endpoint`（没有则省略）、`free{type,note}`、`requirements{credit_card,real_name,region,other}`、`signup_url`、`status`、`status_note`、`verified_at`、`sources[]`、`recommended_model`（没有设置则为 `null`）、`models[]`、`usage_notes[]`。模型的 `tool_calling` 未填写时导出为 `pending`；`rpm/rpd/tpm/tpd`、`context` 未公布时省略。结构有不兼容变化时 `schema_version` 加一。

## 实测

`tools/catalog` 提供两个子命令，用来给 `verified_at` 提供依据、并发现过期条目：

```bash
cd tools/catalog
go run . smoke                       # 实测所有设置了 key 的供应商
go run . smoke -provider nvidia-nim  # 只测一家
go run . smoke -write                # 两项都成功才把该家 verified_at 改成今天
go run . stale                       # 列出 verified_at 为空或超过 45 天的供应商（Markdown）
go run . stale -days 30              # 自定义天数
```

**`smoke` 做什么**：对每家先 `GET {base_url}/models`，再用一个模型发一次 `max_tokens: 1` 的 chat completion，每个请求只发一次、不重试，默认超时 20 秒（`-timeout` 可调）。输出一张表：供应商、`/models` 与 chat 的状态码、耗时、简短错误信息。没有 key 的供应商直接跳过；一个 key 都没有时输出「全部跳过」。任何一家失败，命令以非零状态退出。

**key 怎么给**：只通过环境变量，变量名是供应商 slug 转大写、`-` 换成 `_`，再加 `_API_KEY`，例如 `nvidia-nim` → `NVIDIA_NIM_API_KEY`，`b-ai` → `B_AI_API_KEY`。

```bash
NVIDIA_NIM_API_KEY=你的key go run . smoke -provider nvidia-nim
```

key 不要写进仓库或脚本；输出里出现的 key 会被替换成 `***`。

**用哪个模型**：默认取 `models[]` 里第一个上首页的「文本生成」模型。想指定别的，在该模型上加 `smoke: true`（每家最多一个，且必须是文本生成模型）。ID 含 `*` 的条目（如 `*:free`）是模型集合，不会被选作实测模型。

**`-write` 的边界**：只有 `/models` 和 chat 都返回 2xx 才会改 `verified_at`，且只改这一行。`home.status` 不会自动修改，两项都成功且当前不是 `ok` 时，输出会提示「可考虑改为 ok」，由人决定。改完后照常运行 `./scripts/validate-layout.sh`。

### GitHub Actions

- `.github/workflows/ci.yml`：push 和 PR 时运行 `go vet`、`go test`、`./scripts/validate-layout.sh` 和 `actionlint`。
- `.github/workflows/weekly.yml`：每周一 02:00 UTC 运行，也可在 Actions 页面手动触发（可填 `provider` 只测一家）。包含三个任务，结果都写进 job summary，不会向仓库提交任何改动：
  - 外部链接检查：检查仓库里 Markdown 的外部链接，失效只报告，不让任务失败。
  - 核验过期检查：运行 `stale`，只报告。
  - 供应商实测：运行 `smoke`；某家实测失败会让该任务变红。

要让某家参与定时实测，在仓库 Settings → Secrets and variables → Actions 里新建同名 Secret，名称与上面的环境变量相同，每家一个：

| 供应商 | Secret 名称 |
|---|---|
| agnes | `AGNES_API_KEY` |
| ai21-labs | `AI21_LABS_API_KEY` |
| b-ai | `B_AI_API_KEY` |
| cerebras | `CEREBRAS_API_KEY` |
| cohere | `COHERE_API_KEY` |
| google-ai-studio | `GOOGLE_AI_STUDIO_API_KEY` |
| groq | `GROQ_API_KEY` |
| iflytek-spark | `IFLYTEK_SPARK_API_KEY` |
| lg-exaone | `LG_EXAONE_API_KEY` |
| modelscope | `MODELSCOPE_API_KEY` |
| nvidia-nim | `NVIDIA_NIM_API_KEY` |
| openrouter | `OPENROUTER_API_KEY` |
| sensenova | `SENSENOVA_API_KEY` |
| siliconflow | `SILICONFLOW_API_KEY` |
| volcengine-ark | `VOLCENGINE_ARK_API_KEY` |
| zhipu-glm | `ZHIPU_GLM_API_KEY` |

没配置的 Secret 对应的供应商会被自动跳过。`base_url` 含 `{占位符}` 的供应商（如 `cloudflare-workers-ai`，URL 里有 `{account_id}`）无论有没有 key 都会被跳过，输出里标注「需手动实测」，不需要配 Secret。新增供应商后，除了建 Secret，还要把 `<SLUG>_API_KEY: ${{ secrets.<SLUG>_API_KEY }}` 加到 `weekly.yml` 的 `smoke` 步骤的 `env` 里，Actions 才会把它传给 `smoke`。

## 必填检查

新增或修改供应商后运行 `./scripts/validate-layout.sh`。

校验只检查结构和数据格式，不代表内容真实。内容真实性必须靠官方来源、原文来源或实际接入结果。

## 来源要求

每个关键结论都应该能追溯到官方来源，尤其是：

- 免费模型
- 是否 OpenAI compatible
- base URL
- model name
- agent / coding client 接入方式
- 核验状态

## 内容要求

只收集免费 LLM 接入所需信息，不收录无关介绍、营销内容或体验描述。

允许写入的信息：

- 官网和官方文档入口
- API key 获取入口
- 接入地址
- 免费模型
- OpenAI-compatible 状态
- Claude Code、OpenClaw、OpenCode、Codex、Hermes 接入方式
- 核验日期和核验状态

禁止写入的信息：

- 真实 API key、token、cookie、session
- 个人账号、手机号、邮箱、用户 ID
- 作者或转载平台的无关信息
- 截图中的敏感字段
- 无法追溯来源的结论

## 不要提交真实密钥

不要把 API key、token、cookie、控制台截图里的敏感字段提交到仓库。

## 首页能力标签

更新免费模型时，在 `provider.yaml` 里维护能力标签和模型分组，再运行生成，同步更新 `README.md` 与 `readme-B.md`，使用[统一能力标签](model-capabilities.md)。

- 先在供应商 `models.md` 记录能力与来源，再写入 `provider.yaml`；不要根据模型 ID 猜测。
- 区分图片理解、图片生成与视频生成；未确认的输入方式不要补写。
- 能力待核验时，供应商标签与模型分组都必须保留「待核验」。
- 保持两列表格，接入与证据链接放在模型下方；不要继续增加横向列。
- 修改后除结构校验外，还要检查桌面与窄屏渲染截图，确保标签、完整模型 ID 和链接不被横向裁切。
