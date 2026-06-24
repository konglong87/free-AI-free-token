# API

## 基本信息

- 接入地址：https://openrouter.ai/api/v1
- 鉴权方式：Bearer API key
- 是否 OpenAI-compatible：是
- Chat completions path：`/chat/completions`
- Models endpoint：`/models`

## 环境变量

```bash
export PROVIDER_API_KEY="你的 API key"
export PROVIDER_BASE_URL="https://openrouter.ai/api/v1"
export PROVIDER_MODEL="cohere/north-mini-code:free"
```

## 接入步骤

1. 注册或登录 OpenRouter。
2. 在 https://openrouter.ai/settings/keys 创建 API key。
3. 从 https://openrouter.ai/models?max_price=0 选择 `:free` 免费模型。
4. 在客户端中配置接入地址、API key、模型名。
5. 发送最小对话请求验证。

## 常见错误

| 错误 | 可能原因 | 处理方式 |
| --- | --- | --- |
| 401 | API key 错误或未启用 | 重新生成 key，确认账号状态 |
| 403 | 模型不可用或账号受限 | 查看控制台和官方模型页 |
| 429 | 限流 | 降低并发，等待限制恢复 |
| 404 | endpoint 或 model 错误 | 对照官方文档核验 |

## 官方来源

- Quickstart：https://openrouter.ai/docs/quickstart
- API Overview：https://openrouter.ai/docs/api-reference/overview
- 模型接口：https://openrouter.ai/api/v1/models
