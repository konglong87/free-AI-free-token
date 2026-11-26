# Models

## 免费模型

| 模型名 | 类型 | 上下文 | 备注 |
| --- | --- | --- | --- |
| `inclusionai/ling-3.2-flash:free` | 文本生成 / 图片理解 | 1048576 | 官方模型接口显示免费，支持 tools 参数 |
| `baidu/ernie-5.5-flash-preview:free` | 文本生成 | 1024000 | 官方模型接口显示免费，支持 tools 参数 |
| `liquid/lfm-3-1.7b:free` | 文本生成 | 262144 | 官方模型接口显示免费，支持 tools 参数 |
| `particle-ai/particle-1:free` | 文本生成 / 图片理解 | 1048576 | 官方模型接口显示免费，支持 tools 参数 |
| `recursive/flint-v1.0:free` | 文本生成 | 262144 | 官方模型接口显示免费，支持 tools 参数 |
| `nvidia/nemotron-3.5-super-vl:free` | 文本生成 / 图片理解 | 1000000 | 官方模型接口显示免费，支持 tools 参数 |
| `nvidia/nemotron-3.5-lightning:free` | 文本生成 | 1000000 | 官方模型接口显示免费，支持 tools 参数 |
| `nvidia/nemotron-3-ultra-550b-a55b:free` | 文本生成 | 1000000 | 官方模型接口显示免费，支持 tools 参数 |
| `nvidia/nemotron-3-nano-omni-30b-a3b-reasoning:free` | 文本生成 / 图片理解 | 256000 | 官方模型接口显示免费，支持 tools 参数 |
| `google/gemma-4-26b-a4b-it:free` | 文本生成 / 图片理解 | 262144 | 官方模型接口显示免费，支持 tools 参数 |
| `google/gemma-4-31b-it:free` | 文本生成 / 图片理解 | 262144 | 官方模型接口显示免费，支持 tools 参数 |

## 模型选择建议

- 默认模型：从 `https://openrouter.ai/models?max_price=0` 选择当前可用 `:free` 模型。
- 编程推荐模型：`nvidia/nemotron-3-ultra-550b-a55b:free`（官方模型接口显示支持 tools 参数）。
- 长上下文模型：`inclusionai/ling-3.2-flash:free`、`particle-ai/particle-1:free`、`nvidia/nemotron-3-ultra-550b-a55b:free`。
- 免费限额：每分钟 20 次；累计购买不足 10 credits 时每天 50 次，购满 10 credits 后每天 1,000 次（[官方限额](https://openrouter.ai/docs/api-reference/limits)）。

## 模型来源

- 官方模型列表：https://openrouter.ai/models?max_price=0
- 官方模型接口：https://openrouter.ai/api/v1/models
- 最后核验日期：2026-11-26（模型清单按官方模型接口当天结果更新）

## 说明

OpenRouter 免费模型通常带 `:free` 后缀。免费模型会变化，接入前以官方模型列表和模型接口为准。
