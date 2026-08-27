# Groq

## 当前状态

- 文档状态：已整理，待实测
- API 状态：官方入口已记录，待实测
- 免费模型状态：公开参考清单
- 接入文档状态：待核验
- 最后核验日期：2026-08-27

## 一句话说明

Groq 提供高速推理 API，并通过 OpenAI-compatible endpoint 暴露多款免费模型。

## 适合场景

- 低延迟对话
- 代码和推理原型
- agent 接入
- 速率受限下的批量测试

## 快速开始

1. 打开 [Groq API Key 页面](https://console.groq.com/keys) 注册并创建 key。
2. 设置 Base URL `https://api.groq.com/openai/v1`。
3. 从 [模型列表](models.md) 选择模型。
4. 发送最小 Chat Completions 请求验证。

## 核心信息

- 官网：[https://groq.com/](https://groq.com/)
- 控制台：[https://console.groq.com/](https://console.groq.com/)
- 文档：[https://console.groq.com/docs](https://console.groq.com/docs)
- 接入地址：`https://api.groq.com/openai/v1`
- 推荐免费模型：`moonshotai/kimi-k2-instruct`、`groq/compound`
- 是否 OpenAI-compatible：官方提供 OpenAI-compatible endpoint；完整能力待实测
- 核验状态：公开入口已整理，API 与客户端待实测

## 目录

- [官方链接](official-links.md)
- [模型列表](models.md)
- [API 接入](api.md)
- [Claude Code 接入](integrations/claude-code.md)
- [OpenClaw 接入](integrations/openclaw.md)
- [OpenCode 接入](integrations/opencode.md)
- [Codex 接入](integrations/codex.md)
- [Hermes 接入](integrations/hermes.md)
