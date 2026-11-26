# Cerebras 免费模型与规则

## 试用额度可用模型

来源：[模型清单](https://inference-docs.cerebras.ai/models/overview) 及各模型页，两个模型均标注 Free Trial 与 Pay as You Go。

| 模型 ID | 能力 | 上下文（试用 / 付费） | 最大输出（试用 / 付费） | 官方写明的功能 |
| --- | --- | --- | --- | --- |
| `gpt-oss-120b` | 📝 文本生成 | 65k / 131k | 32k / 40k | Reasoning、Streaming、Structured Outputs、Tool Calling、Prompt Caching |
| `qwen-3.8-27b` | 📝 文本生成 · 👁 图片理解 | 64K（65,536）/ 128K（131,072） | 32K（32,768）/ 40K（40,960） | Image Inputs、Reasoning、Streaming、Sampling Controls、Structured Outputs、Tool Calling、Parallel Tool Calling、Prompt Caching |

- `qwen-3.8-27b` 图片输入：试用每请求 2 张（PayGo 10 张），请求体总大小 10 MiB；只接受 base64 的 PNG 或 JPEG，不支持外部 HTTPS 图片 URL；输出为文本。
- `gpt-oss-120b` 模型页没有写图片输入，不推断。
- 官方对试用层上下文有限制，约为付费层的一半。

## 已弃用或下线（不要当成可用模型）

来源：[弃用页](https://inference-docs.cerebras.ai/support/deprecation)。

- `gemma-4-31b`：自 2026-09-03 起不再提供 Shared Inference（OpenCode 接入页的模型表仍列着它，官方文档内部不一致，以弃用页为准）
- `zai-glm-4.7`：2026-08-17
- `llama3.1-8b`、`qwen-3-235b-a22b-instruct-2507`：2026-05-27
- `qwen-3-32b`、`llama-3.3-70b`：2026-02-16
- `zai-glm-4.6`：2026-01-20

## 免费规则

- **类型**：一次性赠金，$5，需先添加有效支付方式，激活后 30 天过期；到期或用完后 API 和 Playground 暂停，除非另购 PayGo 额度（[账号与计费](https://inference-docs.cerebras.ai/console/account-billing)）。
- **不是永久免费层**：官方没有写任何永久免费额度。
- **滥用**：赠金不可转让、无现金价值，欺诈、滥用、重复账号可能被撤销。
- **价格对照**（付费层，供估算 $5 能用多少）：`gpt-oss-120b` $0.35/M 输入、$0.75/M 输出；`qwen-3.8-27b` $0.99/M 输入、$1.49/M 输出（[定价页](https://www.cerebras.ai/pricing)）。定价页脚注写明 Developer 层不适合生产使用。

## 限流

来源：[限流页](https://inference-docs.cerebras.ai/support/rate-limits)，以组织为单位。页面没有给 RPD 和并发数。

| 模型 | 层级 | RPM | 未缓存 TPM | 总 TPM | TPH | TPD |
| --- | --- | --- | --- | --- | --- | --- |
| `gpt-oss-120b` | Free Trial | 5 | 30K | 90K | 1M | 1M |
| `qwen-3.8-27b` | Free Trial | 300 | 150K | 750K（临时上调，标准值 450K） | N/A | N/A |
| `gpt-oss-120b` | Developer | 1K | 500K | 1.5M | 无 | 无 |
| `qwen-3.8-27b` | Developer | 300 | 150K | 750K | 无 | 无 |

当前账号的实际限额在控制台 Limits 页（需登录，待核验）。

## 模型选择建议

- 默认模型：`gpt-oss-120b`（但试用层仅 5 RPM、每日 1M token）
- 高频调用：`qwen-3.8-27b`（试用层 300 RPM）
- 图片理解：`qwen-3.8-27b`
- 编程推荐：官方未给编程适配说明，待核验

最后核验日期：2026-11-26
