# Codex Integration

## 当前状态

SenseNova 官方 API 文档提供 OpenAI-compatible 接入；Codex 专用配置待实测。

## 配置字段

```bash
export PROVIDER_API_KEY="你的 SenseNova API key"
export PROVIDER_BASE_URL="https://token.sensenova.cn/v1"
export PROVIDER_MODEL="sensenova-6.7-flash-lite"
```

## 接入步骤

1. 确认 Codex 当前版本是否支持自定义 OpenAI-compatible provider。
2. 写入 SenseNova Base URL、API key、model。
3. 注意 Base URL 使用 `https://token.sensenova.cn/v1`。
4. 运行最小代码问答。

## 待核验

- Codex 配置文件路径：待核验
- 是否支持 reasoning 参数：待核验
- 是否支持 tool calls：待核验

## 验证记录

| 日期 | Codex 版本 | 模型 | 结果 | 备注 |
| --- | --- | --- | --- | --- |
| 2026-07-18 | 待实测 | `sensenova-6.7-flash-lite` | 待核验 | 按 OpenAI-compatible 接入字段整理 |
