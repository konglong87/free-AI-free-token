# B.AI 免费模型资料包设计

## 目标

新增 B.AI 供应商资料，收录用户截图中标记为“赠送”的三个模型，并保留页面截图作为免费状态与 API 接入信息的证据。

## 范围

- 供应商目录：`providers/b-ai/`
- 模型：`DeepSeek-V4-Flash`、`DeepSeek-V4-Flash-Vision-Exp`、`Hy3`
- 来源页面：`https://chat.b.ai/key`、`https://chat.b.ai/leaderboard`
- API Base URL：`https://api.b.ai/v1`
- 证据图片：模型/key 页面截图、调用示例截图
- 根 README 供应商索引

不在本次范围内：API 实测、客户端实测、未出现在截图中的模型、真实密钥或控制台敏感信息。

## 架构与数据流

沿用仓库“一家供应商一个目录”的边界，从 `_template` 复制固定文档结构。`models.md` 只描述截图确认的三个免费模型；`api.md` 记录截图中的 OpenAI-compatible 接入字段，并将未实测能力标记为“待核验”；`official-links.md` 记录来源与核验日期；`README.md` 和五个 integration 文件提供统一入口但不宣称已实测。

证据图片置于 `providers/b-ai/assets/`，文档通过相对路径引用。这样模型清单、接入说明和证据可以独立更新，不影响其他供应商。

## 错误处理与真实性

- 页面动态内容无法在当前环境直接复核时，不补写截图之外的模型或能力。
- 不确定的模型上下文、streaming、tool calling、结构化输出和客户端兼容性统一标记为“待核验”。
- 图片中如含 API key、token、cookie 或账号信息，不纳入仓库；本次仅提交用户提供的公开示例与模型列表截图。

## 验证方案

1. 运行 `./scripts/validate-layout.sh`，确认固定目录结构完整。
2. 用 `file` 检查两张证据图片为有效 PNG。
3. 用 `rg` 检查三个模型名、Base URL 与来源链接均出现在对应文档中，且没有真实密钥。
4. 查看 git diff 后提交并推送。
