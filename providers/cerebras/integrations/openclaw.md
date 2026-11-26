# OpenClaw 接入 Cerebras Inference

## 当前状态

待核验。官方集成页未列出 OpenClaw；以下按 OpenAI 兼容接口字段整理，**未实测**。

字段来自官方 OpenAI 兼容说明。

## 可用于配置的字段

| 字段 | 值 |
| --- | --- |
| API Key | 本人的 Cerebras API key；不可提交到仓库 |
| OpenAI Base URL | `https://api.cerebras.ai/v1` |
| Model ID | `gpt-oss-120b`（试用层 5 RPM）或 `qwen-3.8-27b` |

## 说明与待核验

- 官方集成页（Aider、Cerebras for VS Code、Cline、KiloCode、OpenCode）没有 OpenClaw。
- 不要把 `/chat/completions` 写进 Base URL。
- 先用短问答验证，再测试流式与工具调用。

## 验证记录

| 日期 | OpenClaw 版本 | 模型 | 结果 | 备注 |
| --- | --- | --- | --- | --- |
| 2026-11-26 | 待核验 | 待核验 | 待核验 | 仅整理官方文档，未持有 key，未实测 |

## 来源

- [官方集成列表](https://inference-docs.cerebras.ai/integrations)
- [OpenAI 兼容说明](https://inference-docs.cerebras.ai/resources/openai)
