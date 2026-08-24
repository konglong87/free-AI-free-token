# B.AI 免费模型资料包 Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Add a fixed-layout B.AI provider directory documenting the three screenshot-confirmed free models with committed evidence images.

**Architecture:** Keep all B.AI facts inside `providers/b-ai/`, copied from the repository template. Use Markdown as the source of truth, explicitly separating screenshot-confirmed facts from unverified API/client capabilities; link the evidence images by relative paths.

**Tech Stack:** Markdown, PNG evidence files, Bash layout validator, Git.

---

### Task 1: Record the approved design

**Files:**
- Create: `docs/superpowers/specs/2026-08-24-b-ai-free-models-design.md`

- [x] **Step 1: Write the approved scope and verification design**

  Record the provider boundary, three model IDs, source URLs, evidence policy, explicit non-goals, and validation commands.

- [x] **Step 2: Commit the design**

  Run `git add docs/superpowers/specs/2026-08-24-b-ai-free-models-design.md && git commit -m "docs: define b.ai free models scope"`.

### Task 2: Scaffold the provider directory

**Files:**
- Create: `providers/b-ai/README.md`
- Create: `providers/b-ai/official-links.md`
- Create: `providers/b-ai/models.md`
- Create: `providers/b-ai/api.md`
- Create: `providers/b-ai/integrations/claude-code.md`
- Create: `providers/b-ai/integrations/openclaw.md`
- Create: `providers/b-ai/integrations/opencode.md`
- Create: `providers/b-ai/integrations/codex.md`
- Create: `providers/b-ai/integrations/hermes.md`

- [x] **Step 1: Copy the fixed template layout**

  Create `providers/b-ai/` from `providers/_template/` so the required file set remains compatible with `scripts/validate-layout.sh`.

- [x] **Step 2: Replace placeholders with B.AI facts**

  Use `https://chat.b.ai/` as the official entry, `https://chat.b.ai/key` as the key/model source, `https://chat.b.ai/leaderboard` as the leaderboard source, and `https://api.b.ai/v1` as the screenshot-confirmed Base URL. List only `DeepSeek-V4-Flash`, `DeepSeek-V4-Flash-Vision-Exp`, and `Hy3`; mark all untested capabilities and client integrations as `待核验`.

### Task 3: Add evidence and index entry

**Files:**
- Create: `providers/b-ai/assets/b-ai-free-models.png`
- Create: `providers/b-ai/assets/b-ai-api-example.png`
- Modify: `README.md`

- [x] **Step 1: Copy the supplied PNG evidence**

  Store the model/key page screenshot as `b-ai-free-models.png` and the curl/python/nodejs example as `b-ai-api-example.png`; do not alter image pixels or add sensitive files.

- [x] **Step 2: Link evidence from provider docs**

  Reference both images from `README.md`, `official-links.md`, and `models.md` where the corresponding claim is made.

- [x] **Step 3: Add the root README row**

  Add B.AI to the free-provider table with the three model IDs, Base URL, and status `已整理，待实测`.

### Task 4: Verify and publish

**Files:**
- Verify: all files above

- [x] **Step 1: Run structural validation**

  Run `./scripts/validate-layout.sh`; expected output is `Provider layout OK`.

- [x] **Step 2: Verify evidence files and content**

  Run `file providers/b-ai/assets/*.png`, search for all three model names and `https://api.b.ai/v1`, and confirm no API key/token values are present.

- [x] **Step 3: Review the diff and commit**

  Run `git diff --check` and `git status --short`, then commit with `git commit -m "docs: add b.ai free models"`.

- [x] **Step 4: Push the current branch**

  Run `git push origin main` and confirm the push succeeds.
