# Google AI Studio 免费模型与规则

## 免费层文本 / 多模态模型

来源：[定价页](https://aistudio.google.com/docs/pricing) 逐模型的 Free Tier 列（标 Free of charge）；能力与上下文来自[模型清单](https://aistudio.google.com/docs/models)。官方没有单独写免费层的上下文长度和最大输出，下表「上下文」是模型页对该模型系列的描述，**免费层是否相同待核验**。

| 模型 ID | 能力 | 上下文 | 备注 |
| --- | --- | --- | --- |
| `gemini-4-hydrogen` | 📝 文本生成 · 👁 图片理解 | 官方：Gemini 4 系列输入 1M tokens；输出 64K 或 1M | Gemini 4 系列输入含文本、图片、音频、视频、PDF |
| `gemini-3.8-flash` | 📝 文本生成（其他能力待核验） | 未写明 | OpenAI 兼容页示例模型 |
| `gemini-3.6-flash` | 📝 文本生成（其他能力待核验） | 未写明 |  |
| `gemini-3.5-flash-lite` | 📝 文本生成（其他能力待核验） | 未写明 | 免费层不提供 context caching |
| `gemini-3.1-flash-lite` | 📝 文本生成（其他能力待核验） | 未写明 | **计划 2027-05-07 停服**，替换为 `gemini-3.5-flash-lite` |
| `gemini-3-flash-preview` | 📝 文本生成（其他能力待核验） | 未写明 | 预览模型 |
| `gemini-2.5-pro` | 📝 文本生成（其他能力待核验） | 未写明 | **2026-12-01 起请求自动路由到 `gemini-3.8-flash`**；已不能在 AI Studio 选择 |
| `gemini-2.5-flash` | 📝 文本生成（其他能力待核验） | 未写明 | **2026-12-01 起请求自动路由到 `gemini-3.5-flash-lite`**；已不能在 AI Studio 选择 |
| `gemini-2.5-flash-lite` | 📝 文本生成（其他能力待核验） | 未写明 |  |

能力说明：官方只明确了 Gemini 4 系列的多模态输入。其他模型页面没有写输入类型，这里不靠模型名推断，保留「待核验」。

## 免费层的其他模型（非对话，本仓库统一标签未覆盖）

- 音频：`gemini-3.5-transcribe`、`gemini-3.5-transcribe-live`、`gemini-3.8-live`、`gemini-3.8-live-extended-thinking`、`gemini-3.1-flash-live-preview`、`gemini-3.5-live-translate-preview`、`gemini-2.5-flash-native-audio-preview-12-2025`（计划 2027-03-31 停服）
- TTS：`gemini-3.8-flash-tts`、`gemini-3.8-flash-lite-tts`、`gemini-3.1-flash-tts-preview`（计划 2026-12-18 停服）、`gemini-2.5-flash-preview-tts`（计划 2026-12-18 停服）
- embedding：`gemini-embedding-2`、`gemini-embedding-001`
- 其他：`gemini-robotics-er-2-preview`、`gemini-robotics-er-2-streaming-preview`、`gemini-decision-1`（官方描述为输出 0 token 的评估模型，通过 Decisions API 调用）；Gemma 4 页面标免费，但没有给出 API 模型 ID（待核验）

## 免费层明确不可用（Not available）

`gemini-4-carbon`、`gemini-4-argon`、`gemini-4-neon`、`gemini-3.1-pro-preview`、`gemini-3.1-pro-preview-customtools`、`genie-4`、`gemini-omni-1.2-flash`、`gemini-omni-1.1-flash`，以及**所有图片生成类**（`gemini-nano-banana-2.1`、`gemini-3.1-flash-image`、`gemini-3.1-flash-lite-image`、`gemini-3-pro-image`、`gemini-2.5-flash-image`）、`gemini-2.5-pro-preview-tts`、`lyria-3.5`、`lyria-3-*`。Google Search grounding 免费层也不提供。因此免费层没有文生图能力。

## 已停服

`gemini-2.0-flash`、`gemini-2.0-flash-lite`、`gemini-3-pro-preview`、`gemini-3.1-flash-lite-preview`。弃用页写明所列日期是最早可能的退役日期，确切日期会提前通知（[弃用页](https://aistudio.google.com/docs/deprecations)）。

## 限流

- 按**项目**计，不按 API key（[限流页](https://aistudio.google.com/docs/rate-limits)）。
- 官方页面只公布了免费层 `gemini-decision-1`：200 RPM、50,000 TPM、10,000 RPD。
- 其余免费模型的 RPM / TPM / RPD 页面写「在 AI Studio 查看」，需登录，**待核验**。
- 官方声明限额不保证，实际容量可能变化；免费层请求使用可被抢占的容量（[状态页](https://aistudio.google.com/status)）。

## 数据条款（免费层务必注意）

- Unpaid Services（含免费层）：Google 可使用你提交的内容和生成的回复来改进产品，人工审核员可能读取、标注输入输出（审核前已与账号、key、项目脱钩）；定价页对免费层行写「用于改进产品：是」。不要提交敏感、机密或个人信息。
- Paid Services：不用于改进产品。
- **EEA、瑞士、英国**：条款写明 Paid Services 的数据条款适用于这些地区用户的所有服务，包括 AI Studio 与 Gemini API 的免费额度（[服务条款](https://aistudio.google.com/docs/terms)）。

## 门槛

- 信用卡：官方对免费层没有写要求，也没有写「不需要」，**待核验**。
- 设备验证：第一个 AI Studio 项目通常不需要；第二个及以后的免费项目、或在 AI Studio 之外创建 / 导入的项目，需用 Android 或 iPhone 扫码验证，官方称不需要手机号；未验证返回 403（[项目验证](https://aistudio.google.com/docs/verification)）。
- 地区：可用地区覆盖英国、瑞士和全部 EEA 成员国（[可用地区](https://aistudio.google.com/docs/available-regions)）。

## 模型选择建议

- 默认模型：`gemini-4-hydrogen`（免费层可用，数值限流待核验）
- 低延迟 / 低成本：`gemini-3.5-flash-lite`
- 长上下文：Gemini 4 系列（官方：输入 1M tokens）
- 多模态：`gemini-4-hydrogen`（官方明确支持图片等输入）
- 编程推荐：官方未给免费层编程适配说明，待核验

最后核验日期：2026-11-26
