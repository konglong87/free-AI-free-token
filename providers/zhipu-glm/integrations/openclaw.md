# OpenClaw Integration

## 当前状态

智谱官方提供 OpenAI API 兼容接入；OpenClaw 专用配置待实测。

## 配置字段

```bash
export PROVIDER_API_KEY="你的智谱 API key"
export PROVIDER_BASE_URL="https://open.bigmodel.cn/api/paas/v4/"
export PROVIDER_MODEL="glm-4.7-flash"
```

## 接入步骤

1. 在 OpenClaw 中选择 OpenAI-compatible provider。
2. 填写智谱 API key。
3. 填写 Base URL `https://open.bigmodel.cn/api/paas/v4/`。
4. 填写模型名 `glm-4.7-flash`。
5. 运行最小对话请求。

## 待核验

- OpenClaw 配置文件路径：待核验
- 是否支持 streaming：待核验
- 是否支持 tool calling：待核验

## 验证记录

| 日期 | OpenClaw 版本 | 模型 | 结果 | 备注 |
| --- | --- | --- | --- | --- |
| 2026-06-25 | 待实测 | `glm-4.7-flash` | 待核验 | 按 OpenAI-compatible 接入字段整理 |
