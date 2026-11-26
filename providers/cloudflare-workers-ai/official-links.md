# 官方链接与核验记录

## 官方入口

- [Workers AI 文档首页](https://developers.cloudflare.com/workers-ai/)
- [模型目录](https://developers.cloudflare.com/workers-ai/models/)（页面显示 76 个模型，文本生成 41 个）
- [定价与免费额度](https://developers.cloudflare.com/workers-ai/platform/pricing/)
- [限流](https://developers.cloudflare.com/workers-ai/platform/limits/)
- [REST 入门](https://developers.cloudflare.com/workers-ai/get-started/rest-api/)
- [OpenAI 兼容](https://developers.cloudflare.com/workers-ai/configuration/open-ai-compatibility/)
- [错误码](https://developers.cloudflare.com/workers-ai/platform/errors/)
- [更新日志](https://developers.cloudflare.com/workers-ai/changelog/)
- [模型下线公告（2026-11-05）](https://developers.cloudflare.com/changelog/post/2026-11-05-planned-model-deprecations/)
- [模型搜索 API](https://developers.cloudflare.com/api/resources/ai/subresources/models/methods/list/)
- [控制台](https://dash.cloudflare.com)（需登录，本次访问返回 403）
- [状态页](https://www.cloudflarestatus.com)
- [服务条款](https://www.cloudflare.com/terms/)

## 来源核验记录

| 日期 | 来源 | 核验内容 | 结果 |
| --- | --- | --- | --- |
| 2026-11-26 | [定价](https://developers.cloudflare.com/workers-ai/platform/pricing/) | Workers Free 每日 10,000 Neurons；每日 00:00 UTC 重置；超额后操作报错；Workers Paid 每日同样含 10,000，超出 $0.011 / 1000 Neurons；2026-11-18 版写明部分模型需付费方式 | 已核对 |
| 2026-11-26 | [错误码](https://developers.cloudflare.com/workers-ai/platform/errors/) | 超额：错误码 3036、HTTP 429，提示已用完每日 10,000 Neurons | 已核对 |
| 2026-11-26 | [限流](https://developers.cloudflare.com/workers-ai/platform/limits/) | 文本生成 300 请求 / 分钟；需 Workers Paid 的模型：标准计费 20 / 分钟、预充值 AI Gateway credits 50 / 分钟；嵌入 3000（`bge-large-en-v1.5` 为 1500）；文生图 720；语音识别 720 | 已核对 |
| 2026-11-26 | [OpenAI 兼容](https://developers.cloudflare.com/workers-ai/configuration/open-ai-compatibility/) | Base URL、`/v1/chat/completions`、`/v1/embeddings`；`/v1/responses` 仅 gpt-oss 两款且只支持 `stream:false`；API token 需 Workers AI Read + Edit 权限 | 已核对，未实测 |
| 2026-11-26 | [模型目录](https://developers.cloudflare.com/workers-ai/models/)及各模型页 | 模型 ID、上下文、能力徽标；抽查 9 个模型页无「Paid access required」 | 已核对（抽查）；其余模型待核验 |
| 2026-11-26 | [下线公告](https://developers.cloudflare.com/changelog/post/2026-11-05-planned-model-deprecations/) | 2026-11-20 起 24 个旧模型返回 HTTP 410（错误码 5028） | 已核对 |
| 2026-11-26 | 官方文档检索 | Anthropic 兼容端点 | 未见 |
| 2026-11-26 | 官方 `agent-setup/claude-code` 页与 `agents/harnesses/opencode` 页 | 前者讲 Claude Code 使用 Cloudflare 的 Skills 与 MCP，后者讲在 Durable Object 中运行 OpenCode，均不是把 Workers AI 当模型后端 | 已核对 |
