# OpenClaw 接入 硅基流动 SiliconFlow（国内站）

## 当前状态

待核验。官方有接入页（见来源），本目录未逐页核对内容，也未实测。

以下按 OpenAI 兼容接口字段整理，具体配置格式以官方 OpenClaw 接入页为准。

## 可用于配置的字段

| 字段 | 值 |
| --- | --- |
| API Key | 本人的 SiliconFlow key；不可提交到仓库 |
| OpenAI Base URL | `https://api.siliconflow.cn/v1` |
| Model ID | `Qwen/Qwen3-8B`（示例，免费状态以价格页为准） |

## 说明与待核验

- 官方 OpenClaw 接入页的配置格式、推荐模型未核对。
- 不要把 `/chat/completions` 写进 Base URL。
- 先用短问答验证，再测试流式、工具调用；注意免费模型限速。

## 验证记录

| 日期 | OpenClaw 版本 | 模型 | 结果 | 备注 |
| --- | --- | --- | --- | --- |
| 2026-11-26 | 待核验 | 待核验 | 待核验 | 仅整理官方文档，未持有 key，未实测 |

## 来源

- [官方 OpenClaw 接入页](https://docs.siliconflow.cn/docs/usercases/use-siliconcloud-in-OpenClaw)
