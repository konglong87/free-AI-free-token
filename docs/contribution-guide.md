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

## 内容要求

只收集免费 LLM 接入所需信息，不收录无关介绍、营销内容或体验描述。

## 不要提交真实密钥

不要把 API key、token、cookie、控制台截图里的敏感字段提交到仓库。
