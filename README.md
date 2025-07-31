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

## Expected Results
Monitor 1hour, found some GOAWAY related logs,
```bash
$ grep GOAWAY http2-debug.log 
2025/07/31 12:42:23 Watch channel closed. Possibly due to GOAWAY or timeout. Reconnecting...
2025/07/31 12:45:08 http2: Framer 0xc0000f0000: read GOAWAY len=8 LastStreamID=1 ErrCode=NO_ERROR Debug=""
2025/07/31 12:45:08 http2: Transport received GOAWAY len=8 LastStreamID=1 ErrCode=NO_ERROR Debug=""
2025/07/31 12:45:10 Watch channel closed. Possibly due to GOAWAY or timeout. Reconnecting...
2025/07/31 12:47:23 Watch channel closed. Possibly due to GOAWAY or timeout. Reconnecting...
2025/07/31 12:50:32 Watch channel closed. Possibly due to GOAWAY or timeout. Reconnecting...
2025/07/31 12:52:42 Watch channel closed. Possibly due to GOAWAY or timeout. Reconnecting...
2025/07/31 12:54:25 Watch channel closed. Possibly due to GOAWAY or timeout. Reconnecting...
2025/07/31 12:56:26 Watch channel closed. Possibly due to GOAWAY or timeout. Reconnecting...
2025/07/31 12:58:25 Watch channel closed. Possibly due to GOAWAY or timeout. Reconnecting...
2025/07/31 13:00:07 Watch channel closed. Possibly due to GOAWAY or timeout. Reconnecting...
2025/07/31 13:08:10 Watch channel closed. Possibly due to GOAWAY or timeout. Reconnecting...
2025/07/31 13:11:13 Watch channel closed. Possibly due to GOAWAY or timeout. Reconnecting...
2025/07/31 13:13:48 http2: Framer 0xc000346000: read GOAWAY len=8 LastStreamID=1 ErrCode=NO_ERROR Debug=""
2025/07/31 13:13:48 http2: Transport received GOAWAY len=8 LastStreamID=1 ErrCode=NO_ERROR Debug=""
2025/07/31 13:20:53 Watch channel closed. Possibly due to GOAWAY or timeout. Reconnecting...
2025/07/31 13:29:41 Watch channel closed. Possibly due to GOAWAY or timeout. Reconnecting...
2025/07/31 13:38:31 Watch channel closed. Possibly due to GOAWAY or timeout. Reconnecting...
2025/07/31 13:41:20 Watch channel closed. Possibly due to GOAWAY or timeout. Reconnecting...
2025/07/31 13:42:44 Watch channel closed. Possibly due to GOAWAY or timeout. Reconnecting...
2025/07/31 13:44:47 Watch channel closed. Possibly due to GOAWAY or timeout. Reconnecting...
2025/07/31 13:48:45 Watch channel closed. Possibly due to GOAWAY or timeout. Reconnecting...
```

## Contributing
Pull requests are welcome. For major changes, open an issue first.

## License
[MIT](LICENSE)