# NVIDIA NIM

## 当前状态

- 文档状态：已整理，待实测
- API 状态：官方入口已记录，待实测
- 免费模型状态：参考项目当前清单
- 接入文档状态：待核验
- 最后核验日期：2026-08-27

## 一句话说明

NVIDIA NIM 提供托管在 NVIDIA API Catalog 的免费模型，并暴露 OpenAI 风格的统一推理入口。

## 适合场景

- 推理和代码原型
- 长上下文评估
- 多模型对比
- agent 接入

## 快速开始

1. 打开 [NVIDIA API Key 页面](https://build.nvidia.com/settings/api-keys) 并登录。
2. 创建 API key。
3. 设置 Base URL `https://integrate.api.nvidia.com/v1`。
4. 从 [模型列表](models.md) 选择模型。
5. 发送最小 Chat Completions 请求并记录结果。

## 核心信息

- 官网：[https://www.nvidia.com/](https://www.nvidia.com/)
- 模型目录：[https://build.nvidia.com/models](https://build.nvidia.com/models)
- 文档：[https://docs.nvidia.com/nim/](https://docs.nvidia.com/nim/)
- 接入地址：`https://integrate.api.nvidia.com/v1`
- 推荐免费模型：`z-ai/glm-5.2`、`poolside/laguna-xs-2.1`
- 是否 OpenAI-compatible：官方 API 采用 OpenAI 风格；完整能力待实测
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
