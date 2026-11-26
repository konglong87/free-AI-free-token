# Codex 接入 Cohere

## 当前状态

待核验。官方文档未见 Codex 接入说明；以下按 OpenAI 兼容接口字段整理，**未实测**。

字段来自官方 OpenAI 兼容页。

## 可用于配置的字段

| 字段 | 值 |
| --- | --- |
| API Key | 本人的 Cohere Trial key；不可提交到仓库 |
| OpenAI Base URL | `https://api.cohere.ai/compatibility/v1` |
| Model ID | `command-a-03-2025` |

## 说明与待核验

- 官方文档没有 Codex 的接入页。
- 不要把 `/chat/completions` 写进 Base URL。
- 先用短问答验证；Trial key 每月 1,000 次、不得商用。

## 验证记录

| 日期 | Codex 版本 | 模型 | 结果 | 备注 |
| --- | --- | --- | --- | --- |
| 2026-11-26 | 待核验 | 待核验 | 待核验 | 仅整理官方文档，未持有 key，未实测 |

## 来源

- [OpenAI 兼容页](https://docs.cohere.com/docs/compatibility-api)
