# 官方链接与核验记录

## 官方入口

- [官网](https://cohere.com)
- [控制台 / API keys](https://dashboard.cohere.com/api-keys)（需登录，HTTP 200 但需登录后查看）
- [文档首页](https://docs.cohere.com)
- [模型列表](https://docs.cohere.com/docs/models)
- [限流（含 Trial key 限制）](https://docs.cohere.com/docs/rate-limits)
- [定价说明](https://docs.cohere.com/v2/docs/how-does-cohere-pricing-work)、[定价页与 FAQ](https://cohere.com/pricing)
- [OpenAI 兼容 API](https://docs.cohere.com/docs/compatibility-api)
- [FAQ](https://docs.cohere.com/docs/cohere-faqs)
- [状态页](https://status.cohere.com)
- [服务条款](https://cohere.com/terms-of-use)

## 来源核验记录

| 日期 | 来源 | 核验内容 | 结果 |
| --- | --- | --- | --- |
| 2026-11-26 | [限流](https://docs.cohere.com/docs/rate-limits) | Trial key 每月 1,000 次调用；Chat 每模型 20 请求 / 分钟；Embed 2,000 inputs / 分钟、Embed(Images) 5、EmbedJob 5 请求 / 分钟、Rerank 10、Tokenize 100、Audio Transcriptions 5、Parse 500，其余默认 500 请求 / 分钟；Production key 在 Command A、R+、R、R7B、North Mini Code 为 500 请求 / 分钟 | 已核对 |
| 2026-11-26 | [定价说明](https://docs.cohere.com/v2/docs/how-does-cohere-pricing-work) | Trial key 可用全部模型与 API | 已核对 |
| 2026-11-26 | [定价页](https://cohere.com/pricing) FAQ | 「Trial key 的调用免费；但限速且不得用于生产或商业用途」 | 已核对；条款页未逐段核对 |
| 2026-11-26 | [OpenAI 兼容页](https://docs.cohere.com/docs/compatibility-api) | Base URL、chat completions（流式、结构化输出、工具调用）、embeddings、audio transcriptions；不支持的参数；`reasoning_effort` 只接受 `none` 和 `high` | 已核对，未实测 |
| 2026-11-26 | [模型列表](https://docs.cohere.com/docs/models) | 模型 ID、上下文、最大输出、模态；旧别名弃用与退役日期 | 已核对 |
| 2026-11-26 | 官方文档检索 | Anthropic 兼容端点 | 未见（文档 MCP server 仅供编码代理读文档） |
