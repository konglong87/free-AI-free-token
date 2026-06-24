# Claude Code Integration

## 当前状态

官网常见接入文档列出 Claude CLI 和 Claude Desktop；Claude Code 专用配置待实测。

## 配置字段

```bash
export PROVIDER_API_KEY="你的 Agnes AI API key"
export PROVIDER_BASE_URL="https://apihub.agnes-ai.com/v1"
export PROVIDER_MODEL="agnes-2.0-flash"
```

## 接入步骤

1. 打开 Agnes AI 常见接入文档：https://agnes-ai.com/doc/cid1。
2. 参考 Claude CLI 或 Claude Desktop 接入方法。
3. 确认 Claude Code 是否支持 OpenAI-compatible provider。
4. 设置 Base URL 为 `https://apihub.agnes-ai.com/v1`。
5. 设置模型为 `agnes-2.0-flash`。
6. 发起一次最小 prompt。

## 待核验

- Claude Code 配置文件路径：待核验
- Claude Code 是否可直接使用 Agnes AI OpenAI-compatible 接口：待核验
- 是否支持 streaming：待核验
- 是否支持 tool use：待核验

## 验证记录

| 日期 | Claude Code 版本 | 模型 | 结果 | 备注 |
| --- | --- | --- | --- | --- |
| 2026-06-24 | 待实测 | `agnes-2.0-flash` | 待核验 | 官网常见接入文档列出 Claude CLI 和 Claude Desktop |
