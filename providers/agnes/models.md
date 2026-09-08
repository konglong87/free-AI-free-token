# Models

## 免费模型

| 模型名 | 类型 | 接口 | 备注 |
| --- | --- | --- | --- |
| `agnes-3.0-flash` | 文本 / 图片 URL 输入 / Agent / 工具调用 | `/v1/chat/completions`、`/v1/responses`、`/v1/messages` | 512K 上下文，最大输出 65,536 tokens；官方文档显示当前输入、缓存输入和输出价格均为 $0 |
| `agnes-2.5-flash` | 对话 / 文本生成 | `/v1/chat/completions` | 2026-07-29 两条调用记录的消费金额均为 0 美分，已实测免费 |
| `agnes-2.0-flash` | 对话 / 文本生成 / 推理 | `/v1/chat/completions` | 官网文档显示为 Agnes 2.0 Flash |
| `agnes-image-2.0-flash` | 文生图 / 图像编辑 | 待核验 | Agnes AI 官网截图显示为免费模型 |
| `agnes-image-2.1-flash` | 文生图 / 图像编辑 | 待核验 | 官网文档显示为 Agnes Image 2.1 Flash |
| `agnes-video-v2.0` | 视频生成 | 待核验 | Agnes AI 官网截图显示为免费模型 |

## 能力范围

截图中明确提到 Agnes AI API 可用于：

- AI 对话与文本生成
- 逻辑推理与内容理解
- 文生图与图像编辑
- 图生视频与视频生成
- 音视频同步生成
- Agent 工具与自动化工作流
- 创意内容生成与多模态交互应用

## 模型来源

- Agnes 3.0 Flash：https://www.agnes-ai.com/en/docs/agnes-30-flash
- Agnes 2.0 Flash：https://agnes-ai.com/doc/agnes-20-flash
- Agnes Image 2.1 Flash：https://agnes-ai.com/doc/agnes-image-21-flash
- 免费模型截图：`assets/agnes-free-models.png`
- Agnes 2.5 Flash 零消费调用记录：`assets/agnes-2.5-flash-zero-cost.png`
- 最后核验日期：2026-09-08
