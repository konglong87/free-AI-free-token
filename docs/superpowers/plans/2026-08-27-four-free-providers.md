# Four Free Providers Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Add NVIDIA NIM, ModelScope, Groq, and AI21 Labs as template-compliant free-model provider directories with root README navigation.

**Architecture:** Keep each provider isolated under `providers/<slug>/` and use Markdown as the source of truth. Copy the existing `_template` layout, populate provider facts and official links, and keep the root README limited to summary rows.

**Tech Stack:** Markdown, Bash (`scripts/validate-layout.sh`), Git.

---

### Task 1: Scaffold Provider Directories

**Files:**
- Create: `providers/nvidia-nim/` from `providers/_template/`
- Create: `providers/modelscope/` from `providers/_template/`
- Create: `providers/groq/` from `providers/_template/`
- Create: `providers/ai21-labs/` from `providers/_template/`

- [ ] **Step 1: Copy the fixed template layout**

Run:

```bash
for provider in nvidia-nim modelscope groq ai21-labs; do
  cp -R providers/_template "providers/$provider"
done
```

Expected: each new directory contains the four top-level Markdown files and five files under `integrations/`.

- [ ] **Step 2: Verify the scaffold before content edits**

Run `./scripts/validate-layout.sh`.

Expected: `Provider layout OK`.

### Task 2: Populate Provider Documentation

**Files:**
- Modify: all Markdown files under `providers/nvidia-nim/`
- Modify: all Markdown files under `providers/modelscope/`
- Modify: all Markdown files under `providers/groq/`
- Modify: all Markdown files under `providers/ai21-labs/`

- [ ] **Step 1: Replace template placeholders with provider metadata**

For every provider, set the status date to `2026-08-27`, add official homepage, key/token page, documentation and model source links, and use the approved Base URL.

- [ ] **Step 2: Add the approved model tables and guidance**

Use these exact model IDs:

```text
NVIDIA NIM: z-ai/glm-5.2, poolside/laguna-xs-2.1, z-ai/glm-5.1, qwen/qwen3.5-397b-a17b
ModelScope: MiniMax/MiniMax-M2.5, qwen-qwen3-5-35b-a3b, qwen-qwen3-5-27b
Groq: moonshotai/kimi-k2-instruct, moonshotai/kimi-k2-instruct-0905, groq/compound
AI21 Labs: jambalarge-1-7, jambamini-2
```

Mark aggregator-only context, rate limits, `/models`, streaming, tool calling, structured output, and client-specific compatibility as `待核验` unless an official source explicitly confirms them. Explain in AI21 Labs documents that the current reference list confirms only two free models.

- [ ] **Step 3: Fill all five integration documents per provider**

Each integration includes the provider API key variable, Base URL, a representative model, setup steps, and a verification table with status `待核验`; do not claim real-key testing.

- [ ] **Step 4: Scan provider docs for unresolved template placeholders**

Run:

```bash
rg -n "待填写|YYYY-MM-DD|这里写供应商|PROVIDER_BASE_URL=\"待填写\"|PROVIDER_MODEL=\"待填写\"" providers/nvidia-nim providers/modelscope providers/groq providers/ai21-labs
```

Expected: no output.

### Task 3: Update Root README Index

**Files:**
- Modify: `README.md` in the `免费供应商` HTML table

- [ ] **Step 1: Add four summary rows**

Add one row per provider with directory link, approved representative model IDs, Base URL, and status wording consistent with existing entries. Keep details in provider directories.

- [ ] **Step 2: Verify root links and model IDs**

Run:

```bash
rg -n "providers/(nvidia-nim|modelscope|groq|ai21-labs)|z-ai/glm-5.2|MiniMax/MiniMax-M2.5|moonshotai/kimi-k2-instruct|jambalarge-1-7" README.md
```

Expected: each provider path and representative model appears in the table.

### Task 4: Validate, Review, and Commit

**Files:**
- Verify: all new provider Markdown files and `README.md`

- [ ] **Step 1: Run structural and whitespace checks**

Run `./scripts/validate-layout.sh` and `git diff --check`.

Expected: `Provider layout OK` and no diff-check errors.

- [ ] **Step 2: Scan sensitive values and endpoint consistency**

Run:

```bash
rg -n "sk-[A-Za-z0-9]{20,}|Bearer [A-Za-z0-9_-]{20,}|token=[A-Za-z0-9_-]{20,}|session" providers/nvidia-nim providers/modelscope providers/groq providers/ai21-labs
rg -n "https://(integrate\.api\.nvidia\.com/v1|api-inference\.modelscope\.cn/v1|api\.groq\.com/openai/v1|api\.ai21\.com/studio/v1)" providers/nvidia-nim providers/modelscope providers/groq providers/ai21-labs
```

Expected: no credential-like values; expected Base URLs are present.

- [ ] **Step 3: Review the final diff**

Run `git diff --stat` and inspect the changed provider files and `README.md`; confirm only approved documentation changes are present.

- [ ] **Step 4: Stage implementation files and approved planning documents**

Run:

```bash
git add README.md providers/nvidia-nim providers/modelscope providers/groq providers/ai21-labs docs/superpowers/plans/2026-08-27-four-free-providers.md
git add docs/superpowers/specs/2026-08-27-four-free-providers-design.md
git commit -m "docs: add nvidia modelscope groq ai21 providers"
git push
```

Expected: commit succeeds and the current branch is pushed with both planning documents and the four provider additions.
