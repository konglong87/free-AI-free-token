# Claude Code 接入 Cloudflare Workers AI

## 当前状态

不支持直连。Claude Code 使用 Anthropic Messages 协议，Workers AI 官方文档未见 Anthropic 兼容端点。

Cloudflare 官方 `agent-setup/claude-code` 页讲的是 Claude Code 如何使用 Cloudflare 的 Skills 与 MCP，不是把 Workers AI 当 Claude Code 后端。

## 可用于配置的字段

| 字段 | 值 |
| --- | --- |
| 上游 Base URL | `https://api.cloudflare.com/client/v4/accounts/{account_id}/ai/v1` |
| 上游鉴权 | `CLOUDFLARE_API_TOKEN`（不可提交到仓库） |
| 上游 Model ID | `@cf/openai/gpt-oss-120b` |

## 说明与待核验

- 如确需使用，需要本地自建 Anthropic → OpenAI 的转换代理，本目录未验证任何具体代理。
- 每日 10,000 Neurons 很快会被 Claude Code 的多轮调用耗尽。

## 验证记录

| 日期 | Claude Code 版本 | 模型 | 结果 | 备注 |
| --- | --- | --- | --- | --- |
| 2026-11-26 | 待核验 | 待核验 | 待核验 | 仅整理官方文档，未持有 key，未实测 |

## 来源

- [官方 Claude Code 页（Skills / MCP）](https://developers.cloudflare.com/agent-setup/claude-code/)
- [OpenAI 兼容](https://developers.cloudflare.com/workers-ai/configuration/open-ai-compatibility/)
