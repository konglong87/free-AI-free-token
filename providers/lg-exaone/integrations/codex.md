# Codex CLI 接入 K-EXAONE 2.5

## 当前状态

待核验。EXAONE 官方未提供 Codex 接入指南，也只提供 Chat Completions 接口，没有 OpenAI Responses 接口。如果你的 Codex 版本只支持 Responses（`wire_api = "responses"`），就无法直连，需要自行加一层 Chat Completions → Responses 的转换代理。

## 可用于自定义 Provider 的字段

| 字段 | 值 |
| --- | --- |
| Base URL | `https://api.exaone.ai/v1` |
| API Key 环境变量 | `EXAONE_API_KEY`（不可提交到仓库） |
| Model ID | `k-exaone_v2.5` |
| 预期路径 | `/chat/completions` |

在当前 Codex 版本支持 Chat Completions 的前提下，把上述字段写入本地 `~/.codex/config.toml` 的自定义 provider，先验证短问答，再验证流式和工具调用。具体配置语法与兼容性仍待核验。

来源：[EXAONE Quick Start](https://docs.exaone.ai/ko/docs/quick-start)、[K-EXAONE 2.5 接口说明](https://docs.exaone.ai/ko/docs/k-exaone-2.5)。
