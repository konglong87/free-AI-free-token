# Cerebras Inference

## 当前状态

- 文档状态：官方文档、定价页、限流页、弃用页已核对；控制台需登录，未使用 key 实测
- 免费方式：**仅新用户一次性赠金** $5（试用额度），激活后 **30 天过期**；没有永久免费层，也没有每日循环额度
- 使用门槛：**需先绑定有效支付方式**才能获得 $5（官方称不会扣费、不会自动开通付费）；到期或用完后 API 与 Playground 暂停
- 试用层上下文约为付费层的一半（64K–65K）；限流以组织为单位
- 最后核验日期：2026-11-26

## 一句话说明

Cerebras 提供高速推理 API 和 OpenAI 兼容接口；新账号绑定支付方式后获得 $5 一次性试用额度，可调用 `gpt-oss-120b` 与 `qwen-3.8-27b`。

## 快速开始

1. 在 [Cerebras Cloud](https://cloud.cerebras.ai) 注册并按官方要求绑定有效支付方式，领取 $5 试用额度（30 天内有效）。
2. 在控制台创建 API key，只在本地 shell 设置环境变量，不要提交到仓库。
3. Base URL 用 `https://api.cerebras.ai/v1`，模型 ID 用 `gpt-oss-120b` 或 `qwen-3.8-27b`。
4. 发最小 Chat Completions 请求验证，并在控制台 Limits 页确认当前账号的实际限额。

## 核心信息

| 项目 | 值 |
| --- | --- |
| 官网 / 定价 | https://www.cerebras.ai/pricing |
| 控制台 / API key | https://cloud.cerebras.ai（官方文档的 Get an API key 入口；需登录，本次未打开） |
| 文档 | https://inference-docs.cerebras.ai/introduction |
| OpenAI Base URL | `https://api.cerebras.ai/v1`（对话 `/chat/completions`） |
| Anthropic 兼容 | 官方文档未见 |
| 免费类型 | 一次性赠金 $5，30 天过期 |
| 免费模型 | `gpt-oss-120b`、`qwen-3.8-27b`（试用与付费层均可用） |
| 试用层上下文 | `gpt-oss-120b` 65k；`qwen-3.8-27b` 64K（65,536）；付费层约 131k / 128K |
| 限流（Free Trial） | `gpt-oss-120b`：5 RPM、总 TPM 90K、TPD 1M；`qwen-3.8-27b`：300 RPM、总 TPM 750K（临时上调，标准值 450K） |
| 门槛 | 绑定有效支付方式；手机号 / 实名 / 地区限制官方页面未见（待核验） |
| 核验状态 | 官方页面已核对，未实测 |



## 目录

- [官方链接与核验记录](official-links.md)
- [模型与免费规则](models.md)
- [API 接入、鉴权和限流](api.md)
- [Claude Code](integrations/claude-code.md) · [OpenClaw](integrations/openclaw.md) · [OpenCode](integrations/opencode.md) · [Codex](integrations/codex.md) · [Hermes](integrations/hermes.md)
