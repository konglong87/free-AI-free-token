<h1 align="center">free-AI-free-token</h1>

<p align="center">免费 LLM 供应商、模型与 API 接入资料库</p>

<p align="center">
  <a href="https://github.com/konglong87/free-AI-free-token"><img src="https://img.shields.io/badge/providers-9-2563eb?style=flat-square" alt="供应商数量"></a>
  <a href="https://github.com/konglong87/free-AI-free-token/tree/main/providers"><img src="https://img.shields.io/badge/models-curated-059669?style=flat-square" alt="模型清单"></a>
  <a href="https://github.com/konglong87/free-AI-free-token"><img src="https://img.shields.io/badge/docs-Markdown-f59e0b?style=flat-square" alt="文档格式"></a>
  <a href="https://github.com/konglong87/free-AI-free-token/commits/main"><img src="https://img.shields.io/badge/updated-2026--08--27-7c3aed?style=flat-square" alt="最后更新日期"></a>
</p>

<p align="center">
  <a href="#免费供应商">浏览供应商</a> ·
  <a href="#快速开始">快速开始</a> ·
  <a href="readme-B.md">方案 B 预览</a> ·
  <a href="docs/architecture.md">架构说明</a> ·
  <a href="docs/contribution-guide.md">贡献指南</a>
</p>

<p align="center">
  <img src="providers/agnes/assets/agnes-free-models.png" width="280" alt="Agnes AI 免费模型证据截图">
  <img src="providers/sensenova/assets/sensenova-model-docs.png" width="280" alt="SenseNova 模型文档证据截图">
  <img src="providers/b-ai/assets/b-ai-free-models.png" width="280" alt="B.AI 免费模型证据截图">
</p>

<p align="center"><sub>证据图片来自仓库内已归档的供应商页面截图；详细来源见各供应商目录。</sub></p>

## 快速开始

