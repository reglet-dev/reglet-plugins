#!/bin/bash
# Run tests for all plugins
# Usage: ./scripts/test-all.sh

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(dirname "$SCRIPT_DIR")"
PLUGINS_DIR="$ROOT_DIR/plugins"

echo "Testing all Reglet plugins..."
echo ""

# Track results
PASSED=0
FAILED=0
TOTAL_TESTS=0

for plugin_dir in "$PLUGINS_DIR"/*/; do
    plugin_name=$(basename "$plugin_dir")
    echo "Testing $plugin_name..."
    
    if output=$(cd "$plugin_dir" && go test -v ./... 2>&1); then
        # Count tests
        test_count=$(echo "$output" | grep -c "^--- PASS" || true)
        TOTAL_TESTS=$((TOTAL_TESTS + test_count))
        echo "  ✓ $plugin_name: $test_count tests passed"
        ((PASSED++))
    else
        echo "  ✗ $plugin_name: tests failed"
        echo "$output" | tail -20
        ((FAILED++))
    fi
    echo ""
done

echo "================================"
echo "Test Summary:"
echo "  Plugins Passed: $PASSED"
echo "  Plugins Failed: $FAILED"
echo "  Total Tests:    $TOTAL_TESTS"
echo "================================"

if [ $FAILED -gt 0 ]; then
    exit 1
fi
