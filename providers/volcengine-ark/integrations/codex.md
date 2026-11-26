# Codex 接入 火山方舟（字节跳动 · 豆包）

## 当前状态

待核验。官方 Coding Plan 概览页列出了该客户端的接入说明，但针对付费套餐；普通免费额度 key 的接入未核对，也未实测。

以下按 OpenAI 兼容接口字段整理，配置格式以官方为准。

## 可用于配置的字段

| 字段 | 值 |
| --- | --- |
| API Key | 本人的 ARK API Key；不可提交到仓库 |
| OpenAI Base URL | `https://ark.cn-beijing.volces.com/api/v3` |
| Model ID | `doubao-seed-2-0-mini-260428`（示例；免费额度覆盖待核验） |

## 说明与待核验

- 官方 Coding Plan 概览页列出了 Codex，其配置使用 Coding Plan 专属端点，不适用于免费额度。
- 普通 key 使用 `/api/v3`；不要把 `/chat/completions` 写进 Base URL。
- 先用短问答验证，并开启安心体验模式避免超额计费。

## 验证记录

| 日期 | Codex 版本 | 模型 | 结果 | 备注 |
| --- | --- | --- | --- | --- |
| 2026-11-26 | 待核验 | 待核验 | 待核验 | 仅整理官方文档，未持有 key，未实测 |

## 来源

- [Coding Plan 概览（官方，付费套餐）](https://docs.volcengine.com/docs/ark/coding-plan-personal-plan-overview)
- [Base URL 及鉴权](https://docs.volcengine.com/docs/ark/base-url-and-authentication?lang=zh)
