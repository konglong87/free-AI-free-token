# Cloudflare Workers AI API 接入

## 协议与地址

- OpenAI 兼容 Base URL：`https://api.cloudflare.com/client/v4/accounts/{account_id}/ai/v1`（替换 `{account_id}`）
- 对话：`/chat/completions`；嵌入：`/embeddings`。`/responses` 仅 `@cf/openai/gpt-oss-120b`、`@cf/openai/gpt-oss-20b` 支持，且只支持 `stream:false`（[OpenAI 兼容](https://developers.cloudflare.com/workers-ai/configuration/open-ai-compatibility/)）
- 原生 REST：`POST https://api.cloudflare.com/client/v4/accounts/{ACCOUNT_ID}/ai/run/@cf/...`（[REST 入门](https://developers.cloudflare.com/workers-ai/get-started/rest-api/)）
- 模型列表：OpenAI 兼容页未提 `/v1/models`；官方另有模型搜索接口 `GET /accounts/{account_id}/ai/models/search`
- 鉴权：`Authorization: Bearer <api_token>`；token 需 `Workers AI - Read` 与 `Workers AI - Edit` 权限
- Anthropic Messages 端点：官方文档未见

## 最小请求

```bash
export CLOUDFLARE_ACCOUNT_ID="<你的账号 ID>"
export CLOUDFLARE_API_TOKEN="<你的 API token>"
curl "https://api.cloudflare.com/client/v4/accounts/${CLOUDFLARE_ACCOUNT_ID}/ai/v1/chat/completions" \
  -H "Authorization: Bearer ${CLOUDFLARE_API_TOKEN}" \
  -H 'Content-Type: application/json' \
  -d '{"model":"@cf/openai/gpt-oss-120b","messages":[{"role":"user","content":"你好"}]}'
```

以上为按官方文档整理的示例，**尚未用真实 token 调用**。

## 限流与额度

见 [models.md](models.md)。要点：每日 10,000 Neurons（00:00 UTC 重置），超额 429；文本生成 300 请求 / 分钟。

## 门槛

- 需要 Cloudflare 账号、账号 ID 与 API token。
- 信用卡、手机号、地区限制：官方 Workers AI 文档未提及，注册流程需登录，待核验。

## 常见错误

| 错误 | 可能原因 | 处理方式 |
| --- | --- | --- |
| 401 / 403 | token 错误或缺少 Workers AI Read + Edit 权限 | 重新创建 token，检查权限 |
| 404 | 账号 ID 或模型 ID 错误（缺少 `@cf/` 前缀） | 对照模型目录 |
| 410（5028） | 模型已下线 | 改用现役模型，见官方下线公告 |
| 429（3036） | 已用完每日 10,000 Neurons | 等待 00:00 UTC 重置，或升级 Workers Paid |
| 429（其他） | 触发 RPM 限制 | 降低并发 |
