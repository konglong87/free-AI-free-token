# OpenCode 接入 K-EXAONE 2.5

## 当前状态

EXAONE 官方 Agent 接入页提供 OpenCode 配置，以下照官方示例整理，**未持 Key 实测**。

## 接入步骤

1. 按官方页面安装 OpenCode：`npm install -g @opencode/cli`
2. 在运行 OpenCode 的终端设置 Key：`export EXAONE_API_KEY="<你的 EXAONE API Key>"`
3. 把以下内容合并到项目的 `opencode.json`：

```json
{
  "$schema": "https://opencode.ai/config.json",
  "model": "exaone/k-exaone_v2.5",
  "providers": {
    "exaone": {
      "package": "@opencode/ai/providers/openai-compatible",
      "name": "EXAONE",
      "settings": {
        "baseURL": "https://api.exaone.ai/v1",
        "apiKey": "{env:EXAONE_API_KEY}"
      },
      "models": {
        "k-exaone_v2.5": {
          "name": "K-EXAONE 2.5",
          "capabilities": {
            "tools": true,
            "input": ["text"],
            "output": ["text"],
            "temperature": true,
            "reasoning": { "supported": true }
          },
          "compatibility": { "reasoningField": "reasoning_content" },
          "body": {
            "reasoning_effort": "medium",
            "temperature": 1.0,
            "top_p": 0.95
          }
        }
      }
    }
  }
}
```

4. 在项目目录运行 `opencode --standalone`（使用当前终端的环境变量），会话中也可用 `/models` 切换。

`compatibility.reasoningField` 让 OpenCode 把 `reasoning_content` 带回下一轮请求；`body.reasoning_effort` 可选 `low` / `medium` / `high` / `max`。Key 只通过 `{env:EXAONE_API_KEY}` 引用，不要写死在 `opencode.json`。

来源：[EXAONE 官方 Agent 接入页](https://docs.exaone.ai/ko/docs/agent-start)。
