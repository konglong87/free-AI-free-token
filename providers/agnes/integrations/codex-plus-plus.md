# Codex++ Integration

## 当前状态

官网常见接入文档列出 Codex++ 接入方法，待本地实测。

## 配置字段

```bash
export PROVIDER_API_KEY="你的 Agnes AI API key"
export PROVIDER_BASE_URL="https://apihub.agnes-ai.com/v1"
export PROVIDER_MODEL="agnes-2.0-flash"
```

## 接入步骤

1. 打开 Agnes AI 常见接入文档：https://agnes-ai.com/doc/cid1。
2. 选择 Codex++ 接入方法。
3. 配置 Base URL 为 `https://apihub.agnes-ai.com/v1`。
4. 填写 Agnes AI API key 和模型名。
5. 注意 Base URL 不要拼上 `/chat/completions`。
6. 运行最小代码问答。

## 待核验

- Codex++ 配置文件路径：待核验
- 是否支持 reasoning 参数：待核验
- 是否支持 tool calls：待核验

## 验证记录

| 日期 | Codex++ 版本 | 模型 | 结果 | 备注 |
| --- | --- | --- | --- | --- |
| 2026-06-24 | 待实测 | `agnes-2.0-flash` | 待核验 | 官网常见接入文档列出 Codex++ |
