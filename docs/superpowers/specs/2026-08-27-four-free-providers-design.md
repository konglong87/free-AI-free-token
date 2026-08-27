# NVIDIA NIM、ModelScope、Groq、AI21 Labs 免费模型设计

## 目标

在现有 `free-AI-free-token` 文档架构中新增四家免费 LLM 供应商：NVIDIA NIM、ModelScope、Groq、AI21 Labs。每家供应商独立维护固定结构的 Markdown 文档，并在根 README 中提供可发现的索引入口。

## 范围

- 新增目录：`providers/nvidia-nim/`、`providers/modelscope/`、`providers/groq/`、`providers/ai21-labs/`。
- 每个目录包含模板规定的 9 个文件：总览、官方链接、模型列表、API 接入，以及 Claude Code、OpenClaw、OpenCode、Codex、Hermes 五类 integrations 文档。
- 每家收录 3–5 个代表性免费模型；AI21 Labs 当前公开参考清单仅确认 2 个模型，因此只收录这 2 个并在文档中说明数量限制。
- 根 README 增加四家供应商索引行，包含目录、代表模型和 Base URL。
- 参考资料使用已下载的公开中文 README；关键链接同时指向供应商官方入口或官方文档。

## 模型清单

| 供应商 | 模型 ID | 参考上下文 | 参考限流 |
| --- | --- | --- | --- |
| NVIDIA NIM | `z-ai/glm-5.2` | 1M | Up to 40 RPM |
| NVIDIA NIM | `poolside/laguna-xs-2.1` | 262K | Up to 40 RPM |
| NVIDIA NIM | `z-ai/glm-5.1` | 202K | Up to 40 RPM |
| NVIDIA NIM | `qwen/qwen3.5-397b-a17b` | 262K | 待供应商核验 |
| ModelScope | `MiniMax/MiniMax-M2.5` | 204K | 待供应商核验 |
| ModelScope | `qwen-qwen3-5-35b-a3b` | 131K | 2,000 RPD 总量，单模型上限待核验 |
| ModelScope | `qwen-qwen3-5-27b` | 131K | 2,000 RPD 总量，单模型上限待核验 |
| Groq | `moonshotai/kimi-k2-instruct` | 131K | 待供应商核验 |
| Groq | `moonshotai/kimi-k2-instruct-0905` | 131K | 待供应商核验 |
| Groq | `groq/compound` | 131K | 30 RPM，250 RPD |
| AI21 Labs | `jamba-large-1-7` | 256K | 200 RPM，10 RPS |
| AI21 Labs | `jamba-mini-2` | 256K | 200 RPM，10 RPS |

参考上下文和限流是发现来源中的当前记录；未由官方页面或真实 key 实测确认的字段必须标记为“待核验”。

## 架构与数据流

沿用“一家供应商一个目录”的边界。供应商事实只写入对应目录；根 README 只保留用于导航的摘要，不复制详细说明。

每家目录内部职责固定：

- `README.md`：状态、定位、快速开始和核心入口。
- `official-links.md`：官网、控制台、文档、Key 页面、模型来源和核验记录。
- `models.md`：代表模型、能力、上下文、限流和选择建议。
- `api.md`：Base URL、鉴权、端点、环境变量、错误处理和兼容性边界。
- `integrations/*.md`：五类客户端的统一配置字段与待核验验证记录。

四家供应商均按 OpenAI-compatible 方向记录配置，但能力声明严格以官方文档或实测为准。无法确认的 `/models`、streaming、tool calling、结构化输出和客户端专用配置保持“待核验”。

## 官方入口

- NVIDIA NIM：Base URL `https://integrate.api.nvidia.com/v1`；Key `https://build.nvidia.com/settings/api-keys`。
- ModelScope：Base URL `https://api-inference.modelscope.cn/v1`；Token `https://modelscope.cn/my/myaccesstoken`。
- Groq：Base URL `https://api.groq.com/openai/v1`；Key `https://console.groq.com/keys`。
- AI21 Labs：Base URL `https://api.ai21.com/studio/v1`；Key `https://studio.ai21.com/account/api-key`。

各目录同时记录官方文档首页、模型列表或模型文档入口。AI21 Labs 的免费状态和模型数量以参考项目当前清单为准，并保留官方模型文档链接供后续复核。

## 错误处理与安全

- 所有 API 文档统一覆盖 401、403、404、429 的原因和处理建议。
- 不提交真实 API key、token、cookie、账号信息或控制台敏感截图。
- 文档状态默认使用“已整理，待实测”或“待核验”，不得把聚合项目记录写成已实测结论。

## 验证方案

1. `./scripts/validate-layout.sh`：确认四个目录均包含模板要求的文件。
2. `git diff --check`：检查空白和 Markdown 差异问题。
3. 关键字段扫描：确认四家模型 ID、Base URL、Key 页面和官方链接均出现且拼写一致。
4. 敏感信息扫描：检查新增文档不包含疑似真实密钥、token 或 session 值。
5. 人工抽查：随机打开每家 `README.md`、`models.md`、`api.md` 与一个 integration，确认状态、来源和配置字段一致。

## 非目标

- 不实现 API SDK、代码生成器、运行时路由或自动同步程序。
- 不声称已完成真实 API key、客户端、streaming 或 tool calling 实测。
- 不同步参考项目中的全部模型，也不为 AI21 Labs 人为补充未确认的第三个模型。
