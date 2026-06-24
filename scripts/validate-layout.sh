#!/usr/bin/env bash
set -euo pipefail

required_files=(
  "README.md"
  "official-links.md"
  "models.md"
  "api.md"
  "integrations/claude-code.md"
  "integrations/openclaw.md"
  "integrations/opencode.md"
  "integrations/codex.md"
  "integrations/hermes.md"
)

root="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
providers_dir="$root/providers"

if [[ ! -d "$providers_dir" ]]; then
  echo "Missing providers directory: $providers_dir" >&2
  exit 1
fi

failed=0

for provider_dir in "$providers_dir"/*; do
  [[ -d "$provider_dir" ]] || continue
  provider="$(basename "$provider_dir")"

  for file in "${required_files[@]}"; do
    if [[ ! -f "$provider_dir/$file" ]]; then
      echo "Missing $provider/$file" >&2
      failed=1
    fi
  done
done

if [[ "$failed" -ne 0 ]]; then
  exit 1
fi

echo "Provider layout OK"
