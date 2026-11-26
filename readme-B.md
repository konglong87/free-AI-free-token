<div align="center">

# free-AI-free-token

免费 LLM 供应商、模型与 API 接入资料库

<!-- providers-count:start --><a href="https://github.com/konglong87/free-AI-free-token"><img src="https://img.shields.io/badge/providers-17-2563eb?style=flat-square" alt="供应商数量"></a><!-- providers-count:end -->
<a href="https://github.com/konglong87/free-AI-free-token/tree/main/providers"><img src="https://img.shields.io/badge/models-curated-059669?style=flat-square" alt="模型清单"></a>
<a href="https://github.com/konglong87/free-AI-free-token"><img src="https://img.shields.io/badge/docs-Markdown-f59e0b?style=flat-square" alt="文档格式"></a>

</div>

## 供应商总览

<table width="100%">
<tr>
<td width="25%"><strong><!-- providers-count:start -->17<!-- providers-count:end --></strong><br><sub>供应商目录</sub></td>
<td width="25%"><strong>OpenAI-compatible</strong><br><sub>优先收录</sub></td>
<td width="25%"><strong>5</strong><br><sub>固定接入文档</sub></td>
<td width="25%"><strong><!-- updated:start -->2026-11-26<!-- updated:end --></strong><br><sub>最后整理</sub></td>
</tr>
</table>

## 从这里开始

