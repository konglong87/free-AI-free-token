# API

## 基本信息

- 接入地址：待填写
- 鉴权方式：Bearer API key
- 是否 OpenAI-compatible：待核验
- Chat completions path：待填写
- Models endpoint：待填写

## 环境变量

```bash
export PROVIDER_API_KEY="你的 API key"
export PROVIDER_BASE_URL="待填写"
export PROVIDER_MODEL="待填写"
```

## 接入步骤

1. 注册或登录供应商控制台。
2. 创建 API key。
3. 选择免费模型。
4. 在客户端中配置接入地址、API key、模型名。
5. 发送最小对话请求验证。

## 常见错误

| 错误 | 可能原因 | 处理方式 |
| --- | --- | --- |
| 401 | API key 错误或未启用 | 重新生成 key，确认账号状态 |
| 403 | 模型不可用或账号受限 | 查看控制台和官方文档 |
| 429 | 限流 | 降低并发，等待额度刷新 |
| 404 | endpoint 或 model 错误 | 对照官方文档核验 |

## OpenAI-compatible 说明

如果供应商兼容 OpenAI API，需要记录：

- base URL
- chat completions endpoint
- models endpoint
- 是否支持 streaming
- 是否支持 tool calling
- 是否支持结构化输出
