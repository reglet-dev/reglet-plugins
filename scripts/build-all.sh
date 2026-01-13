#!/bin/bash
# Build all plugins to WASM
# Usage: ./scripts/build-all.sh

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(dirname "$SCRIPT_DIR")"
PLUGINS_DIR="$ROOT_DIR/plugins"

echo "Building all Reglet plugins..."
echo ""

# Track results
BUILT=0
FAILED=0

for plugin_dir in "$PLUGINS_DIR"/*/; do
    plugin_name=$(basename "$plugin_dir")
    echo "Building $plugin_name..."
    
    if (cd "$plugin_dir" && make build); then
        echo "  ✓ $plugin_name built successfully"
        ((BUILT++))
    else
        echo "  ✗ $plugin_name build failed"
        ((FAILED++))
    fi
    echo ""
done

echo "================================"
echo "Build Summary:"
echo "  Successful: $BUILT"
echo "  Failed:     $FAILED"
echo "================================"

if [ $FAILED -gt 0 ]; then
    exit 1
fi
