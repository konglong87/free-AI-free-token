# K-EXAONE 2.5 模型与免费规则

## `k-exaone_v2.5`

- **能力**：📝 文本输入 → 文本生成；韩语 / 英语，官方定位为 Agent、代码与工作流自动化。Models 页列出工具调用、结构化输出、流式。
- **规模**：Models 页标注 750B 参数（31B 激活）。
- **官方标称上下文**：524,288 Token（512K），`max_tokens` 包含推理 Token；输入 + `max_tokens` 超限返回 400。
- **推理**：官方示例的流式和非流式响应都带 `usage.reasoning_tokens`；流式响应的推理内容在 `delta.reasoning_content`。`reasoning_effort` 在 API 参数表中为 `low` / `high` / `max`（默认 `max`），官方 Agent 接入页则写 `low` / `medium` / `high` / `max` 并在示例中用 `medium`，两页不一致，以实际返回为准。`chat_template_kwargs.preserve_thinking`（默认 `false`）控制多轮对话是否保留之前 assistant 轮次的推理。
- **采样默认值**：`temperature=1.0`、`top_p=0.95`（官方 Agent 示例同样使用这两个值）。
- **免费方式**：技术验证阶段免费，非商用。
- **核验**：官方公告与文档已核对，实际调用待实测。

**注意**：同平台的 EXAONE Tabular 是表格数据预测模型，不是对话 LLM，本仓库不收录。

## 免费规则与限制

- **免费**：官方公告称开发者、初创公司、企业、研究机构等都可以免费调用，拿到 API Key 即可，用于服务开发和研究。条款第 9 条保留日后收费的可能，收费前会在网站公告。
- **非商业用途**：条款第 2 条第 1 项只授予非商业使用权限，不得用于营利，也不得让第三方通过你的产品或服务调用模型。准备接入收费产品前先确认合规。
- **Key 30 天有效**：条款第 3 条第 4 项规定 Key 自签发日起 30 天有效，到期自动停用，需要重新签发。
- **数据用于训练**：条款第 2 条第 3 项规定输入输出默认可用于模型训练，用户可以拒绝；拒绝前已用于训练的部分无法删除。不要发送敏感数据。
- **调用上限**：Rate Limits 页未公开具体数值；同意可选的训练用途条款可获得更高的每日调用上限，不同意也能用。超限返回 429。

## 来源

- [官方公告](https://www.lgresearch.ai/news/view?seq=755)
- [Models](https://docs.exaone.ai/ko/docs/models)
- [K-EXAONE 2.5 接口说明](https://docs.exaone.ai/ko/docs/k-exaone-2.5)
- [Rate Limits](https://docs.exaone.ai/ko/docs/rate-limits)
- [服务条款](https://docs.exaone.ai/ko/terms/terms-of-service)
