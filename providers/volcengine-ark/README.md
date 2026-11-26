# 火山方舟（字节跳动 · 豆包）

## 当前状态

- 文档状态：官方文档已核对；开通管理页、产品活动介绍页需登录，未核对；未使用 key 实测
- 免费方式：**新用户一次性免费额度**（官方写「每个模型提供一定额度」，示例为 50w token；有效期官方未写）；另有**限时活动**「协作奖励计划」，需授权数据换额度。没有永久免费，也没有每日循环额度
- 额度用完：默认 API 调用失败；可开启「安心体验模式」，只消耗赠送额度，用完自动暂停，不扣费
- **不算免费**：Coding Plan、Agent Plan 是付费订阅，且官方明确不可用于通用 API 调用
- 最后核验日期：2026-11-26

## 一句话说明

火山方舟是字节跳动的大模型服务平台，提供豆包（Doubao）等模型的 OpenAI 兼容和 Anthropic 兼容接口；新用户每个模型有一次性免费额度，用完需实名认证并开通后按量付费。

## 快速开始

1. 注册火山引擎账号，打开[方舟控制台](https://ark.volcengine.com/region:cn-beijing/apikey)创建 API Key（入口取自官方[快速入门](https://docs.volcengine.com/docs/ark/quick-start?lang=zh)）。
2. 在控制台「开通管理」确认哪些模型有免费额度（**需登录，待核验**）；想避免产生费用，先开启[安心体验模式](https://docs.volcengine.com/docs/ark/free-tokens-only-mode)。
3. Base URL 用 `https://ark.cn-beijing.volces.com/api/v3`，模型 ID 从[官方模型列表](https://docs.volcengine.com/docs/ark/model-list?lang=zh)取。
4. 发最小 Chat Completions 请求验证；注意控制台额度显示有小时级延迟。

## 核心信息

| 项目 | 值 |
| --- | --- |
| 官网 | https://www.volcengine.com/product/ark（页面为 JS 渲染，本次未读到正文，不引用营销文案） |
| 文档 | https://docs.volcengine.com/docs/ark/ |
| API Key 入口 | https://ark.volcengine.com/region:cn-beijing/apikey |
| OpenAI Base URL | `https://ark.cn-beijing.volces.com/api/v3`（对话 `POST /chat/completions`，另有 Responses API） |
| Anthropic 兼容 | 有：`POST https://ark.cn-beijing.volces.com/api/compatible/v1/messages`，鉴权头 `x-api-key`（[Messages API](https://docs.volcengine.com/docs/ark/messages-api?lang=zh)）；Claude Code 的 `ANTHROPIC_BASE_URL` 写法待核验 |
| 免费类型 | 一次性赠金（新用户，按模型）+ 限时活动（协作奖励计划，需数据授权） |
| 免费额度 | 每模型「一定额度」，官方示例 500k（50w）token；有效期官方未写；仅抵扣在线推理按 token 后付费 |
| 限流 | 官方模型列表给出各模型「最大 RPM / TPM」（非刚性保障）；RPD / TPD 官方未见 |
| 门槛 | 火山引擎账号；用完后继续使用、开启安心体验模式、参加协作奖励计划都要求实名认证（个人或企业）；信用卡、地区限制官方文档未见 |
| 核验状态 | 官方页面已核对，未实测；免费额度覆盖哪些模型待核验 |



## 目录

- [官方链接与核验记录](official-links.md)
- [模型与免费规则](models.md)
- [API 接入、鉴权和限流](api.md)
- [Claude Code](integrations/claude-code.md) · [OpenClaw](integrations/openclaw.md) · [OpenCode](integrations/opencode.md) · [Codex](integrations/codex.md) · [Hermes](integrations/hermes.md)
