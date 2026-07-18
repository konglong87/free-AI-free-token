# API

## 基本信息

- 接入地址：https://token.sensenova.cn/v1
- 鉴权方式：Bearer API key
- 是否 OpenAI-compatible：是
- Chat completions path：`/chat/completions`
- Models endpoint：待核验
- 核验状态：已整理，待实测

## 环境变量

```bash
export PROVIDER_API_KEY="你的 SenseNova API key"
export PROVIDER_BASE_URL="https://token.sensenova.cn/v1"
export PROVIDER_MODEL="sensenova-6.7-flash-lite"
```

## 接入步骤

1. 打开 https://platform.sensenova.cn/docs。
2. 注册或登录平台并创建 API key。
3. 选择免费模型，例如 `sensenova-6.7-flash-lite`。
4. 在客户端中配置接入地址、API key、模型名。
5. 发送最小对话请求验证。

## 常见错误

| 错误 | 可能原因 | 处理方式 |
| --- | --- | --- |
| 401 | API key 错误或未启用 | 重新生成 key，确认账号状态 |
| 403 | 模型不可用或账号受限 | 查看控制台和官方文档 |
| 429 | 限流 | 降低并发，等待限制恢复 |
| 404 | endpoint 或 model 错误 | 对照官方文档核验 |

## OpenAI-compatible 说明

如果客户端要求填写 OpenAI-compatible provider，通常填写：

- Base URL：`https://token.sensenova.cn/v1`
- Chat endpoint：`/chat/completions`
- API key：SenseNova API key
- Model：`sensenova-6.7-flash-lite`

## 官方来源

- Token Plan：https://www.sensenova.cn/token-plan
- API 文档：https://github.com/OpenSenseNova/SenseNova6.7/blob/main/API_CN.md
