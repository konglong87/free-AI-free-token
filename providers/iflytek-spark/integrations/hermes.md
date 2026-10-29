# Hermes 接入 Spark-X2.6

## 当前状态

讯飞官方提供 Hermes Custom endpoint 示例，示例模型为 `spark-x2.5`；替换为官方模型表中的 `spark-x2.6`，**待实际调用验证**。

## 接入步骤

1. 本地运行 `hermes setup`，选择 **Quick Setup** 和 **Custom endpoint / OpenAI-compatible**。
2. Base URL：`https://maas-api.cn-huabei-1.xf-yun.com/v2`。
3. APIKEY：本人的 MaaS Key；Model ID：`spark-x2.6`；Context length：按控制台及实际接口授权填写，不要直接假定全链路支持 1M。
4. 执行 `hermes chat`，先验证短问答，再验证工具调用；如需手动切换模型，官方提供 `/model custom:spark-x2.5` 示例，X2.6 对应写成 `/model custom:spark-x2.6`（待实测）。

密钥建议通过 Hermes 的本地凭证流程保存，不要放进共享的 `config.yaml`。

来源：[讯飞官方 Hermes 指南](https://maas.xfyun.cn/doc/guide/4%E3%80%81%E4%B8%89%E6%96%B9%E5%B7%A5%E5%85%B7%E6%8E%A5%E5%85%A5/4.5%20%E6%8E%A5%E5%85%A5%20Hermes.html)。
