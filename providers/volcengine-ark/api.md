# 火山方舟 API 接入

## 协议与地址

- OpenAI 兼容 Base URL：`https://ark.cn-beijing.volces.com/api/v3`，对话 `POST /chat/completions`（完整地址 `https://ark.cn-beijing.volces.com/api/v3/chat/completions`），另有 Responses API（[Base URL 及鉴权](https://docs.volcengine.com/docs/ark/base-url-and-authentication?lang=zh)）
- Anthropic 兼容：`POST https://ark.cn-beijing.volces.com/api/compatible/v1/messages`，鉴权头 `x-api-key`（[Messages API](https://docs.volcengine.com/docs/ark/messages-api?lang=zh)）
- 鉴权：`Authorization: Bearer $ARK_API_KEY`
- 模型列表接口：官方文档未见公开的 `/models` 路径（待核验）；模型 ID 请查[官方模型列表](https://docs.volcengine.com/docs/ark/model-list?lang=zh)
- 注意：Coding Plan 付费套餐使用专属端点 `https://ark.cn-beijing.volces.com/api/coding`（Anthropic）和 `.../api/coding/v3`（OpenAI），**不要混用**。官方警告 Coding Plan 用户用普通 `/api/v3` 不消耗套餐额度、会产生额外费用（[Coding Plan Claude Code 页](https://docs.volcengine.com/docs/ark/coding-plan-personal-ai-claude-code?lang=zh)）。

## 最小请求

```bash
export ARK_API_KEY="<你的 ARK API Key>"
curl 'https://ark.cn-beijing.volces.com/api/v3/chat/completions' \
  -H "Authorization: Bearer ${ARK_API_KEY}" \
  -H 'Content-Type: application/json' \
  -d '{"model":"doubao-seed-2-0-mini-260428","messages":[{"role":"user","content":"你好"}]}'
```

以上为按官方文档整理的示例，**尚未用真实 key 调用**；该模型是否有免费额度待核验。

## 限流与额度

- 每个账号下每个基础模型有 RPM 与 TPM 限制，所有接入点共享，只含按 token 后付费调用；具体数值见[官方模型列表](https://docs.volcengine.com/docs/ark/model-list?lang=zh)的「最大 RPM / TPM」，页面注明非刚性保障。部分模型另标最大并发（每 1M TPM）。
- RPD、TPD 官方未见。
- 可在控制台「开通管理」给模型或接入点设置用量上限，达到后自动暂停；控制台额度显示有小时级延迟（[开通管理](https://docs.volcengine.com/docs/ark/activation-management?lang=zh)）。
- 开启[安心体验模式](https://docs.volcengine.com/docs/ark/free-tokens-only-mode)后用完赠送额度返回 429 `SetLimitExceeded`，不会产生费用。

## 门槛

- 火山引擎账号；用完额度后继续使用需实名认证并手动开通模型服务。
- 信用卡、地区限制：官方文档未见（待核验）。

## 常见错误

| 错误 | 可能原因 | 处理方式 |
| --- | --- | --- |
| 401 | key 错误，或把 Coding Plan 专属 key 用在 `/api/v3` | 重新生成 key，确认端点与 key 类型匹配 |
| 404 | 模型 ID 错误（含日期后缀）或模型未开通 | 对照官方模型列表，在开通管理开通模型 |
| 429 `SetLimitExceeded` | 安心体验模式下赠送额度已用完 | 按需关闭安心体验模式（不可逆，之后按量计费），或等待下一份额度 |
| 429 其他 | 触发 RPM / TPM 限流 | 降低并发，等待恢复 |

## OpenAI / Anthropic 兼容说明

- Chat Completions、Responses、Messages 三种格式官方均有说明。
- streaming、工具调用、结构化输出：官方模型列表标注部分模型支持，具体行为待实测。
