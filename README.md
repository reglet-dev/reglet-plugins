# Reglet Plugins

Official plugin repository for [Reglet](https://github.com/reglet-dev/reglet), the compliance-as-code framework.

## Plugins

| Plugin | Description |
|--------|-------------|
| [aws](plugins/aws) | AWS infrastructure inspection and compliance |
| [command](plugins/command) | Execute commands and validate output |
| [dns](plugins/dns) | DNS resolution and record validation |
| [file](plugins/file) | File system checks and validation |
| [http](plugins/http) | HTTP/HTTPS request checking and validation |
| [smtp](plugins/smtp) | SMTP connection testing and server validation |
| [tcp](plugins/tcp) | TCP connection testing and TLS validation |

## Quick Start

### Building All Plugins

```bash
make build-all
```

### Testing All Plugins

```bash
make test-all
```

### Building a Single Plugin

```bash
make -C plugins/command build
```

## Development

Each plugin is an independent Go module that compiles to WASM. See [Plugin Development Guide](docs/plugin-development.md) for details.

### Prerequisites

- Go 1.25+
- WASM support (`GOOS=wasip1 GOARCH=wasm`)

### Project Structure

```
reglet-plugins/
├── .github/workflows/ # CI/CD (publish-plugin.yml)
├── plugins/
│   ├── aws/         # AWS infrastructure plugin
│   ├── command/     # Command execution plugin
│   ├── dns/         # DNS resolution plugin
│   ├── file/        # File system plugin
│   ├── http/        # HTTP request plugin
│   ├── smtp/        # SMTP connection plugin
│   └── tcp/         # TCP connection plugin
├── scripts/         # Build and test automation
├── docs/            # Documentation
└── Makefile         # Top-level build commands
```

### Plugin Structure

Each plugin follows this structure:

```
plugin-name/
├── core/            # Plugin definition and config types
│   ├── plugin.go    # PluginDef (name, capabilities, etc.)
│   └── config.go    # Config struct with jsonschema tags
├── services/        # Operation handlers
├── plugin.go        # WASM entry point
├── plugin.json      # Static metadata for CI (name, description, capabilities)
├── go.mod           # Go module definition
├── Makefile         # Build targets
└── README.md        # Plugin documentation
```

## Capabilities

Plugins declare capabilities they require. The Reglet host grants or denies these based on user configuration.

| Plugin | Capability | Pattern |
|--------|-----------|---------|
| aws | network | `outbound:*.amazonaws.com:443` |
| command | exec | `**` |
| dns | network | `outbound:53` |
| file | fs | `read:**` |
| http | network | `outbound:80,443` |
| smtp | network | `outbound:25,465,587` |
| tcp | network | `outbound:*` |

## Publishing

Plugins are published as OCI artifacts to `ghcr.io/reglet-dev/plugins/<name>:<version>`. Push a tag to trigger the CI workflow:

```bash
git tag command/v1.2.0
git push origin command/v1.2.0
```

This builds the WASM binary and pushes it via ORAS with the correct media types for the Reglet host to pull.

## SDK

All plugins use the [Reglet Plugin SDK](https://github.com/reglet-dev/reglet-plugin-sdk):

```go
import "github.com/reglet-dev/reglet-plugin-sdk"
```

## License

Apache 2.0 - See [LICENSE](LICENSE)
