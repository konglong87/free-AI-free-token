# Google AI Studio（Gemini API）

## 当前状态

- 文档状态：官方定价页、限流页、模型页、弃用页、OpenAI 兼容页、条款页已核对；AI Studio 控制台需登录，未使用 key 实测
- 免费方式：**永久免费层**（新账号默认处于 Free Tier，官方未写截止日期），按项目限流；升级付费层才需要绑定计费
- **数据条款**：免费层的输入输出会被 Google 用于改进产品，人工审核员可能读取；EEA、瑞士、英国用户适用付费级数据条款（详见 [models.md](models.md)）
- 使用门槛：官方文档对免费层没有写信用卡要求（待核验）；第二个及以后的免费项目需设备验证（扫码，不要手机号）
- 注意：2026-12-01 起 `gemini-2.5-pro`、`gemini-2.5-flash` 的请求会被自动路由到新模型
- 最后核验日期：2026-11-26

## 一句话说明

Google AI Studio 的 Gemini API 提供永久免费层，覆盖多个 Flash 系列和 `gemini-4-hydrogen`；同时提供 OpenAI 兼容端点。

## 快速开始

1. 打开 [Google AI Studio](https://aistudio.google.com) 登录，按[官方 API key 说明](https://aistudio.google.com/docs/api-key)创建 key（创建页需登录，本次未打开）。
2. 只在本地 shell 设置环境变量，不要提交到仓库。
3. OpenAI 兼容 Base URL 用 `https://generativelanguage.googleapis.com/v1beta/openai/`；模型 ID 从 [models.md](models.md) 选免费模型。
4. 在 [AI Studio 限额页](https://aistudio.google.com/rate-limit)（需登录）查看自己项目的 RPM / TPM / RPD。

## 核心信息

| 项目 | 值 |
| --- | --- |
| 官网 | https://ai.google.dev/（跳转到 https://ai.google/） |
| 控制台 | https://aistudio.google.com |
| 文档 | https://aistudio.google.com/docs |
| OpenAI Base URL | `https://generativelanguage.googleapis.com/v1beta/openai/`（官方原文） |
| Anthropic 兼容 | 官方文档未见 |
| 免费类型 | 永久免费层，按项目限流 |
| 免费模型 | 文本 / 多模态 9 个（`gemini-4-hydrogen`、`gemini-3.8-flash`、`gemini-3.6-flash`、`gemini-3.5-flash-lite` 等），另有音频、TTS、embedding 类 |
| 限流 | 官方页面只公布 `gemini-decision-1`（200 RPM / 50,000 TPM / 10,000 RPD）；其余免费模型需登录 AI Studio 查看，**待核验** |
| 门槛 | 免费层未写信用卡要求（待核验）；第二个及以后免费项目需设备验证；可用地区含英国、瑞士、EEA |
| 核验状态 | 官方页面已核对，未实测 |



## 目录

- [官方链接与核验记录](official-links.md)
- [模型与免费规则](models.md)
- [API 接入、鉴权和限流](api.md)
- [Claude Code](integrations/claude-code.md) · [OpenClaw](integrations/openclaw.md) · [OpenCode](integrations/opencode.md) · [Codex](integrations/codex.md) · [Hermes](integrations/hermes.md)
