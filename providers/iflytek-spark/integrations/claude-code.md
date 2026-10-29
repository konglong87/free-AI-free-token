# Claude Code 接入 Spark-X2.6

## 当前状态

官方确认 Anthropic Messages 接口与 Claude Code 环境变量配置；官方接入示例使用 `spark-x2.5`，以下根据官方模型表替换为 `spark-x2.6`，**组合未持 Key 实测**。

## 配置（仅当前 shell）

```bash
export ANTHROPIC_BASE_URL="https://maas-api.cn-huabei-1.xf-yun.com/anthropic"
export ANTHROPIC_AUTH_TOKEN="<你的 MaaS APIKEY>"
export ANTHROPIC_MODEL="spark-x2.6"
export ANTHROPIC_DEFAULT_HAIKU_MODEL="spark-x2.6"
export CLAUDE_CODE_SUBAGENT_MODEL="spark-x2.6"
claude
```

也可按[官方指南](https://maas.xfyun.cn/doc/guide/4%E3%80%81%E4%B8%89%E6%96%B9%E5%B7%A5%E5%85%B7%E6%8E%A5%E5%85%A5/4.2%20%E6%8E%A5%E5%85%A5%20Claude%20Code.html)写入本地 `~/.claude/settings.json`；不要把真实 Key 写入仓库。Claude Code 会在上面的 Base URL 后拼接 `/v1/messages`，**不要**将 `/v1/messages` 提前加到 Base URL。

## 验证

启动后先试短问答，再依次验证流式、工具调用、长上下文和子 Agent。官方 [Anthropic 接口文档](https://maas.xfyun.cn/doc/guide/3%E3%80%81API%20%E6%8E%A5%E5%8F%A3/3.1%20%E6%98%9F%E7%81%AB%E6%A8%A1%E5%9E%8B%20API/3.1.3%20Anthropic.html)目前列出的单请求 Token 约束与模型标称 1M 不同，按接口限制配置，不保证 Claude Code 可直接使用 1M。
