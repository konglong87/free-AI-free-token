# Claude Code 接入 火山方舟（字节跳动 · 豆包）

## 当前状态

待核验。官方 Anthropic 兼容 Messages 端点存在，但官方 Claude Code 指南针对的是付费的 Coding Plan；普通免费额度 key 直连 Claude Code 的 `ANTHROPIC_BASE_URL` 写法未核对，未实测。

官方 Messages API 是 Anthropic 兼容格式，理论上可直连，具体 Base URL 以官方文档为准。

## 可用于配置的字段

| 字段 | 值 |
| --- | --- |
| Messages 端点 | `https://ark.cn-beijing.volces.com/api/compatible/v1/messages`（官方 curl 示例，鉴权头 `x-api-key`） |
| ANTHROPIC_BASE_URL | 待核验（官方 Claude Code 指南用的是 Coding Plan 专属端点，不适用于免费额度） |
| ANTHROPIC_AUTH_TOKEN | 本人的 ARK API Key；不可提交到仓库 |
| ANTHROPIC_MODEL | 待核验 |

## 说明与待核验

- 官方 Claude Code 指南只覆盖 Coding Plan（付费），专属端点 `https://ark.cn-beijing.volces.com/api/coding`，**不要把它当作免费方案**。
- 官方明确 Coding Plan 额度不可用于普通 API 调用；Coding Plan 的 key 与端点和普通按量 key 是两套，不要混用。
- 免费额度每模型约 50w token，Claude Code 单次会话消耗很快，建议开启安心体验模式避免超额计费。

## 验证记录

| 日期 | Claude Code 版本 | 模型 | 结果 | 备注 |
| --- | --- | --- | --- | --- |
| 2026-11-26 | 待核验 | 待核验 | 待核验 | 仅整理官方文档，未持有 key，未实测 |

## 来源

- [Messages API](https://docs.volcengine.com/docs/ark/messages-api?lang=zh)
- [Coding Plan 的 Claude Code 指南（付费）](https://docs.volcengine.com/docs/ark/coding-plan-personal-ai-claude-code?lang=zh)
- [安心体验模式](https://docs.volcengine.com/docs/ark/free-tokens-only-mode)
