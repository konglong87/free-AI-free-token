# Cloudflare Workers AI

## 当前状态

- 文档状态：官方定价、限流、模型目录、OpenAI 兼容页、错误码页、下线公告已核对；控制台需登录，未使用 token 实测
- 免费方式：**每日循环**，Workers Free 账户每天 10,000 Neurons，每日 00:00 UTC 重置；超额后报错（错误码 3036，HTTP 429），不自动扣费
- 使用门槛：需要 Cloudflare 账号、账号 ID 和 API token；信用卡、手机号、地区限制官方 Workers AI 文档未提及（待核验）
- 注意：2026-11-20 起 24 个旧模型已下线（HTTP 410）；部分新模型需 Workers Paid，免费账户不可用
- 最后核验日期：2026-11-26

## 一句话说明

Cloudflare Workers AI 提供 OpenAI 兼容的 `/ai/v1` 接口，Workers Free 账户每天有 10,000 Neurons 免费额度，可调用 gpt-oss、Gemma、Qwen、GLM、Llama 等模型。

## 快速开始

1. 注册 Cloudflare 账号，记下账号 ID，并按[官方 REST 入门](https://developers.cloudflare.com/workers-ai/get-started/rest-api/)创建带 `Workers AI - Read` 与 `Workers AI - Edit` 权限的 API token（控制台需登录，本次未打开）。
2. 只在本地 shell 设置环境变量，不要提交到仓库。
3. Base URL 用 `https://api.cloudflare.com/client/v4/accounts/{account_id}/ai/v1`（替换 `{account_id}`），模型 ID 带 `@cf/` 前缀。
4. 发最小 Chat Completions 请求验证，并留意每日 10,000 Neurons 额度消耗。

## 核心信息

| 项目 | 值 |
| --- | --- |
| 官网 / 文档 | https://developers.cloudflare.com/workers-ai/ |
| 控制台 | https://dash.cloudflare.com（需登录；本次访问返回 403，待核验） |
| OpenAI Base URL | `https://api.cloudflare.com/client/v4/accounts/{account_id}/ai/v1`（对话 `/chat/completions`，另有 `/embeddings`） |
| Anthropic 兼容 | 官方文档未见 |
| 免费类型 | 每日循环：10,000 Neurons / 天，每日 00:00 UTC 重置 |
| 免费模型 | 9 个已抽查的文本 / 视觉模型（见 [models.md](models.md)）；另有文生图、语音、嵌入类，免费资格待核验 |
| 限流 | 文本生成 300 请求 / 分钟；需 Workers Paid 的模型 20 请求 / 分钟；TPM / TPD / 并发官方未公布 |
| 门槛 | Cloudflare 账号 + 账号 ID + API token；信用卡 / 手机号 / 地区限制待核验 |
| 核验状态 | 官方页面已核对，未实测 |



## 目录

- [官方链接与核验记录](official-links.md)
- [模型与免费规则](models.md)
- [API 接入、鉴权和限流](api.md)
- [Claude Code](integrations/claude-code.md) · [OpenClaw](integrations/openclaw.md) · [OpenCode](integrations/opencode.md) · [Codex](integrations/codex.md) · [Hermes](integrations/hermes.md)
