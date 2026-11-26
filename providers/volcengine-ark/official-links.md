# 官方链接与核验记录

## 官方入口

- [官网产品页](https://www.volcengine.com/product/ark)（JS 渲染，本次未读到正文）
- [方舟控制台 API Key 页](https://ark.volcengine.com/region:cn-beijing/apikey)（需登录；链接取自官方快速入门页，本次未打开）
- [文档中心](https://docs.volcengine.com/docs/ark/)
- [快速入门](https://docs.volcengine.com/docs/ark/quick-start?lang=zh)
- [Base URL 及鉴权](https://docs.volcengine.com/docs/ark/base-url-and-authentication?lang=zh)
- [模型列表](https://docs.volcengine.com/docs/ark/model-list?lang=zh)
- [免费推理额度](https://docs.volcengine.com/docs/ark/free-inference-quota)
- [安心体验模式](https://docs.volcengine.com/docs/ark/free-tokens-only-mode)
- [开通管理](https://docs.volcengine.com/docs/ark/activation-management?lang=zh)
- [协作奖励计划](https://docs.volcengine.com/docs/ark/collaboration-reward-program?lang=zh)
- [Messages API（Anthropic 兼容）](https://docs.volcengine.com/docs/ark/messages-api?lang=zh)
- [Coding Plan 概览（付费套餐）](https://docs.volcengine.com/docs/ark/coding-plan-personal-plan-overview)
- 状态页：官方文档未见
- 服务条款：搜索结果出现过 `https://www.volcengine.com/docs/ark/volcengine-ark-platform-terms`，本次未打开，待核验

## 来源核验记录

| 日期 | 来源 | 核验内容 | 结果 |
| --- | --- | --- | --- |
| 2026-11-26 | [免费推理额度](https://docs.volcengine.com/docs/ark/free-inference-quota)（页面更新 2026-11-24） | 注册火山引擎账号获得免费额度；按模型分别计算，主账号共享；只抵扣按 token 后付费的在线推理，不抵扣批量推理、插件、知识库、上下文缓存存储；额度耗尽后调用失败，继续使用需实名认证并开通 | 已核对 |
| 2026-11-26 | [安心体验模式](https://docs.volcengine.com/docs/ark/free-tokens-only-mode)（更新 2026-11-24） | 只消耗平台赠送的 50w token，接近额度时暂停，返回 HTTP 429 `SetLimitExceeded`；默认关闭，关闭后不可再开启；多模态向量化、同声传译及两个第三方模型不支持 | 已核对 |
| 2026-11-26 | [开通管理](https://docs.volcengine.com/docs/ark/activation-management?lang=zh)（更新 2026-11-24） | 「每个模型提供一定额度的免费调用量」；API 调用需开通模型服务；控制台额度显示有小时级延迟；可设用量上限；模型 RPM / TPM 限制按账号、按基础模型、接入点共享 | 已核对 |
| 2026-11-26 | [协作奖励计划](https://docs.volcengine.com/docs/ark/collaboration-reward-program?lang=zh)（更新 2026-11-24） | 授权数据换额度；每日奖励包 30 个自然日有效；多数模型已于 2026-09-01 / 2026-10-08 退出；活动页入口目前列出 Doubao-Seed-Evolving | 已核对；仍在计划内的完整名单需登录「产品活动介绍页」，待核验 |
| 2026-11-26 | [模型列表](https://docs.volcengine.com/docs/ark/model-list?lang=zh) | 模型 ID、上下文、能力、最大 RPM / TPM | 已核对 |
| 2026-11-26 | [Messages API](https://docs.volcengine.com/docs/ark/messages-api?lang=zh) | Anthropic 兼容 Messages 端点与 `x-api-key` 鉴权 | 已核对，未实测 |
| 2026-11-26 | [Coding Plan 概览](https://docs.volcengine.com/docs/ark/coding-plan-personal-plan-overview) | 付费订阅；套餐额度仅在 AI 编程 / 智能体工具中生效，不可用于 API 调用 | 已核对，不收录为免费 |
| 2026-11-26 | 控制台开通管理页、产品活动介绍页 | 免费额度覆盖的模型、有效期 | 需登录，**待核验** |
