# Hermes 接入 K-EXAONE 2.5

## 当前状态

EXAONE 官方 Agent 接入页提供 Hermes 配置，以下照官方示例整理，**未持 Key 实测**。

## 接入步骤

1. 安装 Hermes（Linux / macOS / WSL2）：`curl -fsSL https://hermes-agent.nousresearch.com/install.sh | bash`
2. 在运行 Hermes 的终端设置 Key：`export EXAONE_API_KEY="<你的 EXAONE API Key>"`
3. 把以下内容合并到本地 `~/.hermes/config.yaml` 的 `providers` 和 `agent` 中（保留已有配置）：

```yaml
providers:
  exaone:
    api: https://api.exaone.ai/v1
    key_env: EXAONE_API_KEY
    transport: chat_completions
    default_model: k-exaone_v2.5
    extra_body:
      temperature: 1.0
      top_p: 0.95
agent:
  reasoning_effort: medium
```

4. 运行 `hermes model` 选择 `exaone` provider 和模型，再在项目目录运行 `hermes`。会话中可用 `/model custom:exaone:k-exaone_v2.5` 切换。

`agent.reasoning_effort` 可选 `low` / `medium` / `high` / `max`。Key 30 天到期后需重新签发并更新环境变量。

来源：[EXAONE 官方 Agent 接入页](https://docs.exaone.ai/ko/docs/agent-start)。
