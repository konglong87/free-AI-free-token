# OpenClaw 接入 K-EXAONE 2.5

## 当前状态

待核验。EXAONE 官方未提供 OpenClaw 接入指南；以下按 OpenAI-compatible 接口字段整理，**未实测**。

## 可用于自定义 Provider 的字段

| 字段 | 值 |
| --- | --- |
| Endpoint compatibility | OpenAI-compatible (`/chat/completions`) |
| Base URL | `https://api.exaone.ai/v1` |
| API Key | 本人的 EXAONE API Key；不可提交到仓库 |
| Model ID | `k-exaone_v2.5` |

在 OpenClaw 的自定义 Provider 中填写上述字段，不要把 `/chat/completions` 写进 Base URL。先用短问答验证，再测试流式、工具调用和推理内容（`reasoning_content`）的回传。

来源：[EXAONE Quick Start](https://docs.exaone.ai/ko/docs/quick-start)。
