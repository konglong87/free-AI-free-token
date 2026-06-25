# Hermes Integration

## 当前状态

智谱官方提供 OpenAI API 兼容接入；Hermes 专用配置待实测。

## 配置字段

```bash
export PROVIDER_API_KEY="你的智谱 API key"
export PROVIDER_BASE_URL="https://open.bigmodel.cn/api/paas/v4/"
export PROVIDER_MODEL="glm-4.7-flash"
```

## 接入步骤

1. 确认 Hermes 的 provider / model 配置格式。
2. 填写智谱 API key、Base URL、model。
3. 运行最小 prompt。

## 待核验

- Hermes 配置文件路径：待核验
- 是否支持 streaming：待核验
- 是否支持 tool calling：待核验

## 验证记录

| 日期 | Hermes 版本 | 模型 | 结果 | 备注 |
| --- | --- | --- | --- | --- |
| 2026-06-25 | 待实测 | `glm-4.7-flash` | 待核验 | 按 OpenAI-compatible 接入字段整理 |
