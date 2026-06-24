# Codex Integration

## 当前状态

官网常见接入文档列出 Codex++ 接入方法；Codex 通用接入按 OpenAI-compatible 字段整理，待本地实测。

## 配置字段

```bash
export PROVIDER_API_KEY="你的 Agnes AI API key"
export PROVIDER_BASE_URL="https://apihub.agnes-ai.com/v1"
export PROVIDER_MODEL="agnes-2.0-flash"
```

## 接入步骤

1. 打开 Agnes AI 常见接入文档：https://agnes-ai.com/doc/cid1。
2. 参考 Codex++ 接入方法。
3. 确认 Codex 当前版本是否支持自定义 OpenAI-compatible provider。
4. 写入 Agnes AI Base URL、API key、model。
5. 注意 Base URL 使用 `https://apihub.agnes-ai.com/v1`，不要拼上 `/chat/completions`。
6. 运行最小代码问答。

## 待核验

- Codex 配置文件路径：待核验
- 是否支持 reasoning 参数：待核验
- 是否支持 tool calls：待核验

## 验证记录

| 日期 | Codex 版本 | 模型 | 结果 | 备注 |
| --- | --- | --- | --- | --- |
| 2026-06-24 | 待实测 | `agnes-2.0-flash` | 待核验 | 官网常见接入文档列出 Codex++ |
