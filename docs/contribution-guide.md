# Contribution Guide

## 新增供应商

```bash
cp -R providers/_template providers/<provider-slug>
```

然后按顺序填写：

1. `provider.yaml`：供应商的唯一结构化数据来源（见下文）
2. `official-links.md`
3. `models.md`
4. `api.md`
5. `integrations/*`
6. `README.md`

不要手改首页的供应商表。改完 yaml 后运行生成，再运行校验：

```bash
cd tools/catalog && go run .   # 重新生成 README.md / readme-B.md 的供应商表和数量
./scripts/validate-layout.sh   # 结构校验；已安装 go 时同时运行 go run . -check
```

## provider.yaml

字段说明见 `providers/_template/provider.yaml`。要点：

- 数据只来自该供应商的官方来源或本目录已有资料；查不到的写 `null` 或「待核验」，不要猜测。
- `home.tags`、模型的 `capabilities` 必须使用[统一能力标签](model-capabilities.md)；`home.order` 决定首页顺序，不可重复。
- 模型按列表顺序展示，`home.pages` 决定出现在哪份首页；相邻且能力相同的模型会合并成一组。
- 生成器会校验必填字段、枚举值、能力标签、slug 与目录名一致、同一家里模型 ID 不重复。

## 必填检查

新增或修改供应商后运行 `./scripts/validate-layout.sh`。

校验只检查结构和数据格式，不代表内容真实。内容真实性必须靠官方来源、原文来源或实际接入结果。

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

更新免费模型时，在 `provider.yaml` 里维护能力标签和模型分组，再运行生成，同步更新 `README.md` 与 `readme-B.md`，使用[统一能力标签](model-capabilities.md)。

- 先在供应商 `models.md` 记录能力与来源，再写入 `provider.yaml`；不要根据模型 ID 猜测。
- 区分图片理解、图片生成与视频生成；未确认的输入方式不要补写。
- 能力待核验时，供应商标签与模型分组都必须保留「待核验」。
- 保持两列表格，接入与证据链接放在模型下方；不要继续增加横向列。
- 修改后除结构校验外，还要检查桌面与窄屏渲染截图，确保标签、完整模型 ID 和链接不被横向裁切。
