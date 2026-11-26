# OpenCode 接入 Cerebras Inference

## 当前状态

官方有接入页，**未实测**。

Cerebras 官方集成页提供 OpenCode 接入说明，具体 provider 配置以官方页面为准。

## 可用于配置的字段

| 字段 | 值 |
| --- | --- |
| API Key | 本人的 Cerebras API key；不可提交到仓库 |
| OpenAI Base URL | `https://api.cerebras.ai/v1` |
| Model ID | `gpt-oss-120b`（试用层 5 RPM）或 `qwen-3.8-27b` |

## 说明与待核验

- 官方 OpenCode 页的模型表仍列着已弃用的 `gemma-4-31b`，以弃用页为准，不要使用。
- 官方页的完整配置格式本目录未逐字核对，待核验。
- 试用层限流较低，先用短问答验证。

## 验证记录

| 日期 | OpenCode 版本 | 模型 | 结果 | 备注 |
| --- | --- | --- | --- | --- |
| 2026-11-26 | 待核验 | 待核验 | 待核验 | 仅整理官方文档，未持有 key，未实测 |

## 来源

- [官方 OpenCode 接入页](https://inference-docs.cerebras.ai/integrations/opencode)
- [弃用页](https://inference-docs.cerebras.ai/support/deprecation)
