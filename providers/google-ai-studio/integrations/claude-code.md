# Claude Code 接入 Google AI Studio（Gemini API）

## 当前状态

不支持直连。Claude Code 使用 Anthropic Messages 协议，Gemini API 官方只提供 OpenAI 兼容和原生接口，官方文档未见 Anthropic 兼容端点。

官方 Coding agents 页提供的 Gemini Docs MCP server 与 Agent Skills，是让 Claude Code 读取 Gemini 文档，不是把 Gemini 作为 Claude Code 的模型。

## 可用于配置的字段

| 字段 | 值 |
| --- | --- |
| 上游 Base URL | `https://generativelanguage.googleapis.com/v1beta/openai/` |
| 上游 API Key | `GEMINI_API_KEY`（不可提交到仓库） |
| 上游 Model ID | `gemini-3.8-flash` |

## 说明与待核验

- 如确需使用，需要本地自建 Anthropic → OpenAI 的转换代理，本目录未验证任何具体代理。
- 免费层数据会被用于改进产品，代理转发代码库内容前请评估风险。
- 官方 `gemini-api` CLI 页写明可与 Claude Code 配合使用，那同样是工具而非模型接入。

## 验证记录

| 日期 | Claude Code 版本 | 模型 | 结果 | 备注 |
| --- | --- | --- | --- | --- |
| 2026-11-26 | 待核验 | 待核验 | 待核验 | 仅整理官方文档，未持有 key，未实测 |

## 来源

- [官方 Coding agents 页](https://aistudio.google.com/docs/coding-agents)
- [官方 CLI 页](https://aistudio.google.com/docs/cli)
- [OpenAI 兼容](https://aistudio.google.com/docs/openai)
