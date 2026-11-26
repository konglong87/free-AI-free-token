# Claude Code 接入 Cerebras Inference

## 当前状态

不支持直连。Claude Code 使用 Anthropic Messages 协议，Cerebras 官方只提供 OpenAI 兼容接口，官方集成页也没有 Claude Code。

如确需使用，需要本地自建把 Anthropic Messages 转成 OpenAI Chat Completions 的代理，本目录未验证任何具体代理。

## 可用于配置的字段

| 字段 | 值 |
| --- | --- |
| 上游 Base URL | `https://api.cerebras.ai/v1` |
| 上游 API Key | `CEREBRAS_API_KEY`（不可提交到仓库） |
| 上游 Model ID | `gpt-oss-120b` / `qwen-3.8-27b` |

## 说明与待核验

- 代理会改变工具调用与推理内容的格式，兼容性待核验。
- `gpt-oss-120b` 拒绝同时带 `tools` 与 `response_format`，转换层需注意。
- 试用层 `gpt-oss-120b` 仅 5 RPM，不适合 Claude Code 这类高频调用。

## 验证记录

| 日期 | Claude Code 版本 | 模型 | 结果 | 备注 |
| --- | --- | --- | --- | --- |
| 2026-11-26 | 待核验 | 待核验 | 待核验 | 仅整理官方文档，未持有 key，未实测 |

## 来源

- [官方集成页（未列 Claude Code）](https://inference-docs.cerebras.ai/integrations)
- [OpenAI 兼容说明](https://inference-docs.cerebras.ai/resources/openai)
