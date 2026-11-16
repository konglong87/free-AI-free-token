# K-EXAONE 2.5 API 接入

## 协议与地址

- 协议：OpenAI-compatible Chat Completions
- Base URL：`https://api.exaone.ai/v1`
- SDK 拼接路径：`/chat/completions`（完整地址 `https://api.exaone.ai/v1/chat/completions`）
- 鉴权：`Authorization: Bearer <EXAONE_API_KEY>`
- 模型 ID：`k-exaone_v2.5`

官方未提供 Anthropic Messages 或 OpenAI Responses 接口。

## 最小请求

先在自己的 shell 中设置密钥，切勿把真实值写进仓库或分享的终端记录：

```bash
export EXAONE_API_KEY="<你的 EXAONE API Key>"
curl 'https://api.exaone.ai/v1/chat/completions' \
  -H "Authorization: Bearer ${EXAONE_API_KEY}" \
  -H 'Content-Type: application/json' \
  -d '{"model":"k-exaone_v2.5","messages":[{"role":"user","content":"你好"}],"temperature":1.0,"top_p":0.95,"reasoning_effort":"low","stream":false}'
```

- 推理始终开启，推理内容在 `choices[0].message.reasoning_content`；`content` 开头可能带 `"\n\n"`，做字符串比较前先 trim。
- 流式：`"stream": true` 返回 SSE；加 `"stream_options":{"include_usage":true}` 可在最后一个 chunk 拿到用量。

## 限制与排错

- **429**：超出调用上限。响应头会给出剩余额度；有 `retry-after` / `retry-after-ms` 时按它等待，没有时退避重试，避免多 Agent 并发突发。
- **400**：先检查输入 + `max_tokens` 是否超过 524,288。
- **401 / 403**：先检查 Key 是否已过 30 天有效期，再检查拼写和鉴权头。
- 只在后端或本地 shell 使用 Key，不要放进浏览器或 App 客户端代码（Authentication 页要求）。
- 本仓库未持有 Key，以上为官方文档参数及离线示例，**尚未做真实 API 请求**。

## 官方文档

- [Quick Start](https://docs.exaone.ai/ko/docs/quick-start)
- [K-EXAONE 2.5 接口说明](https://docs.exaone.ai/ko/docs/k-exaone-2.5)
- [Rate Limits](https://docs.exaone.ai/ko/docs/rate-limits)
- [Error Codes](https://docs.exaone.ai/ko/docs/error-codes)
