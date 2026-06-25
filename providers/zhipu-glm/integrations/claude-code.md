# Claude Code Integration

## 当前状态

智谱官方提供 OpenAI API 兼容接入；Claude Code 专用配置待实测。

## 配置字段

```bash
export PROVIDER_API_KEY="你的智谱 API key"
export PROVIDER_BASE_URL="https://open.bigmodel.cn/api/paas/v4/"
export PROVIDER_MODEL="glm-4.7-flash"
```

## 接入步骤

1. 确认 Claude Code 当前版本是否支持自定义 OpenAI-compatible provider。
2. 设置 Base URL 为 `https://open.bigmodel.cn/api/paas/v4/`。
3. 设置 API key 为智谱开放平台 API key。
4. 设置模型为 `glm-4.7-flash`。
5. 发起一次最小 prompt。

## 待核验

- Claude Code 配置文件路径：待核验
- 是否支持 streaming：待核验
- 是否支持 tool use：待核验

## 验证记录

| 日期 | Claude Code 版本 | 模型 | 结果 | 备注 |
| --- | --- | --- | --- | --- |
| 2026-06-25 | 待实测 | `glm-4.7-flash` | 待核验 | 按 OpenAI-compatible 接入字段整理 |
