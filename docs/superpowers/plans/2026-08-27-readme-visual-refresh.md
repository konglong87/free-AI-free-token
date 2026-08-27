# README Visual Refresh Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Rework the root README into a centered, image-rich project landing document while preserving the existing provider facts and contribution rules.

**Architecture:** Keep README as a single GitHub-rendered Markdown document. Use a centered HTML header for visual hierarchy, existing repository images for evidence, and the existing HTML provider table as the structured index.

**Tech Stack:** GitHub-flavored Markdown, inline HTML, existing PNG assets, Bash validation.

---

### Task 1: Recompose the README Header

**Files:** Modify `README.md`.

- [ ] Add a centered hero with project title, Chinese subtitle, shields badges, and quick links.
- [ ] Add a compact table of contents linking to the existing sections.
- [ ] Keep badge labels factual: provider count, documented model count, Markdown format, and repository license.

### Task 2: Add Evidence Image Showcase

**Files:** Modify `README.md`.

- [ ] Add a centered evidence gallery using existing Agnes AI, SenseNova, and B.AI PNG files.
- [ ] Set explicit image widths and meaningful Chinese `alt` text.
- [ ] Keep decorative content separate from factual provider rows.

### Task 3: Tighten Navigation and Section Hierarchy

**Files:** Modify `README.md`.

- [ ] Rename section headings for quick scanning without changing facts.
- [ ] Add a concise “从这里开始” path to the provider table and architecture docs.
- [ ] Preserve the provider table, architecture tree, collection boundary, fixed template, and contribution commands.

### Task 4: Verify and Commit

**Files:** Verify `README.md` and existing image paths.

- [ ] Run `git diff --check`.
- [ ] Run `./scripts/validate-layout.sh`.
- [ ] Verify every README image source exists and every new anchor target is present.
- [ ] Scan for prohibited project names and credential-like values.
- [ ] Commit with `docs: refresh readme visual layout` and push if credentials permit.
