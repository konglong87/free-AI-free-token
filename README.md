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
| 获取接入地址 | 打开对应供应商的 `api.md` |
| 配置 Agent / Coding Client | 查看供应商目录下的 `integrations/` |
| 新增供应商 | [贡献指南](docs/contribution-guide.md) |

## 免费供应商

下表是导航索引；模型能力、限流和核验状态以供应商目录内文档为准。

<table width="100%">
<thead>
<tr><th width="10%">供应商</th><th width="9%">目录</th><th width="31%">代表性免费模型</th><th width="19%">接入地址</th><th width="19%">证据</th><th width="12%">状态</th></tr>
</thead>
<tbody>
<tr><td>Agnes AI</td><td><a href="providers/agnes/">目录</a></td><td><code>agnes-2.0-flash</code>、<code>agnes-2.5-flash</code>、<code>agnes-image-2.0-flash</code>、<code>agnes-image-2.1-flash</code>、<code>agnes-video-v2.0</code></td><td><a href="https://apihub.agnes-ai.com/v1">API</a></td><td><a href="providers/agnes/assets/agnes-free-models.png">模型截图</a></td><td>✅ 已整理</td></tr>
<tr><td>智谱 GLM</td><td><a href="providers/zhipu-glm/">目录</a></td><td><code>glm-4.7-flash</code>、<code>glm-4-flash-250414</code>、<code>glm-4.6v-flash</code>、<code>glm-4.1v-thinking-flash</code>、<code>glm-4v-flash</code>、<code>cogview-3-flash</code>、<code>cogvideox-flash</code></td><td><a href="https://open.bigmodel.cn/api/paas/v4/">API</a></td><td>-</td><td>✅ 已整理</td></tr>
<tr><td>商汤 SenseNova</td><td><a href="providers/sensenova/">目录</a></td><td><code>sensenova-6.7-flash-lite</code>、SenseNova U1 Fast、<code>glm-5.2</code>、DeepSeek V4 Flash</td><td><a href="https://token.sensenova.cn/v1">API</a></td><td><a href="providers/sensenova/assets/sensenova-model-docs.png">模型截图</a></td><td>✅ 已整理</td></tr>
<tr><td>OpenRouter</td><td><a href="providers/openrouter/">目录</a></td><td><code>*:free</code> 模型</td><td><a href="https://openrouter.ai/api/v1">API</a></td><td>-</td><td>✅ 已整理</td></tr>
<tr><td>B.AI</td><td><a href="providers/b-ai/">目录</a></td><td><code>DeepSeek-V4-Flash</code>、<code>DeepSeek-V4-Flash-Vision-Exp</code>、<code>Hy3</code></td><td><a href="https://api.b.ai/v1">API</a></td><td><a href="providers/b-ai/assets/b-ai-free-models.png">模型截图</a></td><td>✅ 已整理</td></tr>
<tr><td>NVIDIA NIM</td><td><a href="providers/nvidia-nim/">目录</a></td><td><code>z-ai/glm-5.2</code>、<code>poolside/laguna-xs-2.1</code>、<code>z-ai/glm-5.1</code>、<code>qwen/qwen3.5-397b-a17b</code></td><td><a href="https://integrate.api.nvidia.com/v1">API</a></td><td>-</td><td>🟡 待实测</td></tr>
<tr><td>ModelScope</td><td><a href="providers/modelscope/">目录</a></td><td><code>MiniMax/MiniMax-M2.5</code>、<code>qwen-qwen3-5-35b-a3b</code>、<code>qwen-qwen3-5-27b</code></td><td><a href="https://api-inference.modelscope.cn/v1">API</a></td><td>-</td><td>🟡 待实测</td></tr>
<tr><td>Groq</td><td><a href="providers/groq/">目录</a></td><td><code>moonshotai/kimi-k2-instruct</code>、<code>moonshotai/kimi-k2-instruct-0905</code>、<code>groq/compound</code></td><td><a href="https://api.groq.com/openai/v1">API</a></td><td>-</td><td>🟡 待实测</td></tr>
<tr><td>AI21 Labs</td><td><a href="providers/ai21-labs/">目录</a></td><td><code>jamba-large-1-7</code>、<code>jamba-mini-2</code></td><td><a href="https://api.ai21.com/studio/v1">API</a></td><td>-</td><td>🟡 待实测</td></tr>
<tr><td>模板</td><td><a href="providers/_template/">目录</a></td><td>待填写</td><td>待填写</td><td>-</td><td>✅ 可复制</td></tr>
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
