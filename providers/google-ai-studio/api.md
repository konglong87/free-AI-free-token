# Google AI Studio API 接入

## 协议与地址

- OpenAI 兼容 Base URL：`https://generativelanguage.googleapis.com/v1beta/openai/`（官方原文，注意末尾斜杠）
- 对话：`chat/completions`，官方写明支持流式、函数调用、图片与音频输入、结构化输出；另有 `images/generations`、embeddings、batch、模型列表（[OpenAI 兼容](https://aistudio.google.com/docs/openai)）
- 文件上传与下载走原生 `genai` 客户端，不走兼容层
- 鉴权：API key；[API key 说明](https://aistudio.google.com/docs/api-key)
- Anthropic Messages 端点：官方文档未见（针对性检索也未找到官方页面）
- 示例模型 ID（官方兼容页）：`gemini-4-hydrogen`、`gemini-4-neon`（免费层不可用）、`gemini-3.8-flash`、`gemini-embedding-2`、`gemini-embedding-001`

## 最小请求

```bash
export GEMINI_API_KEY="<你的 Gemini API key>"
curl 'https://generativelanguage.googleapis.com/v1beta/openai/chat/completions' \
  -H "Authorization: Bearer ${GEMINI_API_KEY}" \
  -H 'Content-Type: application/json' \
  -d '{"model":"gemini-3.8-flash","messages":[{"role":"user","content":"你好"}]}'
```

以上为按官方 OpenAI 兼容页整理的示例，**尚未用真实 key 调用**；`gemini-4-neon` 等免费层不可用的模型会失败。

## 限流与门槛

见 [models.md](models.md) 的限流、数据条款与门槛。要点：按项目限流；免费层限额需登录 AI Studio 查看；第二个及以后免费项目需设备验证。

## 常见错误

| 错误 | 可能原因 | 处理方式 |
| --- | --- | --- |
| 400 / 404 | 模型 ID 错误，或模型已弃用 / 停服 | 对照模型清单与弃用页 |
| 403 `FREE_TIER_VERIFICATION_REQUIRED` | 第二个及以后的免费项目未做设备验证 | 按官方项目验证页完成扫码验证 |
| 403（其他） | 免费层不可用的模型（如 `gemini-4-neon`、图片生成类） | 改用免费层模型，或升级付费层 |
| 429 | 触发项目限额 | 降低并发；免费层使用可被抢占容量 |
