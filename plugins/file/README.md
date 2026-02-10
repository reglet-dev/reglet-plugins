# File Plugin

File system checks and validation.

## Input

| Field | Type | Required | Default | Description |
|-------|------|----------|---------|-------------|
| `path` | string | yes | | File or directory path |
| `operation` | string | yes | | Check to perform (`exists`, `permissions`, `checksum`, `content`) |
| `permissions` | string | no | | Expected permissions in octal (e.g., `0644`) |
| `checksum` | string | no | | Expected checksum to compare against |
| `algorithm` | string | no | `sha256` | Checksum algorithm (`md5`, `sha256`, `sha512`) |
| `contains` | string | no | | String to search for in file |

## Output

| Field | Type | Always | Description |
|-------|------|--------|-------------|
| `path` | string | yes | Checked path |
| `exists` | bool | yes | Whether path exists |
| `is_dir` | bool | yes | Whether path is a directory |
| `size` | int | no | File size in bytes |
| `permissions` | string | no | File permissions (octal) |
| `mode` | string | no | File mode (octal) |
| `mod_time` | string | no | Last modification time |
| `uid` | int | yes | File owner UID |
| `gid` | int | yes | File group GID |
| `readable` | bool | yes | Whether file is readable |
| `is_symlink` | bool | yes | Whether path is a symlink |
| `checksum` | string | no | File checksum (with `checksum` operation) |
| `contains` | bool | no | Whether file contains search string (with `content` operation) |
