# Models

## 免费模型

| 模型名 | 类型 | 上下文 | 备注 |
| --- | --- | --- | --- |
| `cohere/north-mini-code:free` | 文本 / 代码 | 256000 | 官方模型接口显示免费 |
| `qwen/qwen3-coder:free` | 文本 / 代码 | 1048576 | 官方模型接口显示免费 |
| `openai/gpt-oss-120b:free` | 文本 | 131072 | 官方模型接口显示免费 |
| `openai/gpt-oss-20b:free` | 文本 | 131072 | 官方模型接口显示免费 |
| `meta-llama/llama-3.3-70b-instruct:free` | 文本 | 131072 | 官方模型接口显示免费 |
| `meta-llama/llama-3.2-3b-instruct:free` | 文本 | 131072 | 官方模型接口显示免费 |
| `nousresearch/hermes-3-llama-3.1-405b:free` | 文本 | 131072 | 官方模型接口显示免费 |
| `nvidia/nemotron-3-ultra-550b-a55b:free` | 文本 | 1000000 | 官方模型接口显示免费 |

## 模型选择建议

- 默认模型：从 `https://openrouter.ai/models?max_price=0` 选择当前可用 `:free` 模型。
- 编程推荐模型：`cohere/north-mini-code:free`、`qwen/qwen3-coder:free`。
- 长上下文模型：`qwen/qwen3-coder:free`、`nvidia/nemotron-3-ultra-550b-a55b:free`。

## 模型来源

- 官方模型列表：https://openrouter.ai/models?max_price=0
- 官方模型接口：https://openrouter.ai/api/v1/models
- 最后核验日期：2026-06-24

## 说明

OpenRouter 免费模型通常带 `:free` 后缀。免费模型会变化，接入前以官方模型列表和模型接口为准。
