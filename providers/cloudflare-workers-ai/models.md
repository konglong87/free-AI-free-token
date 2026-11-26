# Cloudflare Workers AI 免费模型与规则

## 已抽查的免费可用模型

来源：[模型目录](https://developers.cloudflare.com/workers-ai/models/)各模型页（上下文与能力徽标）；[定价页](https://developers.cloudflare.com/workers-ai/platform/pricing/)的付费模型名单。我们抽查了下表 9 个模型页，均没有「Paid access required」提示，但其余模型页未逐个核对，**免费资格仍以官方为准**。

| 模型 ID（原样，带 `@cf/`） | 能力 | 上下文 | 官方写明的功能 |
| --- | --- | --- | --- |
| `@cf/openai/gpt-oss-120b` | 📝 文本生成 | 128,000 | 函数调用、推理（low / medium / high）、Batch |
| `@cf/openai/gpt-oss-20b` | 📝 文本生成 | 128,000 | 函数调用、推理、Batch |
| `@cf/google/gemma-4-26b-a4b-it` | 📝 文本生成 · 👁 图片理解 | 256,000 | 视觉、推理、函数调用、Batch |
| `@cf/zai-org/glm-4.7-flash` | 📝 文本生成 | 131,072 | 函数调用、推理 |
| `@cf/qwen/qwen3.8-27b` | 📝 文本生成 · 👁 图片理解 | 262,144 | 视觉、推理、函数调用、Batch |
| `@cf/meta/llama-3.3-70b-instruct-fp8-fast` | 📝 文本生成 | 24,000 | 函数调用、Batch |
| `@cf/meta/llama-4-scout-17b-16e-instruct` | 📝 文本生成 · 👁 图片理解 | 131,000 | 视觉、函数调用、Batch |
| `@cf/mistralai/mistral-small-3.1-24b-instruct` | 📝 文本生成 · 👁 图片理解（待核验） | 128,000 | 能力徽标只列函数调用；介绍文字写有视觉理解，两处不一致 |
| `@cf/meta/llama-3.2-11b-vision-instruct` | 👁 图片理解（文本输出） | 128,000 | 视觉 |

另有 `@cf/qwen/qwen2.5-coder-32b-instruct`、`@cf/ibm-granite/granite-4.0-h-micro`、`@cf/aisingapore/gemma-sea-lion-v4-27b-it` 等，上下文与免费资格待核验。

### 非文本类（免费资格待核验）

官方目录列出：文生图 `@cf/black-forest-labs/flux-1-schnell`、`flux-2-dev`、`flux-2-klein-4b`、`flux-2-klein-9b`；语音识别 `@cf/openai/whisper-large-v3-turbo`；TTS `@cf/deepgram/aura-2-en`；嵌入 `@cf/baai/bge-m3` 等。它们是否消耗同一份免费 Neurons、是否在免费账户可用，本目录没有逐个核验，因此**不记为免费模型**。

## 免费账户不可用（需 Workers Paid 或预充值 AI Gateway credits）

定价页 2026-11-18 版写「Some models require a paid billing method」，名单：`@cf/moonshotai/kimi-k2.6`、`kimi-k2.7-code`、`@cf/zai-org/glm-5.2`、`glm-5.3`、`glm-5.3-flash`、`@cf/deepseek-ai/deepseek-v4-flash-0731`、`deepseek-v4.1-flash-1017`、`deepseek-v4-pro-0813`、`@cf/thinkingmachines/inkling-2-small`。

## 已下线

[官方公告](https://developers.cloudflare.com/changelog/post/2026-11-05-planned-model-deprecations/)：2026-11-20 起 24 个旧模型返回 HTTP 410（错误码 5028），LLM 包括 `deepseek-r1-distill-qwen-32b`、`nemotron-3-120b-a12b`、`deepseek-v4-pro-0813`、`deepseek-v4-flash-0731`、`qwq-32b`、`qwen3-30b-a3b-fp8`、`kimi-k2.6`、`kimi-k2.7-code`、`glm-5.2`、`glm-5.3-flash`、`llama-3.1-8b-instruct-fp8`、`llama-3.2-1b-instruct`、`llama-3.2-3b-instruct` 及 3 个 LoRA 变体。目录页今天仍显示其中部分并标 Deprecated，不要使用。官方 OpenAI 兼容页的示例仍用 `@cf/meta/llama-3.1-8b-instruct`，该模型在 2026-05-30 的下线名单里，示例已过时。

`@cf/cloudflare/clef`、`clef-omni`、`clef-flash`、`clairic`、`clairic-flash` 是 Cloudflare 自家的决策 / 抽取类模型，不是通用聊天模型，不收录。

## 免费规则

- **类型**：每日循环，Workers Free 账户 10,000 Neurons / 天，每日 00:00 UTC 重置。
- **超额**：操作报错，错误码 3036、HTTP 429，提示升级 Workers Paid（[错误码](https://developers.cloudflare.com/workers-ai/platform/errors/)）；不会自动扣费。
- **付费对照**：Workers Paid 每天同样含 10,000 Neurons，超出 $0.011 / 1000 Neurons。
- Neurons 与 token 的换算官方只在各模型定价里按模型给出，本目录不做换算。

## 限流

来源：[限流页](https://developers.cloudflare.com/workers-ai/platform/limits/)。

- 文本生成：300 请求 / 分钟（需要 Workers Paid 的模型除外）。
- 需要 Workers Paid 的模型：标准计费 20 请求 / 分钟（每账户每模型）；预充值 AI Gateway credits 50 请求 / 分钟。
- 嵌入 3000 请求 / 分钟（`bge-large-en-v1.5` 为 1500）；文生图 720；语音识别 720。
- TPM、TPD、并发：官方未公布。

## 模型选择建议

- 默认模型：`@cf/openai/gpt-oss-120b`（128K，函数调用、推理）
- 长上下文：`@cf/qwen/qwen3.8-27b`（262,144）、`@cf/google/gemma-4-26b-a4b-it`（256,000）
- 图片理解：`@cf/google/gemma-4-26b-a4b-it`、`@cf/qwen/qwen3.8-27b`、`@cf/meta/llama-4-scout-17b-16e-instruct`
- 编程推荐：官方未给编程适配说明，待核验

最后核验日期：2026-11-26
