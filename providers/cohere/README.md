# Cohere

## 当前状态

- 文档状态：官方限流页、定价说明、模型页、兼容 API 页已核对；控制台需登录，未使用 key 实测
- 免费方式：**Trial key 永久免费**，官方未写截止日期；每月最多 1,000 次调用，Chat 每模型 20 请求 / 分钟
- **限制**：定价页 FAQ 写明 Trial key 不得用于生产或商业用途
- 使用门槛：需注册 Cohere 账号；信用卡、手机号、地区限制官方文档未提（待核验）。个人账号只能用 Trial key
- 最后核验日期：2026-11-26

## 一句话说明

Cohere 为注册账号提供免费的 Trial key，可调用 Command 系列等模型；官方提供 OpenAI 兼容接口，但 Trial key 限量且禁止商用。

## 快速开始

1. 在 [Cohere 控制台](https://dashboard.cohere.com/api-keys) 注册并创建 Trial API key（页面需登录，本次未打开）。
2. 只在本地 shell 设置环境变量，不要提交到仓库。
3. OpenAI 兼容 Base URL 用 `https://api.cohere.ai/compatibility/v1`；模型 ID 从 [models.md](models.md) 选择。
4. 发最小 Chat 请求验证；注意每月 1,000 次调用上限，且不得用于生产 / 商用。

## 核心信息

| 项目 | 值 |
| --- | --- |
| 官网 / 文档 | https://cohere.com / https://docs.cohere.com |
| API key | https://dashboard.cohere.com/api-keys |
| OpenAI Base URL | `https://api.cohere.ai/compatibility/v1`（官方 Python 示例；对话 `/chat/completions`；音频转写示例用 `api.cohere.com`） |
| Anthropic 兼容 | 官方文档未见 |
| 免费类型 | 永久（Trial key），每月 1,000 次调用 |
| 免费模型 | 官方称 Trial key 可用全部模型与 API；限流页列出的 Chat 型号见 [models.md](models.md) |
| 限流（Trial） | 每月 1,000 次调用；Chat 每模型 20 请求 / 分钟；Embed 2,000 inputs / 分钟；Rerank 10 请求 / 分钟；TPM / 并发官方未公布 |
| 使用限制 | Trial key 不得用于生产或商业用途（定价页 FAQ） |
| 门槛 | 注册账号；信用卡 / 手机号 / 地区限制待核验 |
| 核验状态 | 官方页面已核对，未实测 |



## 目录

- [官方链接与核验记录](official-links.md)
- [模型与免费规则](models.md)
- [API 接入、鉴权和限流](api.md)
- [Claude Code](integrations/claude-code.md) · [OpenClaw](integrations/openclaw.md) · [OpenCode](integrations/opencode.md) · [Codex](integrations/codex.md) · [Hermes](integrations/hermes.md)
