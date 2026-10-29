# OpenClaw 接入 Spark-X2.6

## 当前状态

讯飞官方提供 OpenClaw 自定义 Provider 配置，示例模型为 `spark-x2.5`；以下将 ID 替换为 `spark-x2.6`，**未用实际账号验证**。

## 接入步骤

1. 本地运行 `openclaw onboard --install-daemon`，选择 **Custom Provider**。
2. Endpoint compatibility 选 **OpenAI-compatible (/chat/completions)**。
3. Base URL 填 `https://maas-api.cn-huabei-1.xf-yun.com/v2`；APIKEY 填本人的 MaaS Key；Model ID 填 `spark-x2.6`。
4. 先用短问答验证，随后再测试流式、Function Calling 与工具参数回填。

最终请求路径应为 `/v2/chat/completions`，不要把 `/chat/completions` 写入 Base URL。限流按 API Key + 模型计，429 时限制并发并退避。

来源：[讯飞官方 OpenClaw 指南](https://maas.xfyun.cn/doc/guide/4%E3%80%81%E4%B8%89%E6%96%B9%E5%B7%A5%E5%85%B7%E6%8E%A5%E5%85%A5/4.4%20%E6%8E%A5%E5%85%A5%20OpenClaw.html)。
