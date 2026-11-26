# 硅基流动 免费模型与规则

## 免费对话 / 视觉模型

下表的「免费」判定来自[价格页](https://www.siliconflow.cn/pricing)；模型 ID、上下文和能力标签来自[公开模型中心](https://www.siliconflow.cn/models)。价格页只显示展示名，免费版与收费版的 ID 对应关系依据[官方 FAQ](https://docs.siliconflow.cn/docs/userguide/faqs/misc)（免费版用原名，收费版加 `Pro/` 前缀）推断，**未在控制台逐个核验，入库使用前请在模型广场确认**。

| 模型 ID | 能力 | 上下文 | 官方标签 | 备注 |
| --- | --- | --- | --- | --- |
| `Qwen/Qwen3-8B` | 📝 文本生成 | 128K | 工具调用 | 价格页「免费」 |
| `deepseek-ai/DeepSeek-R1-0528-Qwen3-8B` | 📝 文本生成 | 128K | 官方描述为 DeepSeek-R1-0528 蒸馏到 Qwen3 8B；未见工具调用标签 | 价格页「(Free)」 |
| `Qwen/Qwen3.5-4B` | 📝 文本生成 · 👁 图片理解 | 256K | 视觉输入、工具调用 | 价格页「免费」 |
| `Qwen/Qwen2.5-7B-Instruct` | 📝 文本生成 | 32K | 工具调用、前缀续写 | 价格页「(Free)」；`Pro/Qwen/Qwen2.5-7B-Instruct` 是收费版 |
| `THUDM/GLM-4-9B-0414` | 📝 文本生成 | 32K | 工具调用 | 价格页「免费」 |
| `THUDM/GLM-Z1-9B-0414` | 📝 文本生成 | 128K | 工具调用 | 价格页「免费」 |
| `tencent/Hunyuan-MT-7B` | 📝 文本生成（翻译） | 32K | 翻译；前缀续写 | 价格页「免费」 |
| `deepseek-ai/DeepSeek-OCR` | 👁 图片理解（OCR） | 8K | 视觉输入（OCR） | 价格页「免费」 |
| `PaddlePaddle/PaddleOCR-VL-1.5` | 👁 图片理解（OCR） | 官方卡片未写 | 视觉输入（OCR） | 价格页「免费」 |
| `XingChenAGI/Xing4.0-29B` | 📝 文本生成 | 256K | 工具调用 | **官方公告：2026-12-04 下线** |

能力只采用官方卡片明确写出的标签：「视觉输入」记为图片理解，OCR 模型不等于通用看图问答。「工具调用」只表示官方标注，未实测。

## 其他免费模型（非对话）

价格页另有标「免费」的：嵌入 / 重排 `bge-m3`、`bge-reranker-v2-m3`、`bge-large-zh-v1.5`、`bge-large-en-v1.5`；生图 `Kolors`；语音 `SenseVoiceSmall`、XingChen ASR 系列。ID 未逐个核对（待核验）。

## 不要当成免费

- 所有 `Pro/…` 版本；`Qwen3.5-9B`、`Qwen3-14B`、`Qwen3-VL-8B` 等价格页标价的模型；DeepSeek-V4、GLM-5.3 等主力模型。
- 已下线：`THUDM/glm-4-9b-chat`、`Qwen/Qwen2-7B-Instruct`（2026-03-17）；`DeepSeek-R1-Distill-Qwen-7B/14B/32B` 等（2026-04-29），见[更新公告](https://docs.siliconflow.cn/docs/release-notes/overview)。
- 国际站旧文档里的「免费大语言模型列表」是历史残留，不采信。

## 免费规则

- **类型**：永久免费模型，限速免费。官方：「免费模型调用免费，账户的费用账单中看到此类模型的费用为 0」，并保留免费模型的解释权，模型可能随时下线（[Rate Limits](https://docs.siliconflow.cn/docs/userguide/faqs/rate-limit-and-upgradation)）。
- **不是赠金**：官方文档未见「新用户注册赠送金额」说明（待核验）。2025-12-17 公告称，2025-11-30 前未用完的赠送余额转为代金券，此后平台激励以代金券发放（[更新公告](https://docs.siliconflow.cn/docs/release-notes/overview)）。
- **门槛**：实名认证后才能使用全部免费模型；2026-05-15 起未实名账户无法使用平台功能。

## 模型选择建议

- 默认文本模型：`Qwen/Qwen3-8B`（128K，工具调用）
- 长上下文：`Qwen/Qwen3.5-4B`、`XingChenAGI/Xing4.0-29B`（256K，后者将下线）
- 推理：`THUDM/GLM-Z1-9B-0414`、`deepseek-ai/DeepSeek-R1-0528-Qwen3-8B`（是否带推理输出待核验）
- 图片理解：`Qwen/Qwen3.5-4B`；OCR：`deepseek-ai/DeepSeek-OCR`
- 编程 agent：官方未给免费模型的编程适配说明，参数规模较小，效果待核验

## 来源

- [价格页](https://www.siliconflow.cn/pricing)
- [公开模型中心](https://www.siliconflow.cn/models)
- [Rate Limits](https://docs.siliconflow.cn/docs/userguide/faqs/rate-limit-and-upgradation)
- [更新公告](https://docs.siliconflow.cn/docs/release-notes/overview)

最后核验日期：2026-11-26
