# OpenClaw Integration

## 当前状态

原文提供 OpenAI provider 接入字段，待本地实测。

## 配置字段

```bash
export PROVIDER_API_KEY="你的 API key"
export PROVIDER_BASE_URL="https://openrouter.ai/api/v1"
export PROVIDER_MODEL="cohere/north-mini-code:free"
```

## 接入步骤

1. 在 OpenClaw 中选择 OpenAI-compatible provider。
2. 填写 OpenRouter API key。
3. 填写接入地址 `https://openrouter.ai/api/v1`。
4. 填写免费模型名，例如 `cohere/north-mini-code:free`。
5. 运行最小对话请求。

## 待核验

- OpenClaw 配置文件路径：原文未提供
- 是否支持 streaming：待核验
- 是否支持 tool calling：待核验

## 验证记录

| 日期 | OpenClaw 版本 | 模型 | 结果 | 备注 |
| --- | --- | --- | --- | --- |
| 2026-06-24 | 待实测 | `cohere/north-mini-code:free` | 待核验 | 来源文章提供接入字段 |
