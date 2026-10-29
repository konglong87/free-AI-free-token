# OpenCode 接入 Spark-X2.6

## 当前状态

待核验。讯飞 MaaS 提供 OpenAI-compatible Chat Completions，但截至 2026-10-29 未在讯飞公开接入指南中找到针对 OpenCode 的官方配置，也未在本仓库用 Key 实测。因此这里只列出已核对的接口字段，**不冒充已验证的 OpenCode 配置**。

## 可用于自定义 OpenAI-compatible Provider 的字段

| 字段 | 值 |
| --- | --- |
| Base URL | `https://maas-api.cn-huabei-1.xf-yun.com/v2` |
| API Key | 从 MaaS 控制台获取；不可提交到仓库 |
| Model ID | `spark-x2.6` |
| 预期路径 | `/chat/completions` |

下一步在当前 OpenCode 版本的自定义 Provider 配置中填写上述字段，先验证短问答，再验证流式输出和工具调用。OpenCode 的具体配置语法及完整兼容性仍待核验。

来源：[讯飞官方 API 接入文档](https://maas.xfyun.cn/doc/guide/1%E3%80%81%E5%BF%AB%E9%80%9F%E5%BC%80%E5%A7%8B/1.1%20API%E6%8E%A5%E5%85%A5.html)。
