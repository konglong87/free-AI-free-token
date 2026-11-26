# Claude Code 接入 Cohere

## 当前状态

不支持直连。Claude Code 使用 Anthropic Messages 协议，Cohere 官方文档未见 Anthropic 兼容端点，也没有 Claude Code 接入页。

如确需使用，需要本地自建 Anthropic → OpenAI 的转换代理，本目录未验证任何具体代理。

## 可用于配置的字段

| 字段 | 值 |
| --- | --- |
| 上游 Base URL | `https://api.cohere.ai/compatibility/v1` |
| 上游 API Key | `COHERE_API_KEY`（不可提交到仓库） |
| 上游 Model ID | `command-a-03-2025` |

## 说明与待核验

- Trial key 每月仅 1,000 次调用，且不得用于生产 / 商业用途，不适合 Claude Code 这类高频使用。
- 官方 MCP server 只是让编码代理读 Cohere 文档，不是模型接入。

## 验证记录

| 日期 | Claude Code 版本 | 模型 | 结果 | 备注 |
| --- | --- | --- | --- | --- |
| 2026-11-26 | 待核验 | 待核验 | 待核验 | 仅整理官方文档，未持有 key，未实测 |

## 来源

- [OpenAI 兼容页](https://docs.cohere.com/docs/compatibility-api)
- [限流](https://docs.cohere.com/docs/rate-limits)
