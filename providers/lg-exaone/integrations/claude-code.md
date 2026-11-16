# Claude Code 接入 K-EXAONE 2.5

## 当前状态

不支持直连。Claude Code 使用 Anthropic Messages 协议（`/v1/messages`），EXAONE API Platform 只提供 OpenAI-compatible Chat Completions，官方也没有 Claude Code 接入指南。

## 可选做法（待核验）

如确需在 Claude Code 中使用，需要在本地自行运行一个把 Anthropic Messages 转成 OpenAI Chat Completions 的代理，再把 `ANTHROPIC_BASE_URL` 指向该代理。上游字段：

| 字段 | 值 |
| --- | --- |
| 上游 Base URL | `https://api.exaone.ai/v1` |
| 上游 API Key | `EXAONE_API_KEY`（不可提交到仓库） |
| 上游 Model ID | `k-exaone_v2.5` |

代理会改变工具调用和推理内容的格式，本仓库未验证任何具体代理与 K-EXAONE 2.5 的组合。需要开箱即用的 Agent 接入，优先用官方支持的 [Hermes](hermes.md) 或 [OpenCode](opencode.md)。

来源：[EXAONE 官方 Agent 接入页](https://docs.exaone.ai/ko/docs/agent-start)。
