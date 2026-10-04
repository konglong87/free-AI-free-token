# Models

## 免费模型

| 模型名 | 类型 | 接口 | 备注 |
| --- | --- | --- | --- |
| DeepSeek V4.1 Flash（`deepseek-flash`） | 文本 / 推理 / 代码 / 工具调用 / 视觉理解 | `/chat/completions`、`/responses` | 热门首选；552B MoE，面向复杂推理、代码开发、多模态理解及多步骤 Agent 任务 |
| Kimi K3（`kimi-k3`） | 文本 / 视觉 / 推理 / 代码 / Agent | `/chat/completions`、`/responses` | 热门首选；2.8T 参数、1M Token 上下文，适合长程编程、知识工作和复杂推理 |
| `sensenova-6.7-flash-lite` | 文本 / 对话 | `/chat/completions` | Token Plan 显示为基础免费模型；API 文档提供示例 |
| SenseNova U1 Fast | 文本 / 对话 | 待核验 | Token Plan 显示为基础免费模型；API 模型 ID 待核验 |
| `glm-5.2` | 文本 / 对话 / 推理 | `/chat/completions` | 官方模型文档截图显示模型 ID 为 `glm-5.2` |
| DeepSeek V4 Flash | 文本 / 对话 | 待核验 | Token Plan 模型列表显示；API 模型 ID 待核验 |

## 模型选择建议

- 默认热门模型：`deepseek-flash`
- 长程编程 / 多模态 Agent：`kimi-k3`
- 轻量快速模型：SenseNova U1 Fast
- 长任务 / 推理模型：`glm-5.2`
- DeepSeek 系列：DeepSeek V4.1 Flash（`deepseek-flash`）

## 模型来源

- Token Plan：https://www.sensenova.cn/token-plan
- SenseNova 官方模型文档：https://platform.sensenova.cn/docs
- API 文档：https://github.com/OpenSenseNova/SenseNova6.7/blob/main/API_CN.md
- 模型文档截图：`assets/sensenova-model-docs.png`
- 公测免费套餐截图：`assets/sensenova-free-plan.png`
- 最后核验日期：2026-10-04

## 说明

官网 Token Plan 标注：公测期间基础模型免费；特殊模型以文档标注为准。接入前应以官方 Token Plan 和 API 文档为准。
