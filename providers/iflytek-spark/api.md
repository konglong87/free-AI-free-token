# Spark-X2.6 API 接入

## 协议与地址

| 客户端协议 | Base URL（填入客户端） | 请求路径（SDK 自动拼接） | 鉴权 |
| --- | --- | --- | --- |
| OpenAI Chat Completions | `https://maas-api.cn-huabei-1.xf-yun.com/v2` | `/chat/completions` | `Authorization: Bearer <APIKEY>` |
| OpenAI Responses | `https://maas-api.cn-huabei-1.xf-yun.com/v1` | `/responses` | `Authorization: Bearer <APIKEY>` |
| Anthropic Messages / Claude Code | `https://maas-api.cn-huabei-1.xf-yun.com/anthropic` | `/v1/messages` | `x-api-key: <APIKEY>`（Claude Code 用 `ANTHROPIC_AUTH_TOKEN`） |

不要把三个 Base URL 混用，也不要在 SDK 的 Base URL 里额外填 `/chat/completions`、`/responses` 或 `/v1/messages`。模型 ID：`spark-x2.6`。API Key 可在 MaaS 控制台的「API调用 / APIKEY管理」获取。

## 最小请求（OpenAI Chat）

先在自己的 shell 中设置密钥，切勿粘贴真实值到仓库或分享的终端记录：

```bash
export SPARK_API_PASSWORD="<你的 MaaS APIKEY>"
curl 'https://maas-api.cn-huabei-1.xf-yun.com/v2/chat/completions' \
  -H "Authorization: Bearer ${SPARK_API_PASSWORD}" \
  -H 'Content-Type: application/json' \
  -d '{"model":"spark-x2.6","messages":[{"role":"user","content":"你好"}],"stream":false}'
```

Claude Code 使用另一套协议及 Base URL，见 [Claude Code 接入](integrations/claude-code.md)。Codex 用 Responses 而非 Chat Completions，见 [Codex 接入](integrations/codex.md)。

## 限制与排错

- 赠额与实际可用模型、API Key 权限以控制台为准；未领取不能假定免费。
- 官方限流维度为 API Key + 模型下的并发和 TPM。出现 429 时降低并发、排队并退避重试，避免多 Agent 同时突发请求；当前默认值及提升办法见[速率限制文档](https://maas.xfyun.cn/doc/guide/1%E3%80%81%E5%BF%AB%E9%80%9F%E5%BC%80%E5%A7%8B/1.3%20%E9%80%9F%E7%8E%87%E9%99%90%E5%88%B6.html)。
- 401/403 先检查 Key 与权限，404 先检查模型 ID 和接口路径。Anthropic 接口文档另列单请求 Token 限制；不要把模型宣传的 1M 当作所有协议的无条件可用值。
- 本仓库未持有 Key，以上为官方文档参数及离线示例，**尚未做真实 API 请求**。

## 官方文档

- [MaaS API 接入](https://maas.xfyun.cn/doc/guide/1%E3%80%81%E5%BF%AB%E9%80%9F%E5%BC%80%E5%A7%8B/1.1%20API%E6%8E%A5%E5%85%A5.html)
- [Anthropic Messages](https://maas.xfyun.cn/doc/guide/3%E3%80%81API%20%E6%8E%A5%E5%8F%A3/3.1%20%E6%98%9F%E7%81%AB%E6%A8%A1%E5%9E%8B%20API/3.1.3%20Anthropic.html)
