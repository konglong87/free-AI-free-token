# ModelScope

## 当前状态

- 文档状态：已整理，待实测
- API 状态：官方推理入口已记录，待实测
- 免费模型状态：参考项目当前清单
- 接入文档状态：待核验
- 最后核验日期：2026-08-27

## 一句话说明

ModelScope 提供模型社区和在线推理 API，当前参考清单包含多款可免费调用的模型。

## 适合场景

- 中文模型评估
- 推理与代码原型
- 多模型对比
- agent 接入

## 快速开始

1. 注册或登录 [ModelScope](https://modelscope.cn/)。
2. 在 [访问令牌页面](https://modelscope.cn/my/myaccesstoken) 创建 Token。
3. 设置 Base URL `https://api-inference.modelscope.cn/v1`。
4. 从 [模型列表](models.md) 选择模型。
5. 发送最小 Chat Completions 请求验证。

## 核心信息

- 官网：[https://modelscope.cn/](https://modelscope.cn/)
- 模型列表：[https://modelscope.cn/models](https://modelscope.cn/models)
- 文档：[https://modelscope.cn/docs](https://modelscope.cn/docs)
- 接入地址：`https://api-inference.modelscope.cn/v1`
- 推荐免费模型：`MiniMax/MiniMax-M2.5`、`qwen-qwen3-5-35b-a3b`
- 是否 OpenAI-compatible：参考入口采用 OpenAI 风格；完整能力待核验
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
