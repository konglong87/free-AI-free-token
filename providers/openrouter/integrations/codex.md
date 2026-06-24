# Codex Integration

## 当前状态

原文未提供 Codex 专用配置，保留通用 OpenAI-compatible 接入字段，待实测。

## 配置字段

```bash
export PROVIDER_API_KEY="你的 API key"
export PROVIDER_BASE_URL="https://openrouter.ai/api/v1"
export PROVIDER_MODEL="cohere/north-mini-code:free"
```

## 接入步骤

1. 确认 Codex 当前版本是否支持自定义 OpenAI-compatible provider。
2. 写入 OpenRouter base URL、API key、model。
3. 运行最小代码问答。
4. 记录成功命令和输出摘要。

## 待核验

- Codex 配置文件路径：原文未提供
- provider 配置示例：原文未提供
- 是否支持 reasoning 参数：待核验
- 是否支持 tool calls：待核验

## 验证记录

| 日期 | Codex 版本 | 模型 | 结果 | 备注 |
| --- | --- | --- | --- | --- |
| 2026-06-24 | 待实测 | `cohere/north-mini-code:free` | 待核验 | 原文未提供 Codex 专用配置 |
