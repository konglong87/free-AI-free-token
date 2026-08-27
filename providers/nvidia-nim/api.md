# API

## 基本信息

- 接入地址：`https://integrate.api.nvidia.com/v1`
- 鉴权方式：Bearer API key
- 是否 OpenAI-compatible：官方示例采用 OpenAI Chat Completions 风格，完整兼容性待核验
- Chat completions path：`/chat/completions`
- Models endpoint：待核验
- 核验状态：入口已记录，尚未使用真实 key 实测

## 环境变量

```bash
export PROVIDER_API_KEY="你的 NVIDIA API key"
export PROVIDER_BASE_URL="https://integrate.api.nvidia.com/v1"
export PROVIDER_MODEL="z-ai/glm-5.2"
```

## 接入步骤

1. 在 https://build.nvidia.com/settings/api-keys 创建 API key。
2. 设置 `PROVIDER_BASE_URL` 和 `PROVIDER_API_KEY`。
3. 从 [模型列表](models.md) 选择模型。
4. 发送最小 Chat Completions 请求验证。
5. 记录限流、streaming、tool calling 和结构化输出结果。

## 常见错误

| 错误 | 可能原因 | 处理方式 |
| --- | --- | --- |
| 401 | API key 错误或未启用 | 重新生成 key，确认账号状态 |
| 403 | 模型不可用或账号受限 | 查看 Build 页面与模型权限 |
| 429 | 限流 | 降低并发，等待限制恢复 |
| 404 | endpoint 或 model 错误 | 对照 NIM 文档核验 |

## OpenAI-compatible 说明

官方入口采用 Chat Completions 风格。以下能力尚未实测：

- `/models` 列表接口
- streaming 行为
- tool calling
- 结构化输出
