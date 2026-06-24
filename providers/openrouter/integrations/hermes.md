# Hermes Integration

## 当前状态

原文未提供 Hermes 专用配置，保留通用 OpenAI-compatible 接入字段，待实测。

## 配置字段

```bash
export PROVIDER_API_KEY="你的 API key"
export PROVIDER_BASE_URL="https://openrouter.ai/api/v1"
export PROVIDER_MODEL="cohere/north-mini-code:free"
```

## 接入步骤

1. 确认 Hermes 的 provider / model 配置格式。
2. 填写 OpenRouter API key、base URL、model。
3. 运行最小 prompt。
4. 验证上下文、流式输出、工具调用等能力。

## 待核验

- Hermes 配置文件路径：原文未提供
- provider 配置示例：原文未提供
- 是否支持 streaming：待核验

## 验证记录

| 日期 | Hermes 版本 | 模型 | 结果 | 备注 |
| --- | --- | --- | --- | --- |
| 2026-06-24 | 待实测 | `cohere/north-mini-code:free` | 待核验 | 原文未提供 Hermes 专用配置 |
