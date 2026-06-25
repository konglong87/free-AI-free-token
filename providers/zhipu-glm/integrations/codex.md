# Codex Integration

## 当前状态

智谱官方提供 OpenAI API 兼容接入；Codex 专用配置待实测。

## 配置字段

```bash
export PROVIDER_API_KEY="你的智谱 API key"
export PROVIDER_BASE_URL="https://open.bigmodel.cn/api/paas/v4/"
export PROVIDER_MODEL="glm-4.7-flash"
```

## 接入步骤

1. 确认 Codex 当前版本是否支持自定义 OpenAI-compatible provider。
2. 写入智谱 Base URL、API key、model。
3. 注意 Base URL 使用 `https://open.bigmodel.cn/api/paas/v4/`。
4. 运行最小代码问答。

## 待核验

- Codex 配置文件路径：待核验
- 是否支持 reasoning 参数：待核验
- 是否支持 tool calls：待核验

## 验证记录

| 日期 | Codex 版本 | 模型 | 结果 | 备注 |
| --- | --- | --- | --- | --- |
| 2026-06-25 | 待实测 | `glm-4.7-flash` | 待核验 | 按 OpenAI-compatible 接入字段整理 |
