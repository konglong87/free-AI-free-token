# 智谱 GLM

## 当前状态

- 文档状态：已整理，待实测
- API 状态：官方文档已核对
- 免费模型状态：官方模型概览和价格页已核对
- 接入文档状态：OpenAI-compatible 接入已整理，agent 专用配置待实测
- 最后核验日期：2026-06-25

## 一句话说明

智谱开放平台提供 GLM 系列模型 API，官方文档列出多个 Flash 免费模型，并提供兼容 OpenAI API 的接入方式。

## 快速开始

1. 打开智谱开放平台。
2. 创建或复制 API key。
3. 选择免费模型，例如 `glm-4.7-flash`。
4. 使用 Base URL `https://open.bigmodel.cn/api/paas/v4/`。
5. 在客户端中配置 Base URL、API key、模型名。

## 核心信息

- 官网：https://bigmodel.cn/
- 文档：https://docs.bigmodel.cn/
- API key：https://bigmodel.cn/usercenter/proj-mgmt/apikeys
- 接入地址：https://open.bigmodel.cn/api/paas/v4/
- Chat completions endpoint：`https://open.bigmodel.cn/api/paas/v4/chat/completions`
- 是否 OpenAI-compatible：是
- 推荐免费模型：`glm-4.7-flash`
- 核验状态：已整理，待实测

## 目录

- [官方链接](official-links.md)
- [模型列表](models.md)
- [API 接入](api.md)
- [Claude Code 接入](integrations/claude-code.md)
- [OpenClaw 接入](integrations/openclaw.md)
- [OpenCode 接入](integrations/opencode.md)
- [Codex 接入](integrations/codex.md)
- [Hermes 接入](integrations/hermes.md)
