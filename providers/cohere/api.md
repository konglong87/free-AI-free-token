# Cohere API 接入

## 协议与地址

- OpenAI 兼容 Base URL：`https://api.cohere.ai/compatibility/v1`（官方 Python 示例）；音频转写示例使用 `api.cohere.com`
- 对话：`/chat/completions`，支持流式、结构化输出、工具调用；另有 embeddings、audio transcriptions（[OpenAI 兼容页](https://docs.cohere.com/docs/compatibility-api)）
- 模型列表：兼容页未提 `/models`
- 鉴权：Cohere API key 作为 Bearer
- Anthropic Messages 端点：官方文档未见（文档的 MCP server `https://docs.cohere.com/_mcp/server` 只供编码代理读文档，不是模型接入端点）

## 最小请求

```bash
export COHERE_API_KEY="<你的 Cohere Trial key>"
curl 'https://api.cohere.ai/compatibility/v1/chat/completions' \
  -H "Authorization: Bearer ${COHERE_API_KEY}" \
  -H 'Content-Type: application/json' \
  -d '{"model":"command-a-03-2025","messages":[{"role":"user","content":"你好"}]}'
```

以上为按官方文档整理的示例，**尚未用真实 key 调用**。

## 兼容性限制（官方原文要点）

- 不支持的 chat 参数：`store`、`metadata`、`logit_bias`、`top_logprobs`、`n`、`modalities`、`prediction`、`audio`、`service_tier`、`parallel_tool_calls`。
- `reasoning_effort` 只接受 `none` 和 `high`。

## 限流与门槛

见 [models.md](models.md)。要点：Trial key 每月 1,000 次、Chat 20 请求 / 分钟；不得用于生产 / 商业用途。信用卡、手机号、地区限制官方文档未提（待核验）。

## 常见错误

| 错误 | 可能原因 | 处理方式 |
| --- | --- | --- |
| 401 | key 错误 | 重新生成 key |
| 403 / 额度提示 | Trial key 每月 1,000 次已用完，或用于不被允许的场景 | 等待下月额度，或申请 Production key |
| 404 | 模型 ID 错误或模型已弃用 / 退役 | 对照官方模型列表 |
| 429 | 触发 20 请求 / 分钟等限流 | 降低并发 |
