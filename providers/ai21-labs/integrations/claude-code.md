# Claude Code Integration

## 当前状态

待核验

## 适用前提

- 已确认该供应商 API 可用。
- 已确认是否 OpenAI-compatible。
- 已确认 Claude Code 当前版本支持自定义 base URL / provider 配置。

## 配置字段

```bash
export PROVIDER_API_KEY="your-ai21-api-key"
export PROVIDER_BASE_URL="https://api.ai21.com/studio/v1"
export PROVIDER_MODEL="jamba-large-1-7"
```

## 接入步骤

1. 获取供应商 API key。
2. 确认 base URL 和 model。
3. 按 Claude Code 当前配置方式写入 provider。
4. 运行一次最小 prompt。
5. 记录成功截图或命令输出摘要。

## 待核验

- Claude Code 配置文件路径：待核验
- provider 名称：待核验
- model 名称：待核验
- 是否支持 streaming：待核验
- 是否支持 tool use：待核验

## 验证记录

| 日期 | Claude Code 版本 | 模型 | 结果 | 备注 |
| --- | --- | --- | --- | --- |
| 2026-08-27 | 待核验 | 待核验 | 待核验 | 待核验 |
