# 科大讯飞星火 Spark-X2.6

## 当前状态

- 文档状态：官方公开页面已核对，API 调用待实测
- 免费额度：讯飞星辰 MaaS 首页标注「零门槛领取10亿tokens免费额度」，**限时赠额，不是永久免费模型**；领取资格、有效期及是否自动付费以登录后的活动规则为准
- 用户提供的资料称额度有效期为 31 天；公开页面未展示该条款，**尚未独立核实**
- 最后核验日期：2026-10-29

## 一句话说明

Spark-X2.6 是科大讯飞的文本生成模型，官方公布 1M 上下文、293B-A30B MoE 架构，并提供 OpenAI Chat、Responses 和 Anthropic Messages 接口。免费额度需在讯飞星辰 MaaS 平台领取；额度用完后按当时生效价格计费。

## 快速开始

1. 打开 [讯飞星辰 MaaS](https://maas.xfyun.cn/)；首页有 Spark-X2.6 的「10亿 tokens 免费额度」入口。登录后先核对活动条款、余额与付费设置，再领取并开通模型。
2. 在控制台「API调用 / APIKEY管理」获取个人 API Key；不要把真实密钥提交到仓库。
3. 选用模型 ID `spark-x2.6`，按客户端协议选择不同 Base URL，见 [API 接入](api.md)。
4. 发一次最小请求，确认模型响应、额度扣减与限制，再用于长任务或多 Agent。

## 核心信息

| 项目 | 值 |
| --- | --- |
| 官网 / 领取入口 | https://maas.xfyun.cn/ |
| 模型 ID | `spark-x2.6`（不是 `Spark-X2.6`，也不是 4B / Preview） |
| OpenAI Chat Base URL | `https://maas-api.cn-huabei-1.xf-yun.com/v2` |
| OpenAI Responses Base URL | `https://maas-api.cn-huabei-1.xf-yun.com/v1` |
| Anthropic / Claude Code Base URL | `https://maas-api.cn-huabei-1.xf-yun.com/anthropic` |
| 官方标称上下文 | 1M Token；不同协议和客户端仍须遵守各自接口限制 |
| 免费方式 | 活动赠额（官方首页标注 10 亿 tokens），非永久免费 |
| 核验状态 | 官方网页与文档已核对；未持有账号 / Key，未实测领取和 API |

## 目录

- [官方链接与核验记录](official-links.md)
- [模型与免费规则](models.md)
- [API 接入、鉴权和限流](api.md)
- [Claude Code](integrations/claude-code.md) · [Codex](integrations/codex.md) · [OpenClaw](integrations/openclaw.md) · [Hermes](integrations/hermes.md) · [OpenCode（待核验）](integrations/opencode.md)
