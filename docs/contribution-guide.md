# Contribution Guide

## 新增供应商

```bash
cp -R providers/_template providers/<provider-slug>
```

然后按顺序填写：

1. `official-links.md`
2. `models.md`
3. `api.md`
4. `integrations/*`
5. `README.md`
6. 根目录 `README.md` 的免费供应商入口

## 必填检查

新增或修改供应商后运行：

```bash
./scripts/validate-layout.sh
```

校验只检查结构，不代表内容真实。内容真实性必须靠官方来源、原文来源或实际接入结果。

## 来源要求

每个关键结论都应该能追溯到官方来源，尤其是：

- 免费模型
- 是否 OpenAI compatible
- base URL
- model name
- agent / coding client 接入方式
- 核验状态

## 内容要求

只收集免费 LLM 接入所需信息，不收录无关介绍、营销内容或体验描述。

允许写入的信息：

- 官网和官方文档入口
- API key 获取入口
- 接入地址
- 免费模型
- OpenAI-compatible 状态
- Claude Code、OpenClaw、OpenCode、Codex、Hermes 接入方式
- 核验日期和核验状态

禁止写入的信息：

- 真实 API key、token、cookie、session
- 个人账号、手机号、邮箱、用户 ID
- 作者或转载平台的无关信息
- 截图中的敏感字段
- 无法追溯来源的结论

## 不要提交真实密钥

不要把 API key、token、cookie、控制台截图里的敏感字段提交到仓库。

## 首页能力标签

更新免费模型时，同步维护 `README.md` 与 `readme-B.md` 的供应商能力概览和模型分组，使用[统一能力标签](model-capabilities.md)。

- 先在供应商 `models.md` 记录能力与来源，再汇总到首页；不要根据模型 ID 猜测。
- 区分图片理解、图片生成与视频生成；未确认的输入方式不要补写。
- 能力待核验时，供应商标签与模型分组都必须保留「待核验」。
- 保持两列表格，接入与证据链接放在模型下方；不要继续增加横向列。
- 修改后除结构校验外，还要检查桌面与窄屏渲染截图，确保标签、完整模型 ID 和链接不被横向裁切。
