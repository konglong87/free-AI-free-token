# API

## 基本信息

- 接入地址：`https://api-inference.modelscope.cn/v1`
- 鉴权方式：Bearer Access Token
- 是否 OpenAI-compatible：参考入口采用 OpenAI Chat Completions 风格，完整兼容性待核验
- Chat completions path：`/chat/completions`
- Models endpoint：待核验
- 核验状态：入口已记录，尚未使用真实 Token 实测

## 环境变量

```bash
export PROVIDER_API_KEY="你的 ModelScope Access Token"
export PROVIDER_BASE_URL="https://api-inference.modelscope.cn/v1"
export PROVIDER_MODEL="MiniMax/MiniMax-M2.5"
```

## 接入步骤

1. 在 https://modelscope.cn/my/myaccesstoken 创建 Access Token。
2. 设置 `PROVIDER_BASE_URL` 和 `PROVIDER_API_KEY`。
3. 从 [模型列表](models.md) 选择模型。
4. 发送最小 Chat Completions 请求验证。
5. 记录实际限流、streaming、tool calling 和结构化输出结果。

## 常见错误

| 错误 | 可能原因 | 处理方式 |
| --- | --- | --- |
| 401 | Token 错误或未启用 | 重新生成 Token，确认账号状态 |
| 403 | 模型不可用或账号受限 | 查看模型页面与权限 |
| 429 | 限流 | 降低并发，等待限制恢复 |
| 404 | endpoint 或 model 错误 | 对照 ModelScope 文档核验 |

## OpenAI-compatible 说明

参考入口使用 Chat Completions 风格。以下能力尚未实测：

- `/models` 列表接口
- streaming 行为
- tool calling
- 结构化输出
