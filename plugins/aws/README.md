# AWS Plugin

AWS infrastructure inspection and compliance checks.

Credentials are read from the environment: `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`, and optionally `AWS_SESSION_TOKEN`.

## Input

| Field | Type | Required | Default | Description |
|-------|------|----------|---------|-------------|
| `service` | string | yes | | AWS service (`iam`, `ec2`) |
| `operation` | string | yes | | Service-specific operation (see below) |
| `region` | string | no | `AWS_REGION` env | AWS region |
| `timeout_seconds` | int | no | 30 | Request timeout in seconds (max 300) |
| `filters` | map[string][]string | no | | AWS API filters (service-specific) |

## Operations

### EC2

#### `describe_security_groups`

Find security groups with open SSH/RDP to the internet.

| Field | Type | Always | Description |
|-------|------|--------|-------------|
| `region` | string | yes | AWS region queried |
| `total_groups` | int | yes | Total security groups |
| `security_groups` | []object | yes | All security group details |
| `open_ssh_groups` | []object | yes | Groups with open SSH to internet |
| `open_rdp_groups` | []object | yes | Groups with open RDP to internet |

#### `describe_instances_metadata`

Check IMDSv2 enforcement on EC2 instances.

| Field | Type | Always | Description |
|-------|------|--------|-------------|
| `region` | string | yes | AWS region queried |
| `total_instances` | int | yes | Total running instances |
| `instances` | []object | yes | All instance metadata details |
| `non_compliant_instances` | []object | yes | Instances not enforcing IMDSv2 |

### IAM

#### `get_account_summary`

Check root MFA and account summary metrics.

| Field | Type | Always | Description |
|-------|------|--------|-------------|
| `root_mfa_enabled` | bool | yes | Whether root account has MFA |
| `users` | int | yes | Number of IAM users |
| `groups` | int | yes | Number of IAM groups |
| `roles` | int | yes | Number of IAM roles |
| `policies` | int | yes | Number of customer-managed policies |
| `mfa_devices` | int | yes | Number of MFA devices |
| `mfa_devices_in_use` | int | yes | Number of MFA devices in use |
| `access_keys_per_user_quota` | int | yes | Access keys per user quota |

#### `get_account_password_policy`

Check IAM password policy configuration.

| Field | Type | Always | Description |
|-------|------|--------|-------------|
| `policy_exists` | bool | yes | Whether a password policy exists |
| `password_policy` | object | no | Password policy configuration |
| `password_policy.minimum_length` | int | | Minimum password length |
| `password_policy.require_symbols` | bool | | Whether symbols are required |
| `password_policy.require_numbers` | bool | | Whether numbers are required |
| `password_policy.require_uppercase` | bool | | Whether uppercase required |
| `password_policy.require_lowercase` | bool | | Whether lowercase required |
| `password_policy.max_age_days` | int | | Maximum password age in days |
| `password_policy.password_reuse_prevention` | int | | Previous passwords to prevent reuse |

#### `list_access_keys_with_usage`

Find unused access keys.

Additional input: `threshold_days` (int, default: 90) — days of inactivity to flag.

| Field | Type | Always | Description |
|-------|------|--------|-------------|
| `total_users` | int | yes | Total IAM users |
| `total_access_keys` | int | yes | Total access keys |
| `access_keys` | []object | yes | All access key details |
| `unused_keys_over_threshold` | []object | yes | Keys unused beyond threshold |
