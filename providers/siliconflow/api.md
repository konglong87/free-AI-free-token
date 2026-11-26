# 硅基流动 API 接入

## 协议与地址

- OpenAI 兼容：`https://api.siliconflow.cn/v1`，对话 `POST /chat/completions`（完整地址 `https://api.siliconflow.cn/v1/chat/completions`）
- Anthropic 兼容：`POST https://api.siliconflow.cn/v1/messages`（官方 [Messages API](https://docs.siliconflow.cn/docs/api/messages-post)）
- 模型列表：`GET /v1/models`（官方 [Models API](https://docs.siliconflow.cn/docs/api/models-get)，未实测）
- 鉴权：`Authorization: Bearer <API key>`
- 域名说明：国内站 `api.siliconflow.cn`；国际站 `api.siliconflow.com` 是另一套账号，不要混用 key

## 最小请求

```bash
export SILICONFLOW_API_KEY="<你的 key>"
curl 'https://api.siliconflow.cn/v1/chat/completions' \
  -H "Authorization: Bearer ${SILICONFLOW_API_KEY}" \
  -H 'Content-Type: application/json' \
  -d '{"model":"Qwen/Qwen3-8B","messages":[{"role":"user","content":"你好"}]}'
```

以上为按官方文档整理的示例，**尚未用真实 key 调用**。

## 限流与用量

- 免费模型限速为固定值，具体 RPM / TPM 写在登录后的模型详情页，官方文档未公开数值（待核验）。
- 官方只给出总体范围：语言模型 RPM 1000–10000、TPM 50000–5000000；L0 等级 RPM 1,000、TPM 40,000，仅适用于收费模型。
- 2026-11-16 起「动态限流」按模型分批开启，限速随负载浮动，开启后的模型不再沿用 L0–L5 等级。
- 限流按账户计，每个模型单独计，不是按 key。超限返回 429：`Request was rejected due to rate limiting`。
- Rate Limits 页仍保留 DeepSeek-R1 / V3 的 RPH / RPD 旧条款，与 2025-03-07 公告取消限制的说法冲突，待核验。
- 来源：[Rate Limits](https://docs.siliconflow.cn/docs/userguide/faqs/rate-limit-and-upgradation)。

## 门槛

- 实名认证：支付宝人脸识别，支持身份证、港澳台通行证和居住证、外国人永久居留证；未满 18 岁无法实名；部分虚拟运营商号段无法注册或绑定手机号（[实名认证 FAQ](https://docs.siliconflow.cn/docs/userguide/faqs/authentication)）。
- 不需要信用卡；充值走支付宝、微信或对公转账；免费模型本身不需要充值。

## 常见错误

| 错误 | 可能原因 | 处理方式 |
| --- | --- | --- |
| 401 | key 错误，或国内站 / 国际站 key 混用 | 重新生成 key，确认域名与账号站点一致 |
| 403 | 未实名认证或模型无权限 | 完成实名认证，对照模型广场 |
| 404 | 模型 ID 错误（如缺少或多了 `Pro/` 前缀） | 对照公开模型中心的 ID |
| 429 | 触发限流 | 降低并发，等待恢复；动态限流下限速会浮动 |

## OpenAI / Anthropic 兼容说明

- 官方同时提供 Chat Completions 与 Messages 两种格式。
- 免费模型是否都能走 Messages 端点、工具调用是否稳定，官方文档未保证，待核验。
- streaming、结构化输出是否支持因模型而异，待核验。
