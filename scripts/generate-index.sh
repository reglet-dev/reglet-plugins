#!/usr/bin/env bash
set -euo pipefail

# Collect plugin metadata from each plugin.json,
# merged with latest git tag version.

script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
repo_root="$(dirname "$script_dir")"

registry="ghcr.io/reglet-dev/plugins"
repo_url="https://github.com/reglet-dev/reglet-plugins"

plugins_json="[]"

# ensure we affect the correct path.

for plugin_json in "$repo_root"/plugins/*/plugin.json; do
    # check if file exists to handle empty case
    [ -e "$plugin_json" ] || continue
    
    plugin_dir="$(dirname "$plugin_json")"
    plugin_name="$(jq -r '.name' "$plugin_json")"

    # Find latest tag for this plugin (e.g. dns/v1.2.0 -> 1.2.0)
    latest=$(git tag -l "${plugin_name}/v*" --sort=-v:refname | head -1 | sed "s|${plugin_name}/v||" || echo "unreleased")
    if [ -z "$latest" ]; then
        latest="unreleased"
    fi

    entry=$(jq --arg v "$latest" '. + {latest: $v}' "$plugin_json")
    plugins_json=$(echo "$plugins_json" | jq --argjson e "$entry" '. + [$e]')
done

jq -n \
    --arg repo "$repo_url" \
    --arg registry "$registry" \
    --arg updated "$(date -u +%Y-%m-%dT%H:%M:%SZ)" \
    --argjson plugins "$plugins_json" \
    '{repository: $repo, registry: $registry, updated: $updated, plugins: $plugins}' \
    > "$repo_root/index.json"

echo "Generated index.json with $(echo "$plugins_json" | jq length) plugins"
