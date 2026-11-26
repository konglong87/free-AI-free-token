# Cohere 模型与免费规则

## Trial key 可用的 Chat 型号

来源：[模型列表](https://docs.cohere.com/docs/models)（上下文、最大输出、模态）；[限流页](https://docs.cohere.com/docs/rate-limits)列出 Trial 可用的 Chat 型号：Command A+、A Reasoning、A Translate、A Vision、A、R+、R、R7B、North Mini Code。下表把这些名称对应到模型页的 ID，对应关系是按名称整理的，使用前请对照官方模型页。

| 模型 ID | 能力 | 上下文 | 最大输出 | 备注 |
| --- | --- | --- | --- | --- |
| `command-a-plus-05-2026` | 📝 文本生成 · 👁 图片理解 | 128k | 64k | 官方模态：文本、图像输入 |
| `command-a-03-2025` | 📝 文本生成 | 256k | 8k | 文本 |
| `command-a-reasoning-08-2025` | 📝 文本生成 | 256k | 32k | 文本（推理型号） |
| `command-a-vision-07-2025` | 📝 文本生成 · 👁 图片理解 | 128K | 8K | 官方模态：文本、图像输入 |
| `command-a-translate-08-2025` | 📝 文本生成（翻译） | 8K | 8k | 文本 |
| `command-r7b-12-2024` | 📝 文本生成 | 128k | 4k | 文本 |
| `command-r-08-2024` | 📝 文本生成 | 128k | 4k | 文本 |
| `command-r-plus-08-2024` | 📝 文本生成 | 128k | 4k | 文本 |
| `north-mini-code-1.0` | 📝 文本生成 | 256k | 64k | 官方称为 agentic coding 训练 |

- 官方称 Trial key 可用「全部模型与 API」，限流页只列出了上述 Chat 型号。`c4ai-aya-vision-32b`（模态写「文本、图像」）、`tiny-aya-*`、`c4ai-aya-expanse-32b` 不在限流页的 Trial 列表内，本目录不记为 Trial 可用。
- 已弃用 / 退役：`command-r`、`command-r-plus` 等旧别名标 Deprecated（2025-09-15）；`c4ai-aya-expanse-8b`、`c4ai-aya-vision-8b` 已 Retired（2026-04-04）。
- 官方称 Production key 在 Command A Reasoning 等新型号上与 Trial key 一样受限，需联系销售。

## 免费规则

- **类型**：永久（Trial key）。官方未写截止日期。
- **次数上限**：每月 1,000 次调用，对所有 Trial 调用生效；是否对 OpenAI 兼容端点与原生 API 合并计算，官方未说明（待核验）。
- **商用限制**：[定价页](https://cohere.com/pricing) FAQ 写明 Trial key 的调用免费，但限速且不得用于生产或商业用途；文档站限流页没有这句。
- **账号类型**：个人账号只能用 Trial key，组织账号需申请才能拿 Production key。

## 限流（Trial key）

来源：[限流页](https://docs.cohere.com/docs/rate-limits)。

- 每月 1,000 次调用。
- Chat：每模型 20 请求 / 分钟。
- 其他端点：Embed 2,000 inputs / 分钟，Embed(Images) 5 inputs / 分钟，EmbedJob 5 请求 / 分钟，Rerank 10 请求 / 分钟，Tokenize 100 请求 / 分钟，Audio Transcriptions 5 请求 / 分钟，Parse 500 请求 / 分钟，其余默认 500 请求 / 分钟。
- TPM、TPD、并发：官方未公布。

## 模型选择建议

- 默认模型：`command-a-03-2025`
- 长上下文：`command-a-03-2025`、`command-a-reasoning-08-2025`、`north-mini-code-1.0`（256k）
- 多模态：`command-a-vision-07-2025`、`command-a-plus-05-2026`（图像输入）
- 编程推荐：`north-mini-code-1.0`（官方描述，效果待实测）

最后核验日期：2026-11-26
