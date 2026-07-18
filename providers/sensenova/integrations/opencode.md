# OpenCode Integration

## 当前状态

SenseNova 官方 API 文档提供 OpenAI-compatible 接入；OpenCode 专用配置待实测。

## 配置字段

```bash
export PROVIDER_API_KEY="你的 SenseNova API key"
export PROVIDER_BASE_URL="https://token.sensenova.cn/v1"
export PROVIDER_MODEL="sensenova-6.7-flash-lite"
```

## 接入步骤

1. 确认 OpenCode 当前版本是否支持 OpenAI-compatible provider。
2. 填写 SenseNova API key。
3. 填写 Base URL `https://token.sensenova.cn/v1`。
4. 填写模型名 `sensenova-6.7-flash-lite`。
5. 运行最小 prompt。

## 待核验

- OpenCode 配置文件路径：待核验
- 是否支持 streaming：待核验

## 验证记录

| 日期 | OpenCode 版本 | 模型 | 结果 | 备注 |
| --- | --- | --- | --- | --- |
| 2026-07-18 | 待实测 | `sensenova-6.7-flash-lite` | 待核验 | 按 OpenAI-compatible 接入字段整理 |
