# Claude Code Integration

## 当前状态

SenseNova 官方 API 文档提供 OpenAI-compatible 接入；Claude Code 专用配置待实测。

## 配置字段

```bash
export PROVIDER_API_KEY="你的 SenseNova API key"
export PROVIDER_BASE_URL="https://token.sensenova.cn/v1"
export PROVIDER_MODEL="sensenova-6.7-flash-lite"
```

## 接入步骤

1. 确认 Claude Code 当前版本是否支持自定义 OpenAI-compatible provider。
2. 设置 Base URL 为 `https://token.sensenova.cn/v1`。
3. 设置 API key 为 SenseNova API key。
4. 设置模型为 `sensenova-6.7-flash-lite`。
5. 发起一次最小 prompt。

## 待核验

- Claude Code 配置文件路径：待核验
- 是否支持 streaming：待核验
- 是否支持 tool use：待核验

## 验证记录

| 日期 | Claude Code 版本 | 模型 | 结果 | 备注 |
| --- | --- | --- | --- | --- |
| 2026-07-18 | 待实测 | `sensenova-6.7-flash-lite` | 待核验 | 按 OpenAI-compatible 接入字段整理 |
