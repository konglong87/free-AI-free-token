# Codex 接入 Cloudflare Workers AI

## 当前状态

待核验。官方文档未见 Codex 接入说明；以下按 OpenAI 兼容接口字段整理，**未实测**。

字段来自官方 OpenAI 兼容页。

## 可用于配置的字段

| 字段 | 值 |
| --- | --- |
| API Token | 本人的 Cloudflare API token（Workers AI Read + Edit）；不可提交到仓库 |
| OpenAI Base URL | `https://api.cloudflare.com/client/v4/accounts/{account_id}/ai/v1`（替换账号 ID） |
| Model ID | `@cf/openai/gpt-oss-120b`（带 `@cf/` 前缀） |

## 说明与待核验

- 官方文档没有 Codex 的接入页。
- Base URL 中的 `{account_id}` 要替换为自己的账号 ID；不要把 `/chat/completions` 写进 Base URL。
- 先用短问答验证，再测试流式与工具调用；留意每日 Neurons 额度。

## 验证记录

| 日期 | Codex 版本 | 模型 | 结果 | 备注 |
| --- | --- | --- | --- | --- |
| 2026-11-26 | 待核验 | 待核验 | 待核验 | 仅整理官方文档，未持有 key，未实测 |

## 来源

- [OpenAI 兼容](https://developers.cloudflare.com/workers-ai/configuration/open-ai-compatibility/)
- [REST 入门](https://developers.cloudflare.com/workers-ai/get-started/rest-api/)
