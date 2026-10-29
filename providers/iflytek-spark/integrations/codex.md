# Codex CLI 接入 Spark-X2.6

## 当前状态

讯飞官方有 Codex Responses 接入指南；其示例使用 `spark-x2.5`，这里替换为官方确认的 `spark-x2.6`，**实际调用待实测**。

## 配置

在本地 shell 设置 API Key，不提交到版本库：

```bash
export SPARK_API_KEY="<你的 MaaS APIKEY>"
```

将以下配置合并到本地 `~/.codex/config.toml`（保留已有其他 provider）：

```toml
model = "spark-x2.6"
model_provider = "spark"

[model_providers.spark]
name = "Spark X2.6"
base_url = "https://maas-api.cn-huabei-1.xf-yun.com/v1"
wire_api = "responses"
env_key = "SPARK_API_KEY"
```

Codex 在 Base URL 后请求 `/responses`，最终地址是 `https://maas-api.cn-huabei-1.xf-yun.com/v1/responses`，不能误填 Chat Completions 的 `/v2`。运行 `codex` 后先验证文本响应，再验证流式与工具调用。

来源：[讯飞官方 Codex 接入指南](https://maas.xfyun.cn/doc/guide/4%E3%80%81%E4%B8%89%E6%96%B9%E5%B7%A5%E5%85%B7%E6%8E%A5%E5%85%A5/4.3%20%E6%8E%A5%E5%85%A5%20Codex.html)、[MaaS API 接入地址](https://maas.xfyun.cn/doc/guide/1%E3%80%81%E5%BF%AB%E9%80%9F%E5%BC%80%E5%A7%8B/1.1%20API%E6%8E%A5%E5%85%A5.html)。
