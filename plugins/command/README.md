# Command Plugin

Execute commands and validate output.

## Input

| Field | Type | Required | Default | Description |
|-------|------|----------|---------|-------------|
| `run` | string | one of `run`/`command` | | Shell command to execute via `/bin/sh -c` |
| `command` | string | one of `run`/`command` | | Executable path for direct execution |
| `args` | []string | no | | Arguments for direct execution |
| `dir` | string | no | | Working directory |
| `env` | []string | no | | Environment variables (`KEY=VALUE`) |
| `timeout_seconds` | int | no | 60 | Execution timeout in seconds |
| `expected_exit` | int | no | 0 | Expected exit code |
| `expected_output` | string | no | | Expected string in stdout |

## Output

| Field | Type | Always | Description |
|-------|------|--------|-------------|
| `command` | string | yes | Executed command |
| `exit_code` | int | yes | Process exit code |
| `stdout` | string | yes | Standard output |
| `stderr` | string | yes | Standard error |
| `duration_ms` | int | yes | Execution time in milliseconds |