| 你要做什么 | 直接入口 |
| --- | --- |
| 找免费模型 | [供应商目录](#供应商目录) |
| 查 Base URL 和鉴权 | 对应目录的 `api.md` |
| 配置 Claude Code / OpenClaw / OpenCode | 对应目录的 `integrations/` |
| 新增一家供应商 | [贡献指南](docs/contribution-guide.md) |

## 供应商目录

按左侧标签找能力，再看右侧对应的模型 ID；点击供应商名称进入目录。

**理解 ≠ 生成**：👁 图片理解是看图回答；🖼 文生图 / 图生图是生成图片；🎬 文生视频 / 图生视频是生成视频。[查看标签说明](docs/model-capabilities.md)。

标签只汇总仓库已收录的模型，不代表平台全部能力或永久免费承诺；免费条件、限流和核验状态以供应商目录及官方资料为准。标注「待核验」的能力不作已确认推荐。

<!-- providers:start -->
<table width="100%">
<thead>
<tr><th width="30%">供应商 · 免费能力</th><th width="70%">代表性模型 · 接入</th></tr>
</thead>
<tbody>
<tr>
<td valign="top"><a href="providers/zhipu-glm/"><strong>智谱 GLM</strong></a><br><sub>📝 文本生成<br>👁 图片理解<br>🖼 文生图<br>🎬 视频生成</sub></td>
<td valign="top"><strong>文本生成</strong><br><code>glm-4.7-flash</code><br><br><strong>图片理解</strong><br><code>glm-4.6v-flash</code><br><br><strong>文生图</strong><br><code>cogview-3-flash</code><br><br><strong>视频生成</strong><br><code>cogvideox-flash</code><br><br><a href="https://open.bigmodel.cn/api/paas/v4">API</a> · <a href="https://bigmodel.cn/">注册</a> · <a href="providers/zhipu-glm/models.md">完整模型</a><br><sub>状态：✅</sub></td>
</tr>
<tr>
<td valign="top"><a href="providers/sensenova/"><strong>商汤 SenseNova</strong></a><br><sub>📝 文本生成<br>👁 图片理解</sub><br><br><sub>免费方式：限时活动</sub></td>
<td valign="top"><strong>文本生成 · 图片理解</strong><br><strong>DeepSeek V4.1 Flash</strong><br><code>deepseek-flash</code><br><strong><code>kimi-k3</code></strong><br><br><a href="https://token.sensenova.cn/v1">API</a> · <a href="https://www.sensenova.cn/">注册</a> · <a href="providers/sensenova/models.md">完整模型</a><br><sub>状态：✅ · <a href="https://platform.sensenova.cn/docs">官方文档</a></sub></td>
</tr>
<tr>
<td valign="top"><a href="providers/iflytek-spark/"><strong>科大讯飞星火</strong></a><br><sub>📝 文本生成</sub><br><br><sub>免费方式：一次性赠金</sub></td>
<td valign="top"><strong>文本生成 · 10 亿 Token 限时赠额</strong><br><code>spark-x2.6</code><br><br><a href="https://maas-api.cn-huabei-1.xf-yun.com/v2">API</a> · <a href="https://maas.xfyun.cn/">领取入口</a> · <a href="providers/iflytek-spark/models.md">完整模型</a><br><sub>状态：🟡 官方活动与接口已核对，领取条件 / 调用待实测 · <a href="providers/iflytek-spark/official-links.md">官方来源</a></sub></td>
</tr>
<tr>
<td valign="top"><a href="providers/volcengine-ark/"><strong>火山方舟（豆包）</strong></a><br><sub>📝 文本生成<br>👁 图片理解<br>（待核验）</sub><br><br><sub>免费方式：一次性赠金</sub></td>
<td valign="top"><strong>文本生成 · 图片理解 · 新用户一次性额度，覆盖范围待核验（待核验）</strong><br><code>doubao-seed-2-0-mini-260428</code><br><br><a href="https://ark.cn-beijing.volces.com/api/v3">API</a> · <a href="https://ark.volcengine.com/region:cn-beijing/apikey">注册</a> · <a href="providers/volcengine-ark/models.md">完整模型</a><br><sub>状态：🟡 官方文档已核对，控制台额度与调用待实测 · <a href="https://docs.volcengine.com/docs/ark/free-inference-quota">免费额度说明</a> · <a href="providers/volcengine-ark/official-links.md">官方来源</a></sub></td>
</tr>
<tr>
<td valign="top"><a href="providers/siliconflow/"><strong>硅基流动（国内站）</strong></a><br><sub>📝 文本生成<br>👁 图片理解</sub><br><br><sub>免费方式：长期免费<br>门槛：无需绑卡 · 需实名</sub></td>
<td valign="top"><strong>文本生成 · 永久免费，需实名认证</strong><br><code>Qwen/Qwen3-8B</code><br><br><strong>文本生成 · 图片理解</strong><br><code>Qwen/Qwen3.5-4B</code><br><br><a href="https://api.siliconflow.cn/v1">API</a> · <a href="https://cloud.siliconflow.cn">注册</a> · <a href="providers/siliconflow/models.md">完整模型</a><br><sub>状态：🟡 官方页面已核对，免费模型 ID 与限速待登录实测 · <a href="https://www.siliconflow.cn/pricing">价格页（免费标记）</a> · <a href="providers/siliconflow/official-links.md">官方来源</a></sub></td>
</tr>
<tr>
<td valign="top"><a href="providers/modelscope/"><strong>ModelScope</strong></a><br><sub>📝 文本生成</sub></td>
<td valign="top"><strong>文本生成</strong><br><code>MiniMax/MiniMax-M2.5</code><br><code>qwen-qwen3-5-35b-a3b</code><br><br><a href="https://api-inference.modelscope.cn/v1">API</a> · <a href="https://modelscope.cn/">注册</a> · <a href="providers/modelscope/models.md">完整模型</a><br><sub>状态：🟡</sub></td>
</tr>
<tr>
<td valign="top"><a href="providers/agnes/"><strong>Agnes AI</strong></a><br><sub>📝 文本生成<br>👁 图片理解<br>🖼 文生图<br>🖼 图生图<br>🎬 文生视频<br>🎬 图生视频</sub></td>
<td valign="top"><strong>文本生成 · 图片理解</strong><br><code>agnes-3.0-flash</code><br><br><strong>文生图 · 图生图</strong><br><code>agnes-image-2.5-flash</code><br><br><strong>文生视频 · 图生视频</strong><br><code>agnes-video-2.5-flash</code><br><br><a href="https://apihub.agnes-ai.com/v1">API</a> · <a href="https://agnes-ai.com/">注册</a> · <a href="providers/agnes/models.md">完整模型</a><br><sub>状态：✅ · <a href="https://www.agnes-ai.com/zh-Hans/docs/pricing">官方定价</a></sub></td>
</tr>
<tr>
<td valign="top"><a href="providers/lg-exaone/"><strong>LG EXAONE</strong></a><br><sub>📝 文本生成</sub><br><br><sub>免费方式：限时活动</sub></td>
<td valign="top"><strong>文本生成 · 免费（非商用，Key 30 天有效）</strong><br><code>k-exaone_v2.5</code><br><br><a href="https://api.exaone.ai/v1">API</a> · <a href="https://docs.exaone.ai">申请入口</a> · <a href="providers/lg-exaone/models.md">完整模型</a><br><sub>状态：🟡 官方公告、接口与条款已核对，申请 / 调用待实测 · <a href="providers/lg-exaone/official-links.md">官方来源</a></sub></td>
</tr>
<tr>
<td valign="top"><a href="providers/openrouter/"><strong>OpenRouter</strong></a><br><sub>📝 文本生成</sub><br><br><sub>免费方式：每日额度</sub></td>
<td valign="top"><strong>文本生成</strong><br><code>*:free</code> 模型（能力随具体模型变化）<br><br><a href="https://openrouter.ai/api/v1">API</a> · <a href="https://openrouter.ai/">注册</a> · <a href="providers/openrouter/models.md">完整模型</a><br><sub>状态：✅</sub></td>
</tr>
<tr>
<td valign="top"><a href="providers/b-ai/"><strong>B.AI</strong></a><br><sub>📝 文本生成<br>👁 图片理解<br>（待核验）</sub></td>
<td valign="top"><strong>文本生成</strong><br><code>DeepSeek-V4-Flash</code><br><code>Hy3</code><br><br><strong>文本生成 · 图片理解（待核验）</strong><br><code>DeepSeek-V4-Flash-Vision-Exp</code><br><br><a href="https://api.b.ai/v1">API</a> · <a href="https://chat.b.ai/">注册</a> · <a href="providers/b-ai/models.md">完整模型</a><br><sub>状态：✅ · <a href="providers/b-ai/assets/b-ai-free-models.png">模型截图</a></sub></td>
</tr>
<tr>
<td valign="top"><a href="providers/nvidia-nim/"><strong>NVIDIA NIM</strong></a><br><sub>📝 文本生成</sub></td>
<td valign="top"><strong>文本生成</strong><br><code>z-ai/glm-5.2</code><br><code>poolside/laguna-xs-2.1</code><br><br><a href="https://integrate.api.nvidia.com/v1">API</a> · <a href="https://build.nvidia.com/">注册</a> · <a href="providers/nvidia-nim/models.md">完整模型</a><br><sub>状态：🟡</sub></td>
</tr>
<tr>
<td valign="top"><a href="providers/groq/"><strong>Groq</strong></a><br><sub>📝 文本生成</sub></td>
<td valign="top"><strong>文本生成</strong><br><code>moonshotai/kimi-k2-instruct</code><br><code>groq/compound</code><br><br><a href="https://api.groq.com/openai/v1">API</a> · <a href="https://console.groq.com/">注册</a> · <a href="providers/groq/models.md">完整模型</a><br><sub>状态：🟡</sub></td>
</tr>
<tr>
<td valign="top"><a href="providers/ai21-labs/"><strong>AI21 Labs</strong></a><br><sub>📝 文本生成</sub></td>
<td valign="top"><strong>文本生成</strong><br><code>jamba-large-1-7</code><br><code>jamba-mini-2</code><br><br><a href="https://api.ai21.com/studio/v1">API</a> · <a href="https://studio.ai21.com/">注册</a> · <a href="providers/ai21-labs/models.md">完整模型</a><br><sub>状态：🟡</sub></td>
</tr>
<tr>
<td valign="top"><a href="providers/cerebras/"><strong>Cerebras</strong></a><br><sub>📝 文本生成<br>👁 图片理解</sub><br><br><sub>免费方式：一次性赠金<br>门槛：需绑卡</sub></td>
<td valign="top"><strong>文本生成 · $5 一次性额度，30 天过期，需绑定支付方式</strong><br><code>gpt-oss-120b</code><br><br><strong>文本生成 · 图片理解</strong><br><code>qwen-3.8-27b</code><br><br><a href="https://api.cerebras.ai/v1">API</a> · <a href="https://cloud.cerebras.ai">注册</a> · <a href="providers/cerebras/models.md">完整模型</a><br><sub>状态：🟡 官方文档已核对，注册与调用待实测 · <a href="https://inference-docs.cerebras.ai/support/rate-limits">限流说明</a> · <a href="providers/cerebras/official-links.md">官方来源</a></sub></td>
</tr>
<tr>
<td valign="top"><a href="providers/google-ai-studio/"><strong>Google AI Studio（Gemini）</strong></a><br><sub>📝 文本生成<br>👁 图片理解</sub><br><br><sub>免费方式：长期免费<br>门槛：无需绑卡 · 有地区限制</sub></td>
<td valign="top"><strong>文本生成 · 图片理解 · 永久免费层，数据用于改进产品</strong><br><code>gemini-4-hydrogen</code><br><br><strong>文本生成</strong><br><code>gemini-3.8-flash</code><br><br><a href="https://generativelanguage.googleapis.com/v1beta/openai/">API</a> · <a href="https://aistudio.google.com">注册</a> · <a href="providers/google-ai-studio/models.md">完整模型</a><br><sub>状态：🟡 官方文档已核对，免费层数值限额与调用待实测 · <a href="https://aistudio.google.com/docs/pricing">免费层定价</a> · <a href="providers/google-ai-studio/official-links.md">官方来源</a></sub></td>
</tr>
<tr>
<td valign="top"><a href="providers/cloudflare-workers-ai/"><strong>Cloudflare Workers AI</strong></a><br><sub>📝 文本生成<br>👁 图片理解</sub><br><br><sub>免费方式：每日额度</sub></td>
<td valign="top"><strong>文本生成 · 每日 10,000 Neurons 免费</strong><br><code>@cf/openai/gpt-oss-120b</code><br><br><strong>文本生成 · 图片理解</strong><br><code>@cf/google/gemma-4-26b-a4b-it</code><br><br><a href="https://api.cloudflare.com/client/v4/accounts/{account_id}/ai/v1">API</a> · <a href="https://dash.cloudflare.com">注册</a> · <a href="providers/cloudflare-workers-ai/models.md">完整模型</a><br><sub>状态：🟡 官方文档已核对，账号注册与调用待实测 · <a href="https://developers.cloudflare.com/workers-ai/platform/pricing/">免费额度说明</a> · <a href="providers/cloudflare-workers-ai/official-links.md">官方来源</a></sub></td>
</tr>
<tr>
<td valign="top"><a href="providers/cohere/"><strong>Cohere</strong></a><br><sub>📝 文本生成<br>👁 图片理解</sub><br><br><sub>免费方式：长期免费</sub></td>
<td valign="top"><strong>文本生成 · Trial key 每月 1,000 次，不得商用</strong><br><code>command-a-03-2025</code><br><br><strong>文本生成 · 图片理解</strong><br><code>command-a-vision-07-2025</code><br><br><a href="https://api.cohere.ai/compatibility/v1">API</a> · <a href="https://dashboard.cohere.com/api-keys">注册</a> · <a href="providers/cohere/models.md">完整模型</a><br><sub>状态：🟡 官方文档已核对，注册与调用待实测；Trial key 不得商用 · <a href="https://docs.cohere.com/docs/rate-limits">限流说明</a> · <a href="providers/cohere/official-links.md">官方来源</a></sub></td>
</tr>
</tbody>
</table>
<!-- providers:end -->

## 证据图片

<table width="100%">
<tr>
<td align="center" width="33%"><img src="providers/agnes/assets/agnes-free-models.png" width="220" alt="Agnes AI 免费模型证据截图"><br><sub>Agnes AI</sub></td>
<td align="center" width="33%"><img src="providers/sensenova/assets/sensenova-model-docs.png" width="220" alt="SenseNova 模型文档证据截图"><br><sub>SenseNova</sub></td>
<td align="center" width="33%"><img src="providers/b-ai/assets/b-ai-free-models.png" width="220" alt="B.AI 免费模型证据截图"><br><sub>B.AI</sub></td>
</tr>
</table>

## 目录结构

```text
providers/<provider-slug>/
├── README.md
├── official-links.md
├── models.md
├── api.md
└── integrations/
    ├── claude-code.md
    ├── openclaw.md
    ├── opencode.md
    ├── codex.md
    └── hermes.md
```

每家供应商独立维护，详细模型能力、限流和核验状态以目录内文档为准。未实测能力统一标记为 `待核验`，不提交真实 API key、token、cookie 或账号信息。

## 新增供应商

```bash
cp -R providers/_template providers/<provider-slug>
# 改 providers/<provider-slug>/provider.yaml，并填写固定文档
cd tools/catalog && go run . && cd ../..   # 生成两份首页的供应商表
./scripts/validate-layout.sh
```

供应商表由 `provider.yaml` 生成，不要手改。来源、核验日期和状态必须可追溯。

## 文档

- [架构说明](docs/architecture.md)
- [贡献指南](docs/contribution-guide.md)
- [供应商目录](providers/)
