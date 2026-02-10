# SMTP Plugin

SMTP connection testing and server validation.

## Input

| Field | Type | Required | Default | Description |
|-------|------|----------|---------|-------------|
| `host` | string | yes | | SMTP server host (hostname or IP) |
| `port` | int | no | 25 | SMTP server port |
| `timeout_ms` | int | no | 5000 | Connection timeout in milliseconds |
| `use_tls` | bool | no | false | Use direct TLS/SSL connection (port 465) |
| `use_starttls` | bool | no | false | Upgrade via STARTTLS (port 587) |

### Common Port Configurations

| Port | Protocol | use_tls | use_starttls |
|------|----------|---------|--------------|
| 25 | Plain SMTP | false | false |
| 465 | SMTPS (implicit TLS) | true | false |
| 587 | Submission (STARTTLS) | false | true |

## Output

| Field | Type | Always | Description |
|-------|------|--------|-------------|
| `host` | string | yes | Connected host |
| `port` | int | yes | Connected port |
| `banner` | string | yes | SMTP banner message |
| `extensions` | []string | no | Supported SMTP extensions |
| `tls_version` | string | no | TLS version (if encrypted) |
| `supports_auth` | bool | yes | Whether AUTH is supported |
