# Claude Code 接入 硅基流动 SiliconFlow（国内站）

## 当前状态

官方有接入指南，**未实测**；免费模型能否稳定驱动 Claude Code 待核验。

硅基流动官方提供 Anthropic 兼容端点，Claude Code 可直连，无需 OpenAI 转换代理。

## 可用于配置的字段

| 字段 | 值 |
| --- | --- |
| ANTHROPIC_BASE_URL | `https://api.siliconflow.cn/`（官方写法，末尾无 `/v1`） |
| ANTHROPIC_API_KEY | 本人的 SiliconFlow key；不可提交到仓库 |
| ANTHROPIC_MODEL | `Qwen/Qwen3-8B`（示例，待核验） |

## 说明与待核验

- 官方页「更新日期 2025年11月17日」，并提示此场景使用将显著增加 token 消耗。
- 官方页还提供一键脚本，需要下载并执行；本目录没有执行，也不建议未审阅就运行。
- 免费模型参数规模较小、限速固定，官方提示该场景 token 消耗显著增加，可能更快触发 429，实际可用性待核验。
- 国内站才有免费模型；国际站官方 Base URL 为 `https://api.siliconflow.com/`，没有免费模型。

## 验证记录

| 日期 | Claude Code 版本 | 模型 | 结果 | 备注 |
| --- | --- | --- | --- | --- |
| 2026-11-26 | 待核验 | 待核验 | 待核验 | 仅整理官方文档，未持有 key，未实测 |

## 来源

- [官方 Claude Code 接入页](https://docs.siliconflow.cn/docs/usercases/use-siliconcloud-in-ClaudeCode)
- [Messages API](https://docs.siliconflow.cn/docs/api/messages-post)
