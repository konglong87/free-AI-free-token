# 官方链接与核验记录

## 官方入口

- [官网](https://ai.google.dev/)（跳转到 https://ai.google/；旧路径 `ai.google.dev/gemini-api/docs/*` 均 301 到下方文档）
- [AI Studio 控制台](https://aistudio.google.com)
- [API key 说明](https://aistudio.google.com/docs/api-key)（创建页 `aistudio.google.com/apikey` 需登录，未打开）
- [文档首页](https://aistudio.google.com/docs)、[全站索引](https://aistudio.google.com/llms.txt)
- [模型清单](https://aistudio.google.com/docs/models)
- [定价与免费层](https://aistudio.google.com/docs/pricing)
- [限流](https://aistudio.google.com/docs/rate-limits)（个人限额页 `aistudio.google.com/rate-limit` 需登录）
- [计费](https://aistudio.google.com/docs/billing)
- [弃用](https://aistudio.google.com/docs/deprecations)
- [OpenAI 兼容](https://aistudio.google.com/docs/openai)
- [项目验证](https://aistudio.google.com/docs/verification)
- [可用地区](https://aistudio.google.com/docs/available-regions)
- [服务条款（含 Unpaid / Paid Services 数据条款）](https://aistudio.google.com/docs/terms)
- [状态页](https://aistudio.google.com/status)
- [Coding agents 页（让编码代理读 Gemini 文档的 MCP / Skills）](https://aistudio.google.com/docs/coding-agents)

## 来源核验记录

| 日期 | 来源 | 核验内容 | 结果 |
| --- | --- | --- | --- |
| 2026-11-26 | [计费](https://aistudio.google.com/docs/billing)、[定价](https://aistudio.google.com/docs/pricing) | 新账号默认 Free Tier；定价页开头写「可免费开始」；逐模型 Free Tier 列标注 Free of charge 或 Not available | 已核对 |
| 2026-11-26 | [限流](https://aistudio.google.com/docs/rate-limits) | 限流按项目计，不按 key；只公布了免费层 `gemini-decision-1` 的数值；其余写「在 AI Studio 查看」；写明限额不保证 | 已核对；其余数值需登录，待核验 |
| 2026-11-26 | [模型清单](https://aistudio.google.com/docs/models) | Gemini 4 系列输入 1M tokens，输出 64K 或 1M，支持文本、图片、音频、视频、PDF 输入；其他文本模型页面未写 token 上限 | 已核对 |
| 2026-11-26 | [OpenAI 兼容](https://aistudio.google.com/docs/openai) | Base URL `https://generativelanguage.googleapis.com/v1beta/openai/`；chat completions（流式、函数调用、图片与音频输入、结构化输出）、`images/generations`、embeddings、batch、模型列表 | 已核对，未实测 |
| 2026-11-26 | [项目验证](https://aistudio.google.com/docs/verification) | 第二个及以后免费项目、或 AI Studio 之外创建 / 导入的项目需扫码设备验证，官方称不需要手机号；未验证返回 403 `FREE_TIER_VERIFICATION_REQUIRED` | 已核对 |
| 2026-11-26 | [服务条款](https://aistudio.google.com/docs/terms) | Unpaid Services 内容可用于改进产品并可能被人工审核；EEA、瑞士、英国适用付费级数据条款 | 已核对 |
| 2026-11-26 | [弃用](https://aistudio.google.com/docs/deprecations) | `gemini-2.5-pro` / `gemini-2.5-flash` 2026-12-01 起自动路由；其他模型停服日期 | 已核对（日期为最早可能日期） |
| 2026-11-26 | [状态页](https://aistudio.google.com/status) | 页面写明免费层请求使用可被抢占的容量 | 已核对 |
| 2026-11-26 | 官方文档检索 | Anthropic 兼容端点 | 未见 |
