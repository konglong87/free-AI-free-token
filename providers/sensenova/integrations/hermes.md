# Hermes Integration

## 当前状态

SenseNova 官方 API 文档提供 OpenAI-compatible 接入；Hermes 专用配置待实测。

## 配置字段

```bash
export PROVIDER_API_KEY="你的 SenseNova API key"
export PROVIDER_BASE_URL="https://token.sensenova.cn/v1"
export PROVIDER_MODEL="sensenova-6.7-flash-lite"
```

## 接入步骤

1. 确认 Hermes 的 provider / model 配置格式。
2. 填写 SenseNova API key、Base URL、model。
3. 运行最小 prompt。

## 待核验

- Hermes 配置文件路径：待核验
- 是否支持 streaming：待核验
- 是否支持 tool calling：待核验

## 验证记录

| 日期 | Hermes 版本 | 模型 | 结果 | 备注 |
| --- | --- | --- | --- | --- |
| 2026-07-18 | 待实测 | `sensenova-6.7-flash-lite` | 待核验 | 按 OpenAI-compatible 接入字段整理 |
