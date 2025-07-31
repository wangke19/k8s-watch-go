# Kubernetes Watch Go

A Go-based tool for monitoring Kubernetes resources and handling watch events, including error recovery for scenarios like `GOAWAY` errors and resource version expiration.

## Features
- Watches Kubernetes pods for `ADDED`, `MODIFIED`, `DELETED`, and `ERROR` events.
- Handles common watch errors (e.g., `GOAWAY`, resource version expiration).
- Configurable timeout for detecting stuck watches.

## Prerequisites
- Go 1.22+
- Kubernetes cluster access (configured via `kubeconfig`).

## Setup
1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd k8s-watch-go
   ```
2. Install dependencies:
   ```bash
   go mod download
   ```

## Usage
Run the tool with default settings:
```bash
GODEBUG=http2debug=2 go run watch-client.go --kubeconfig=/path/to/kubeconfig [--namespace=your-namespace] 2>&1 | tee http2-debug.log
```

### Configuration
- **Timeout**: Adjust the watch timeout in `goaway-test.go` (default: `3m`).
- **Logging**: Modify log output format or level as needed.

## Error Handling
The tool automatically recovers from:
- `GOAWAY` errors (server-initiated connection closure).
- Resource version expiration (stale watches).

## Contributing
Pull requests are welcome. For major changes, open an issue first.

## License
[MIT](LICENSE)