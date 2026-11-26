# Hermes 接入 Google AI Studio（Gemini API）

## 当前状态

待核验。官方文档未见 Hermes 专门接入说明；以下按 OpenAI 兼容接口字段整理，**未实测**。

字段来自官方 OpenAI 兼容页。

## 可用于配置的字段

| 字段 | 值 |
| --- | --- |
| API Key | 本人的 Gemini API key；不可提交到仓库 |
| OpenAI Base URL | `https://generativelanguage.googleapis.com/v1beta/openai/` |
| Model ID | `gemini-3.8-flash`（示例；先在 AI Studio 确认限额） |

## 说明与待核验

- 官方文档没有 Hermes 的接入页。
- Base URL 末尾带斜杠；不要把 `chat/completions` 写进 Base URL。
- 先用短问答验证，再测试流式与工具调用；免费层数据会用于改进产品。

## 验证记录

| 日期 | Hermes 版本 | 模型 | 结果 | 备注 |
| --- | --- | --- | --- | --- |
| 2026-11-26 | 待核验 | 待核验 | 待核验 | 仅整理官方文档，未持有 key，未实测 |

## 来源

- [OpenAI 兼容](https://aistudio.google.com/docs/openai)
- [Coding agents 页](https://aistudio.google.com/docs/coding-agents)
