# Claude Code Integration

## 当前状态

原文提供环境变量接入方式，待本地实测。

## 配置字段

```bash
export PROVIDER_API_KEY="你的 API key"
export PROVIDER_BASE_URL="https://openrouter.ai/api/v1"
export PROVIDER_MODEL="cohere/north-mini-code:free"
```

## 接入步骤

1. 获取 OpenRouter API key。
2. 设置 OpenAI-compatible 接入地址为 `https://openrouter.ai/api/v1`。
3. 设置模型为 OpenRouter 免费模型，例如 `cohere/north-mini-code:free`。
4. 使用 Claude Code 发起一次最小 prompt。

## 待核验

- Claude Code 配置文件路径：原文未提供
- 是否支持 streaming：待核验
- 是否支持 tool use：待核验

## 验证记录

| 日期 | Claude Code 版本 | 模型 | 结果 | 备注 |
| --- | --- | --- | --- | --- |
| 2026-06-24 | 待实测 | `cohere/north-mini-code:free` | 待核验 | 来源文章提供接入字段 |
