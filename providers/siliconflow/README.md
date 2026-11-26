# 硅基流动 SiliconFlow（国内站）

## 当前状态

- 文档状态：官方文档、价格页、公开模型中心已核对；未登录控制台，未使用 API key 实测
- 免费方式：**永久免费模型**（价格页标「免费」的小模型，调用费用为 0，限速免费，不是赠金，也不是每日循环额度）；官方保留免费模型的解释权，模型可能随时下线
- 使用门槛：**必须实名认证**（支付宝人脸识别）；不需要信用卡，免费模型不需要充值
- 注意：国内站（siliconflow.cn）与国际站（siliconflow.com）是两套独立站点和账号；国际站当前未见免费模型，本目录只收录国内站
- 最后核验日期：2026-11-26

## 一句话说明

硅基流动国内站提供 OpenAI 兼容和 Anthropic 兼容接口，价格页上标「免费」的一批小参数模型可以免费通过 API 调用，需要实名认证，限速固定且具体数值在登录后的模型详情页。

## 快速开始

1. 打开[硅基流动控制台](https://cloud.siliconflow.cn)，登录后按官方要求完成[实名认证](https://docs.siliconflow.cn/docs/userguide/faqs/authentication)（2026-05-15 起未实名无法使用平台功能）。
2. 在 [API 密钥页](https://cloud.siliconflow.cn/account/ak) 创建 key；只在本地 shell 设置环境变量，不要提交到仓库。
3. 从[模型与免费规则](models.md)选择免费模型，Base URL 用 `https://api.siliconflow.cn/v1`。
4. 用最小 Chat Completions 请求验证，并在控制台模型详情页记录该模型的实际限速。

## 核心信息

| 项目 | 值 |
| --- | --- |
| 官网 | https://www.siliconflow.cn/ |
| 控制台 | https://cloud.siliconflow.cn |
| 文档 | https://docs.siliconflow.cn/docs/userguide/quickstart |
| OpenAI Base URL | `https://api.siliconflow.cn/v1`（对话 `POST /chat/completions`） |
| Anthropic 兼容 | 有：`POST https://api.siliconflow.cn/v1/messages`；Claude Code 的 `ANTHROPIC_BASE_URL` 官方写 `https://api.siliconflow.cn/`（无 `/v1`） |
| 免费类型 | 永久免费模型（限速免费）；官方可随时调整或下线 |
| 免费模型 | 对话 / 视觉类 10 个（见 [models.md](models.md)），另有嵌入、重排、生图、语音类免费模型 |
| 限流 | 免费模型限速为固定值，具体 RPM / TPM 需登录后在模型详情页查看，**官方文档未公开数值（待核验）**；2026-11-16 起部分模型开启动态限流 |
| 门槛 | 实名认证；短信或邮箱登录；不要求信用卡；地区限制官方文档未见明确写法 |
| 核验状态 | 官方页面已核对，未实测 |

## 国际站说明

国际站 `siliconflow.com`（Base URL `https://api.siliconflow.com/v1`）定价页写「可从 $1 免费额度开始」，页面上没有 $0 模型；官方更新日志显示 `tencent/Hunyuan-MT-7B`、`THUDM/GLM-4-9B-0414`、`THUDM/GLM-Z1-9B-0414` 已于 2026-03-12 在国际站下线。该赠金是否一次性、有无有效期官方页面未写（待核验）。因此国际站不作为免费 API 收录。
来源：[国际站定价页](https://www.siliconflow.com/pricing)、[国际站 Rate Limits](https://docs.siliconflow.com/en/userguide/rate-limits/rate-limit-and-upgradation)。

## 目录

- [官方链接与核验记录](official-links.md)
- [模型与免费规则](models.md)
- [API 接入、鉴权和限流](api.md)
- [Claude Code](integrations/claude-code.md) · [OpenClaw](integrations/openclaw.md) · [OpenCode](integrations/opencode.md) · [Codex](integrations/codex.md) · [Hermes](integrations/hermes.md)
