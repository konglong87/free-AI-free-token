# Cerebras API 接入

## 协议与地址

- OpenAI 兼容 Base URL：`https://api.cerebras.ai/v1`
- 对话：`/chat/completions`（完整地址 `https://api.cerebras.ai/v1/chat/completions`；官方说明多数工具会自动追加该路径）；另有 `/completions`
- 模型列表：`GET https://api.cerebras.ai/v1/models`（`Authorization: Bearer`；官方示例只返回 `gpt-oss-120b` 与 `qwen-3.8-27b`）
- 鉴权：`Authorization: Bearer <CEREBRAS_API_KEY>`
- Anthropic Messages 端点：官方文档未见
- 来源：[OpenAI 兼容说明](https://inference-docs.cerebras.ai/resources/openai)、[List models](https://inference-docs.cerebras.ai/api-reference/models/list-models)

## 最小请求

```bash
export CEREBRAS_API_KEY="<你的 Cerebras API key>"
curl 'https://api.cerebras.ai/v1/chat/completions' \
  -H "Authorization: Bearer ${CEREBRAS_API_KEY}" \
  -H 'Content-Type: application/json' \
  -d '{"model":"gpt-oss-120b","messages":[{"role":"user","content":"你好"}]}'
```

以上为按官方文档整理的示例，**尚未用真实 key 调用**。

## 兼容性限制（官方原文要点）

- `n` 只支持 1。
- 不支持外部 HTTPS 图片 URL，图片需 base64 的 PNG 或 JPEG。
- `gpt-oss-120b` 拒绝同时带 `tools` 和 `response_format`。
- `reasoning_format`、`clear_thinking` 需通过 `extra_body` 传入。

## 门槛

- 需要绑定有效支付方式（信用卡类）才能获得 $5；官方称不会扣费，也不会自动开通付费。
- 手机号、实名、地区限制：官方页面未见（待核验）。

## 常见错误

| 错误 | 可能原因 | 处理方式 |
| --- | --- | --- |
| 401 | key 错误 | 重新生成 key |
| 403 / 额度提示 | $5 试用额度过期（30 天）或已用完 | 查看控制台额度，按需购买 PayGo |
| 404 | 模型 ID 错误或模型已弃用 | 对照模型清单与弃用页 |
| 429 | 触发 RPM / TPM / TPD 限制（`gpt-oss-120b` 试用层仅 5 RPM） | 降低并发，等待恢复 |

## 数据条款

隐私政策写到对 inference 等服务不保留输入输出（Section 6），但只读到开头片段，第 6 节全文及是否用于训练待核验（[隐私政策](https://www.cerebras.ai/privacy-policy)）。
