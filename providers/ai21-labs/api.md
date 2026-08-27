# API

## 基本信息

- 接入地址：`https://api.ai21.com/studio/v1`
- 鉴权方式：Bearer API key
- 是否 OpenAI-compatible：待核验
- Chat completions path：待核验（先以官方 Jamba 文档为准）
- Models endpoint：待核验
- 核验状态：Key 入口与 Base URL 已记录，尚未使用真实 key 实测

## 环境变量

```bash
export PROVIDER_API_KEY="你的 AI21 Labs API key"
export PROVIDER_BASE_URL="https://api.ai21.com/studio/v1"
export PROVIDER_MODEL="jamba-large-1-7"
```

## 接入步骤

1. 在 https://studio.ai21.com/account/api-key 创建 API key。
2. 设置 `PROVIDER_BASE_URL` 和 `PROVIDER_API_KEY`。
3. 从 [模型列表](models.md) 选择模型。
4. 对照 AI21 官方 Jamba 文档确认请求 path 和字段。
5. 使用真实 key 验证响应格式、streaming 与错误码。

## 常见错误

| 错误 | 可能原因 | 处理方式 |
| --- | --- | --- |
| 401 | API key 错误或未启用 | 重新生成 key，确认账号状态 |
| 403 | 模型不可用或账号受限 | 查看 Studio 权限与套餐 |
| 429 | 限流 | 降低并发，等待限制恢复 |
| 404 | endpoint 或 model 错误 | 对照 Jamba 官方文档核验 |

## OpenAI-compatible 说明

当前无法从公开参考清单确认完整 OpenAI 兼容性，以下能力均待核验：

- Chat Completions path 与 `/models` 接口
- streaming 行为
- tool calling
- 结构化输出
