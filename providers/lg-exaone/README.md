# LG AI Research K-EXAONE 2.5

## 当前状态

- 文档状态：官方公告、接口文档、Agent 接入指南与服务条款已核对，API 调用待实测
- 免费方式：LG AI Research 2026-11-09 公告 K-EXAONE 2.5 API **免费开放**，拿到 API Key 即可调用；官方称其为「技术验证模型」，免费期限未公布
- 使用限制：服务条款限定**非商业用途**；API Key **自签发日起 30 天有效**，到期自动停用
- 最后核验日期：2026-11-16

## 一句话说明

K-EXAONE 2.5 是 LG AI Research 自研的 Agent 向文本大模型（韩语 / 英语），在 EXAONE API Platform 以 OpenAI-compatible Chat Completions 接口免费提供，官方给出 Hermes 与 OpenCode 接入指南。

## 快速开始

1. 打开 [EXAONE API Platform](https://docs.exaone.ai)，登录后点击「API Key 발급하기」（签发 API Key）。签发前先读[服务条款](https://docs.exaone.ai/ko/terms/terms-of-service)：非商业用途、30 天有效、输入输出默认可用于训练（可拒绝）。
2. 在本地 shell 设置 `EXAONE_API_KEY`，不要把真实密钥提交到仓库。
3. 模型 ID 用 `k-exaone_v2.5`，Base URL 用 `https://api.exaone.ai/v1`，见 [API 接入](api.md)。
4. 发一次最小请求，确认返回、限流响应头后再接入 Agent。记下签发日期，30 天后需重新签发。

## 核心信息

| 项目 | 值 |
| --- | --- |
| 官网 / 申请入口 | https://docs.exaone.ai |
| 模型 ID | `k-exaone_v2.5`（下划线，不是 `K-EXAONE-2.5`） |
| OpenAI Chat Base URL | `https://api.exaone.ai/v1` |
| 官方标称上下文 | 524,288 Token（512K），输入 + `max_tokens` 超出返回 400 |
| 输入 / 输出 | 仅文本；响应含推理 Token（`reasoning_effort` 可调） |
| 免费方式 | 技术验证阶段免费，非商用；有调用上限，具体数值未公开 |
| Key 有效期 | 30 天，到期自动停用 |
| 核验状态 | 官方网页与文档已核对；未持有账号 / Key，未实测申请和 API |

## ⚠️ 不要和 LG ThinQ API 搞混

搜索「LG API」常会遇到 **LG ThinQ API**（`connect-pat.lgthinq.com`），那是控制 LG 智能家电的平台，不提供大模型 Token。本目录只收录 EXAONE API Platform。

## 目录

- [官方链接与核验记录](official-links.md)
- [模型与免费规则](models.md)
- [API 接入、鉴权和限流](api.md)
- [Hermes](integrations/hermes.md) · [OpenCode](integrations/opencode.md) · [Codex（待核验）](integrations/codex.md) · [OpenClaw（待核验）](integrations/openclaw.md) · [Claude Code（不支持直连）](integrations/claude-code.md)
