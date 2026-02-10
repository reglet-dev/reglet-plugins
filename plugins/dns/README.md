# DNS Plugin

DNS resolution and record lookup.

## Input

| Field | Type | Required | Default | Description |
|-------|------|----------|---------|-------------|
| `hostname` | string | yes | | Domain name to resolve |
| `record_type` | string | no | `A` | DNS record type (`A`, `AAAA`, `CNAME`, `MX`, `TXT`, `NS`) |
| `nameserver` | string | no | | Custom nameserver (e.g., `8.8.8.8:53`) |

## Output

| Field | Type | Always | Description |
|-------|------|--------|-------------|
| `hostname` | string | yes | Queried hostname |
| `record_type` | string | yes | Record type queried |
| `records` | []string | yes | Resolved DNS records |
| `record_count` | int | yes | Number of resolved records |
| `ttl` | int | no | Record TTL in seconds |
| `nameserver` | string | no | Nameserver used (if custom) |
