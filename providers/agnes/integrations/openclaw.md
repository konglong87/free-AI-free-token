# OpenClaw Integration

## 当前状态

官网常见接入文档列出 OpenClaw 接入方法，待本地实测。

## 配置字段

```bash
export PROVIDER_API_KEY="你的 Agnes AI API key"
export PROVIDER_BASE_URL="https://apihub.agnes-ai.com/v1"
export PROVIDER_MODEL="agnes-2.0-flash"
```

## 接入步骤

1. 打开 Agnes AI 常见接入文档：https://agnes-ai.com/doc/cid1。
2. 选择 OpenClaw 接入方法。
3. 在 OpenClaw 中选择 OpenAI-compatible provider。
4. 填写 Agnes AI API key。
5. 填写 Base URL `https://apihub.agnes-ai.com/v1`。
6. 填写模型名 `agnes-2.0-flash`。
7. 运行最小对话请求。

## 待核验

- OpenClaw 配置文件路径：待核验
- 是否支持 streaming：待核验
- 是否支持 tool calling：待核验

## 验证记录

| 日期 | OpenClaw 版本 | 模型 | 结果 | 备注 |
| --- | --- | --- | --- | --- |
| 2026-06-24 | 待实测 | `agnes-2.0-flash` | 待核验 | 官网提供 OpenClaw 文档入口 |
