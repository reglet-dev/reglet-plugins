# Reglet Plugins

Official plugin repository for [Reglet](https://github.com/reglet-dev/reglet), the compliance-as-code framework.

## Plugins

| Plugin | Version | Description |
|--------|---------|-------------|
| [command](plugins/command) | 1.0.0 | Execute commands and validate output |
| [dns](plugins/dns) | 1.0.0 | DNS resolution and record validation |
| [file](plugins/file) | 1.1.0 | File existence, content, and hash checks |
| [http](plugins/http) | 1.0.0 | HTTP/HTTPS request checking and validation |
| [smtp](plugins/smtp) | 1.0.0 | SMTP connection testing and server validation |
| [tcp](plugins/tcp) | 1.0.0 | TCP connection testing and TLS validation |

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
├── plugins/
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
├── go.mod           # Go module definition
├── go.sum           # Dependencies
├── main.go          # WASM entry point (build tag: wasip1)
├── plugin.go        # Plugin implementation
├── plugin_test.go   # Tests (build tag: !wasip1)
├── Makefile         # Build targets
└── README.md        # Plugin documentation
```

## Capabilities

Plugins declare capabilities they require. The Reglet host grants or denies these based on user configuration.

| Plugin | Capability | Pattern |
|--------|-----------|---------|
| command | exec | `**` |
| dns | network | `outbound:53` |
| file | fs | `read:**` |
| http | network | `outbound:80,443` |
| smtp | network | `outbound:25,465,587` |
| tcp | network | `outbound:*` |

## SDK

All plugins use the [Reglet SDK](https://github.com/reglet-dev/reglet-sdk):

```go
import regletsdk "github.com/reglet-dev/reglet-sdk/go"
```

## License

Apache 2.0 - See [LICENSE](LICENSE)