| 目标 | 入口 |
| --- | --- |
| 选择免费模型 | [免费供应商](#免费供应商) · [模型文档](providers/) |
| 查看目录优先版首页 | [readme-B.md](readme-B.md) |
| 获取接入地址 | 打开对应供应商的 `api.md` |
| 配置 Agent / Coding Client | 查看供应商目录下的 `integrations/` |
| 新增供应商 | [贡献指南](docs/contribution-guide.md) |

## 免费供应商

按左侧标签找能力，再看右侧对应的模型 ID；点击供应商名称进入目录。

**理解 ≠ 生成**：👁 图片理解是看图回答；🖼 文生图 / 图生图是生成图片；🎬 文生视频 / 图生视频是生成视频。[查看标签说明](docs/model-capabilities.md)。

标签只汇总仓库已收录的模型，不代表平台全部能力或永久免费承诺；免费条件、限流和核验状态以供应商目录及官方资料为准。标注「待核验」的能力不作已确认推荐。

<table width="100%">
<thead>
<tr><th width="30%">供应商 · 免费能力</th><th width="70%">代表性模型 · 接入</th></tr>
</thead>
<tbody>
<tr>
<td valign="top"><a href="providers/agnes/"><strong>Agnes AI</strong></a><br><sub>📝 文本生成<br>👁 图片理解<br>🖼 文生图<br>🖼 图生图<br>🎬 文生视频<br>🎬 图生视频</sub></td>
<td valign="top"><strong>文本生成 · 图片理解</strong><br><code>agnes-3.0-flash</code><br><br><strong>文本生成</strong><br><code>agnes-2.5-flash</code><br><br><strong>文生图 · 图生图</strong><br><code>agnes-image-2.5-flash</code><br><br><strong>文生图 · 图像编辑</strong><br><code>agnes-image-2.1-flash</code><br><br><strong>文生视频 · 图生视频</strong><br><code>agnes-video-2.5-flash</code><br><br><a href="https://apihub.agnes-ai.com/v1">API</a> · <a href="https://agnes-ai.com/">注册</a> · <a href="providers/agnes/models.md">完整模型</a><br><sub>状态：✅ · <a href="https://www.agnes-ai.com/zh-Hans/docs/pricing">官方定价</a> · <a href="providers/agnes/assets/agnes-free-models.png">模型截图</a></sub></td>
</tr>
<tr>
<td valign="top"><a href="providers/zhipu-glm/"><strong>智谱 GLM</strong></a><br><sub>📝 文本生成<br>👁 图片理解<br>🖼 文生图<br>🎬 视频生成</sub></td>
<td valign="top"><strong>文本生成</strong><br><code>glm-4.7-flash</code><br><code>glm-4-flash-250414</code><br><br><strong>图片理解</strong><br><code>glm-4.6v-flash</code><br><code>glm-4.1v-thinking-flash</code><br><code>glm-4v-flash</code><br><br><strong>文生图</strong><br><code>cogview-3-flash</code><br><br><strong>视频生成</strong><br><code>cogvideox-flash</code><br><br><a href="https://open.bigmodel.cn/api/paas/v4/">API</a> · <a href="https://bigmodel.cn/">注册</a> · <a href="providers/zhipu-glm/models.md">完整模型</a><br><sub>状态：✅</sub></td>
</tr>
<tr>
<td valign="top"><a href="providers/sensenova/"><strong>商汤 SenseNova</strong></a><br><sub>📝 文本生成<br>👁 图片理解</sub></td>
<td valign="top"><strong>文本生成 · 图片理解</strong><br><strong>DeepSeek V4.1 Flash</strong><br><code>deepseek-flash</code><br><strong><code>kimi-k3</code></strong><br><br><strong>文本生成</strong><br><code>sensenova-6.7-flash-lite</code><br><code>glm-5.2</code><br><br><a href="https://token.sensenova.cn/v1">API</a> · <a href="https://www.sensenova.cn/">注册</a> · <a href="providers/sensenova/models.md">完整模型</a><br><sub>状态：✅ · <a href="https://platform.sensenova.cn/docs">官方模型文档</a></sub></td>
</tr>
<tr>
<td valign="top"><a href="providers/openrouter/"><strong>OpenRouter</strong></a><br><sub>📝 文本生成</sub></td>
<td valign="top"><strong>文本生成</strong><br><code>*:free</code> 模型（能力随具体模型变化）<br><br><a href="https://openrouter.ai/api/v1">API</a> · <a href="https://openrouter.ai/">注册</a> · <a href="providers/openrouter/models.md">完整模型</a><br><sub>状态：✅</sub></td>
</tr>
<tr>
<td valign="top"><a href="providers/b-ai/"><strong>B.AI</strong></a><br><sub>📝 文本生成<br>👁 图片理解<br>（待核验）</sub></td>
<td valign="top"><strong>文本生成</strong><br><code>DeepSeek-V4-Flash</code><br><code>Hy3</code><br><br><strong>文本生成 · 图片理解（待核验）</strong><br><code>DeepSeek-V4-Flash-Vision-Exp</code><br><br><a href="https://api.b.ai/v1">API</a> · <a href="https://chat.b.ai/">注册</a> · <a href="providers/b-ai/models.md">完整模型</a><br><sub>状态：✅ · <a href="providers/b-ai/assets/b-ai-free-models.png">模型截图</a></sub></td>
</tr>
<tr>
<td valign="top"><a href="providers/nvidia-nim/"><strong>NVIDIA NIM</strong></a><br><sub>📝 文本生成</sub></td>
<td valign="top"><strong>文本生成</strong><br><code>z-ai/glm-5.2</code><br><code>poolside/laguna-xs-2.1</code><br><code>z-ai/glm-5.1</code><br><code>qwen/qwen3.5-397b-a17b</code><br><br><a href="https://integrate.api.nvidia.com/v1">API</a> · <a href="https://build.nvidia.com/">注册</a> · <a href="providers/nvidia-nim/models.md">完整模型</a><br><sub>状态：🟡</sub></td>
</tr>
<tr>
<td valign="top"><a href="providers/modelscope/"><strong>ModelScope</strong></a><br><sub>📝 文本生成</sub></td>
<td valign="top"><strong>文本生成</strong><br><code>MiniMax/MiniMax-M2.5</code><br><code>qwen-qwen3-5-35b-a3b</code><br><code>qwen-qwen3-5-27b</code><br><br><a href="https://api-inference.modelscope.cn/v1">API</a> · <a href="https://modelscope.cn/">注册</a> · <a href="providers/modelscope/models.md">完整模型</a><br><sub>状态：🟡</sub></td>
</tr>
<tr>
<td valign="top"><a href="providers/groq/"><strong>Groq</strong></a><br><sub>📝 文本生成</sub></td>
<td valign="top"><strong>文本生成</strong><br><code>moonshotai/kimi-k2-instruct</code><br><code>moonshotai/kimi-k2-instruct-0905</code><br><code>groq/compound</code><br><br><a href="https://api.groq.com/openai/v1">API</a> · <a href="https://console.groq.com/">注册</a> · <a href="providers/groq/models.md">完整模型</a><br><sub>状态：🟡</sub></td>
</tr>
<tr>
<td valign="top"><a href="providers/ai21-labs/"><strong>AI21 Labs</strong></a><br><sub>📝 文本生成</sub></td>
<td valign="top"><strong>文本生成</strong><br><code>jamba-large-1-7</code><br><code>jamba-mini-2</code><br><br><a href="https://api.ai21.com/studio/v1">API</a> · <a href="https://studio.ai21.com/">注册</a> · <a href="providers/ai21-labs/models.md">完整模型</a><br><sub>状态：🟡</sub></td>
</tr>
<tr>
<td valign="top"><a href="providers/_template/"><strong>模板</strong></a><br><sub>能力待填写</sub></td>
<td valign="top">复制模板后，按模型已确认的输入 / 输出能力填写标签。<br><a href="docs/contribution-guide.md">贡献指南</a></td>
</tr>
</tbody>
</table>

## 仓库架构

```text
free-AI-free-token/
├── README.md
├── docs/
│   ├── architecture.md
│   └── contribution-guide.md
├── providers/
│   ├── <provider-slug>/
│   │   ├── README.md
│   │   ├── official-links.md
│   │   ├── models.md
│   │   ├── api.md
│   │   └── integrations/
│   └── _template/
└── scripts/
    └── validate-layout.sh
```

## 供应商目录规则

每个供应商必须放在 `providers/<provider-slug>/` 下，并从 `providers/_template/` 复制固定文件结构。这样每家供应商都能独立更新、审阅和核验。

## 采集边界

只收集公开、可追溯、可接入的信息：官网和官方文档、API key 获取入口、接入地址、免费模型、OpenAI-compatible 状态、客户端接入方式、核验日期和核验状态。

不收录真实 API key、token、cookie、账号信息、敏感截图或无法追溯来源的结论。未实测能力统一标记为 `待核验`。

## 固定文档模板

| 文件 | 作用 |
| --- | --- |
| `README.md` | 供应商总览、适用场景、快速开始 |
| `official-links.md` | 官网、控制台、文档、模型来源和核验记录 |
| `models.md` | 免费模型、模型能力、上下文和限制 |
| `api.md` | 接入地址、鉴权、兼容性和错误处理 |
| `integrations/*.md` | Claude Code、OpenClaw、OpenCode、Codex、Hermes 接入说明 |

## 新增供应商

1. 复制模板：`cp -R providers/_template providers/<provider-slug>`
2. 填写五类固定文档和 integrations 文件。
3. 在本页「免费供应商」表中增加入口。
4. 运行 `./scripts/validate-layout.sh`。
5. 保留官方来源链接、核验日期和状态。

## 文档与贡献

- [架构说明](docs/architecture.md)
- [贡献指南](docs/contribution-guide.md)
- [供应商目录](providers/)

欢迎补充新的免费供应商、修正模型状态或完善客户端接入记录。所有贡献都应附带可追溯的官方来源，并避免提交任何敏感凭据。
