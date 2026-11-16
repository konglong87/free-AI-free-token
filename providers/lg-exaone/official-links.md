# 官方链接与核验记录

## 官方入口

- [LG AI Research 官方公告：K-EXAONE 2.5 API Available for Free（2026-11-09）](https://www.lgresearch.ai/news/view?seq=755)
- [EXAONE API Platform（申请 API Key）](https://docs.exaone.ai)
- [Quick Start](https://docs.exaone.ai/ko/docs/quick-start)
- [Authentication](https://docs.exaone.ai/ko/docs/authentication)
- [Models](https://docs.exaone.ai/ko/docs/models)
- [K-EXAONE 2.5 接口说明](https://docs.exaone.ai/ko/docs/k-exaone-2.5)
- [Rate Limits](https://docs.exaone.ai/ko/docs/rate-limits)
- [Agent로 시작하기（Hermes / OpenCode 接入）](https://docs.exaone.ai/ko/docs/agent-start)
- [Error Codes](https://docs.exaone.ai/ko/docs/error-codes)
- [服务条款（이용약관）](https://docs.exaone.ai/ko/terms/terms-of-service)
- [隐私政策](https://docs.exaone.ai/ko/terms/privacy-policy)

## 易混淆

- LG ThinQ API（`connect-pat.lgthinq.com`）是智能家电控制平台，与大模型无关，不收录。

## 核验记录

| 日期 | 内容 | 结论 |
| --- | --- | --- |
| 2026-11-16 | LG AI Research 官方公告 | 确认 2026-11-09 起免费开放，入口为 docs.exaone.ai；官方称其为技术验证模型；未写免费截止日期 |
| 2026-11-16 | 服务条款 | 第 2 条第 1 项：非商业使用权限，不得营利或让第三方经自己的产品调用模型；第 3 条第 4 项：Key 自签发日起 30 天有效，到期自动停用；第 2 条第 3 项：输入输出默认可用于训练，可拒绝；第 9 条第 2 项：转为收费时会在网站公告 |
| 2026-11-16 | Quick Start / K-EXAONE 2.5 / Models 页 | 确认 Base URL、`/v1/chat/completions`、`k-exaone_v2.5`、512K 上下文、仅文本输入、推理参数 |
| 2026-11-16 | Rate Limits 页 | 超限返回 429；同意可选的训练用途条款可提高每日调用上限；未公开具体限额 |
| 2026-11-16 | Agent 接入页 | 官方只给出 Hermes 与 OpenCode 配置；未提及 Claude Code、Codex、OpenClaw |
