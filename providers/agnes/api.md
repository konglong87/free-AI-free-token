# API

## 基本信息

- 接入地址：https://apihub.agnes-ai.com/v1
- Chat completions endpoint：`https://apihub.agnes-ai.com/v1/chat/completions`
- 鉴权方式：Bearer API key
- 是否 OpenAI-compatible：是

## 环境变量

```bash
export PROVIDER_API_KEY="你的 Agnes AI API key"
export PROVIDER_BASE_URL="https://apihub.agnes-ai.com/v1"
export PROVIDER_MODEL="agnes-3.0-flash"
```

## 接入步骤

1. 打开 https://agnes-ai.com/doc。
2. 注册或登录 Agnes AI。
3. 创建 API key。
4. 配置 Base URL 为 `https://apihub.agnes-ai.com/v1`。
5. 配置模型为 `agnes-3.0-flash`。
6. 发送最小对话请求验证。

## 常见错误

| 错误 | 可能原因 | 处理方式 |
| --- | --- | --- |
| 401 | API key 错误或未启用 | 重新生成 key，确认账号状态 |
| 403 | 模型不可用或账号受限 | 查看控制台和官方文档 |
| 429 | 限流 | 降低并发，等待限制恢复 |
| 404 | endpoint 或 model 错误 | 对照官方文档核验 |

## 官方来源

- 模型定价（文本、图片、视频）：https://www.agnes-ai.com/zh-Hans/docs/pricing
- Agnes Image 2.5 Flash：https://www.agnes-ai.com/zh-Hans/docs/agnes-image-25-flash
- Agnes Video 2.5 Flash：https://www.agnes-ai.com/zh-Hans/docs/agnes-video-25-flash
- Quick Start：https://agnes-ai.com/doc/quick-start
- Agnes 2.0 Flash：https://agnes-ai.com/doc/agnes-20-flash
- Codex++ 文档：https://agnes-ai.com/doc/cid8

## 图片生成 API

`agnes-image-2.5-flash` 使用图像生成专用端点，不使用上方的 Chat Completions endpoint：

- Endpoint：`POST https://apihub.agnes-ai.com/v1/images/generations`
- 当前价格：所有支持的输出分辨率（1K、2K、3K、4K）及输入参考图片均为 `$0`
- 文生图必需字段：`model`、`prompt`、`size`
- 图生图或多图合成：在 `extra_body.image` 中传入图片 URL 或 Base64 Data URI

## 视频生成 API

`agnes-video-2.5-flash` 使用异步视频任务接口：

- 创建任务：`POST https://apihub.agnes-ai.com/v1/videos`
- 查询任务：`GET https://apihub.agnes-ai.com/agnesapi?video_id=<VIDEO_ID>&model_name=agnes-video-2.5-flash`
- 模式：`text`（文生视频）、`keyframe`（首尾帧控制）、`reference`（图片或音频参考）
- 专属限制：`size` 固定为 `720P`；时长 `seconds` 使用字符串 `4`–`12`；`n` 仅支持 `1`
- 参考素材限制：最多 5 张图片、3 段音频；不支持参考视频。素材 URL 必须能被 Agnes AI 服务公开访问
- 计费：原价 `$0.025/秒`，当前限时免费（现价 `$0/秒`）；优惠期限以官方公告和账户账单为准
- 查询建议：所有模式都带 `video_id` 与 `model_name`；不带 `model_name` 的查询仅适用于 `text` 模式

两种生成接口的免费政策可能调整；真实使用时以账户账单与官方定价页为准。
