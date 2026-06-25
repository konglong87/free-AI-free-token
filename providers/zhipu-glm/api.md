# API

## 基本信息

- 接入地址：https://open.bigmodel.cn/api/paas/v4/
- 鉴权方式：Bearer API key
- 是否 OpenAI-compatible：是
- Chat completions path：`/chat/completions`
- Models endpoint：待核验
- 核验状态：已整理，待实测

## 环境变量

```bash
export PROVIDER_API_KEY="你的智谱 API key"
export PROVIDER_BASE_URL="https://open.bigmodel.cn/api/paas/v4/"
export PROVIDER_MODEL="glm-4.7-flash"
```

## 接入步骤

1. 打开 https://bigmodel.cn/usercenter/proj-mgmt/apikeys。
2. 创建或复制 API key。
3. 选择免费模型，例如 `glm-4.7-flash`。
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

- Base URL：`https://open.bigmodel.cn/api/paas/v4/`
- Chat endpoint：`/chat/completions`
- API key：智谱开放平台 API key
- Model：`glm-4.7-flash`

## 官方来源

- 快速开始：https://docs.bigmodel.cn/cn/guide/start/quick-start
- Chat Completions：https://docs.bigmodel.cn/cn/api-reference/模型-api/对话补全
- OpenAI API 兼容：https://docs.bigmodel.cn/cn/guide/develop/openai
