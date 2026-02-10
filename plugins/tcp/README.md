# TCP Plugin

TCP connection testing and TLS validation.

## Input

| Field | Type | Required | Default | Description |
|-------|------|----------|---------|-------------|
| `host` | string | yes | | Target host (hostname or IP) |
| `port` | int | yes | | Target port |
| `timeout_ms` | int | no | 5000 | Connection timeout in milliseconds |
| `tls` | bool | no | false | Use TLS/SSL connection |
| `expected_tls_version` | string | no | | Minimum TLS version (`TLS 1.0`, `TLS 1.1`, `TLS 1.2`, `TLS 1.3`) |

## Output

| Field | Type | Always | Description |
|-------|------|--------|-------------|
| `host` | string | yes | Connected host |
| `port` | int | yes | Connected port |
| `connected` | bool | yes | Whether connection succeeded |
| `tls_version` | string | no | Negotiated TLS version |
| `tls_cipher_suite` | string | no | Negotiated cipher suite |
| `tls_expiry` | string | no | Certificate expiry date |
| `response_time_ms` | int | yes | Connection time in milliseconds |
| `error` | string | no | Error message if connection failed |
