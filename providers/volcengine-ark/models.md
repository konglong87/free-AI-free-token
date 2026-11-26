# 火山方舟 模型与免费规则

## 重要说明：官方没有公布「免费模型清单」

官方只写「每个模型提供一定额度的免费调用量」（[开通管理](https://docs.volcengine.com/docs/ark/activation-management?lang=zh)），没有写明覆盖哪些模型，也没有写有效期。下表是[官方模型列表](https://docs.volcengine.com/docs/ark/model-list?lang=zh)里的常用文本模型，**这些模型是否在免费额度内、协作奖励计划是否仍支持，均未确认，需登录控制台「开通管理」核验**。

| 模型 ID | 能力 | 上下文 | 最大 RPM / TPM（官方模型列表） |
| --- | --- | --- | --- |
| `doubao-seed-evolving` | 📝 文本生成 · 👁 图片理解（待核验） | 1024K | 500 / 1,000,000 |
| `doubao-seed-2-1-pro-260915` | 📝 文本生成 · 👁 图片理解（待核验） | 1024K | 500 / 1,000,000 |
| `doubao-seed-2-1-lite-260915` | 📝 文本生成 · 👁 图片理解（待核验） | 1024K | 30,000 / 5,000,000 |
| `doubao-seed-2-0-mini-260428`（标 LTS） | 📝 文本生成 · 👁 图片理解（待核验） | 256K | 30,000 / 5,000,000 |

- 能力栏：官方写的是「深度思考、文本生成、多模态理解、GUI 任务处理、工具调用、结构化输出」。「多模态理解」未在列表中细分输入类型，这里只能写「图片理解（待核验）」，不推断视频或音频。
- 最大 RPM / TPM 是模型列表页给出的上限，官方注明「非刚性保障，受平台负载 / 调用方式影响」。
- 模型 ID 带日期后缀（如 `-260915`），以官方模型列表为准，区分大小写。
- 图像、视频模型（如 `doubao-seedream-5-5-261028`、Seedance 系列）的额度归属待核验，本目录不把它们记为免费。

## 免费规则

- **新用户一次性额度**：注册火山引擎账号即可获得；按模型分别计算，主账号下共享，基础模型与其精调模型共享。只能抵扣按 token 后付费的在线推理，不能抵扣批量推理、插件、知识库，也不能抵扣上下文缓存的存储费（[免费推理额度](https://docs.volcengine.com/docs/ark/free-inference-quota)）。有效期官方未写。
- **额度用完**：默认调用失败；要继续用，需实名认证并手动开通对应模型推理服务，之后按实际 token 计费。
- **安心体验模式**：只消耗赠送的 50w token，接近额度时服务暂停、不产生费用，暂停时返回 HTTP 429（`TooManyRequests` / `SetLimitExceeded`）。默认关闭，**关闭后不能再开启**；未开通过模型服务的账号才可开启（[安心体验模式](https://docs.volcengine.com/docs/ark/free-tokens-only-mode)）。
- **协作奖励计划（限时活动，需授权数据）**：
  - 开通前提：已开通模型服务、关闭安心体验模式、完成实名认证；同一用户仅限一个账号参与。
  - 额度：首次授权某模型的推理接入点获得冷启动资源包；之后每个模型每日最多采集 5M tokens 的推理数据，次日发放等量资源包，每个资源包自发放起 30 个自然日有效。
  - **代价**：你的输入和输出经加密、匿名化后用于火山引擎及其关联公司的模型和算法优化，条款写「可永久使用」，已用于训练的数据无法撤回。不要发送敏感数据。
  - 现状：大部分模型已于 2026-09-01 和 2026-10-08 起退出计划，退出后模型仍可调用但按标准价格计费；活动页入口目前列出 Doubao-Seed-Evolving，完整名单需登录，待核验（[协作奖励计划](https://docs.volcengine.com/docs/ark/collaboration-reward-program?lang=zh)）。
- **体验中心**：官方称未开通前可在体验中心使用免费额度，那是网页体验，不等于 API 免费。

## 不算免费

- **Coding Plan**：付费订阅（Lite 原价 40 元/月、Pro 原价 200 元/月）。官方明确套餐额度「不可用于 API 调用」，在不支持的场景使用其 Base URL 和 key 可能被判定违规而停用或封号（[Coding Plan 概览](https://docs.volcengine.com/docs/ark/coding-plan-personal-plan-overview)）。2026-06-08 至 2026-11-08 的折扣活动已结束。
- **Agent Plan**：付费订阅，按 AFP 积分计费。

## 模型选择建议

- 默认模型：`doubao-seed-2-0-mini-260428`（RPM / TPM 上限较高，是否有免费额度待核验）
- 长上下文：`doubao-seed-evolving`（1024K，协作奖励计划入口列出的模型，待核验）
- 多模态：图片理解待核验
- 编程推荐：官方未给免费额度下的编程适配说明，待核验

## 来源

- [模型列表](https://docs.volcengine.com/docs/ark/model-list?lang=zh)
- [免费推理额度](https://docs.volcengine.com/docs/ark/free-inference-quota)
- [开通管理](https://docs.volcengine.com/docs/ark/activation-management?lang=zh)
- [协作奖励计划](https://docs.volcengine.com/docs/ark/collaboration-reward-program?lang=zh)

最后核验日期：2026-11-26
