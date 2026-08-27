# API

## 基本信息

- 接入地址：`https://api.groq.com/openai/v1`
- 鉴权方式：Bearer API key
- 是否 OpenAI-compatible：是（官方兼容入口）；完整能力待核验
- Chat completions path：`/chat/completions`
- Models endpoint：`/models`（需 key，待实测）
- 核验状态：官方入口已记录，尚未使用真实 key 实测

## 环境变量

```bash
export PROVIDER_API_KEY="你的 Groq API key"
export PROVIDER_BASE_URL="https://api.groq.com/openai/v1"
export PROVIDER_MODEL="moonshotai/kimi-k2-instruct"
```

## 接入步骤

1. 在 https://console.groq.com/keys 创建 API key。
2. 设置 `PROVIDER_BASE_URL` 和 `PROVIDER_API_KEY`。
3. 从 [模型列表](models.md) 选择模型。
4. 发送最小 Chat Completions 请求验证。
5. 记录实际 RPM/RPD、streaming、tool calling 和结构化输出结果。

## 常见错误

| 错误 | 可能原因 | 处理方式 |
| --- | --- | --- |
| 401 | API key 错误或未启用 | 重新生成 key，确认账号状态 |
| 403 | 模型不可用或账号受限 | 查看控制台与模型权限 |
| 429 | 免费层限流 | 降低并发，等待限制恢复 |
| 404 | endpoint 或 model 错误 | 对照 Groq 模型文档核验 |

## OpenAI-compatible 说明

Groq 提供 OpenAI-compatible Base URL。以下能力仍需真实 key 验证：

- `/models` 返回字段
- streaming 行为
- tool calling
- 结构化输出
