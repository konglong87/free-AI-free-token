# API

## 基本信息

- 接入地址：`https://api.b.ai/v1`
- 鉴权方式：Bearer API key
- 是否 OpenAI-compatible：调用示例采用 OpenAI Chat Completions 风格，完整兼容性待核验
- Chat completions path：`/v1/chat/completions`
- Models endpoint：待核验
- 核验状态：调用字段已截图，尚未使用真实 key 实测

## 环境变量

```bash
export PROVIDER_API_KEY="你的 B.AI API key"
export PROVIDER_BASE_URL="https://api.b.ai/v1"
export PROVIDER_MODEL="DeepSeek-V4-Flash"
```

## 接入步骤

1. 打开 https://chat.b.ai/key 并获取 API key。
2. 设置 `PROVIDER_BASE_URL` 和 `PROVIDER_API_KEY`。
3. 从 `DeepSeek-V4-Flash`、`DeepSeek-V4-Flash-Vision-Exp`、`Hy3` 中选择模型。
4. 按官方调用示例发送最小 Chat Completions 请求。
5. 使用真实 key 实测后补充状态、限流与能力记录。

## 调用示例证据

![B.AI API 调用示例](assets/b-ai-api-example.png)

截图中展示的请求字段包括：

- `POST https://api.b.ai/v1/chat/completions`
- `Authorization: Bearer YOUR_API_KEY`
- `Content-Type: application/json`
- `model`、`messages`、`stream`、`temperature`、`max_tokens`

## 常见错误

| 错误 | 可能原因 | 处理方式 |
| --- | --- | --- |
| 401 | API key 错误或未启用 | 重新生成 key，确认账号状态 |
| 403 | 模型不可用或账号受限 | 查看 key 页面和官方文档 |
| 429 | 限流 | 降低并发，等待限制恢复 |
| 404 | endpoint 或 model 错误 | 对照官方调用示例核验 |

## OpenAI-compatible 说明

当前仅能确认官方示例使用 Chat Completions 风格请求。以下能力尚未实测：

- streaming 行为
- tool calling
- 结构化输出
- `/models` 列表接口
