# HTTP Plugin

HTTP/HTTPS request checking and validation.

## Input

| Field | Type | Required | Default | Description |
|-------|------|----------|---------|-------------|
| `url` | string | yes | | URL to request (must include protocol) |
| `method` | string | no | `GET` | HTTP method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `PATCH`) |
| `headers` | map[string]string | no | | Request headers |
| `body` | string | no | | Request body |
| `expected_status` | int | no | 200 | Expected HTTP status code |
| `expected_body_contains` | string | no | | String that response body should contain |
| `body_preview_length` | int | no | 200 | Response body chars to include (0=hash only, -1=full) |
| `follow_redirects` | bool | no | true | Follow HTTP redirects |
| `timeout_seconds` | int | no | 30 | Request timeout in seconds (max 300) |

## Output

| Field | Type | Always | Description |
|-------|------|--------|-------------|
| `url` | string | yes | Requested URL |
| `method` | string | yes | HTTP method used |
| `status_code` | int | yes | HTTP response status code |
| `status` | string | yes | HTTP status text |
| `protocol` | string | yes | HTTP protocol version |
| `headers` | map[string][]string | yes | Response headers |
| `body` | string | no | Response body (truncated if large) |
| `body_size` | int | yes | Response body size in bytes |
| `tls_version` | string | no | TLS version (HTTPS only) |
| `tls_expiry` | string | no | TLS certificate expiry date |
| `tls_days_until_expiry` | int | no | Days until certificate expires |
| `response_time_ms` | int | yes | Response time in milliseconds |
