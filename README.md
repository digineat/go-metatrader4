# MT4 Client Library

A lightweight Go client library for interacting with a MetaTrader 4 (MT4) trading server over TCP.

## Example Usage

```go
import "go.popov.link/metatrader4/mt4"

client := mt4.NewClient("127.0.0.1", 443,
    mt4.WithDialTimeout(3*time.Second),
    mt4.WithAutoClose(true),
)
ctx := context.Background()
params := map[string]string{
    "login":    "55555",
    "password": "_some_password_",
}
res, err := client.Execute(ctx, "WWAPUSER", params)
```

The `Execute` method sends a raw MT4 command. Parameters are encoded using base64 and Windows-1251.
Use `WithAutoClose(false)` if you want to reuse the connection manually via `client.Close()`.

## Options

- `WithDialTimeout(d time.Duration)` sets how long the client waits when establishing a TCP connection. Default is five seconds.
- `WithReadTimeout(d time.Duration)` sets the maximum time allowed to read a server response. Default is five seconds.
- `WithWriteTimeout(d time.Duration)` sets the maximum time allowed to send a request to the server. Default is five seconds.
- `WithAutoClose(enabled bool)` closes the connection after every `Execute` call when set to `true` (default). Set to `false` to reuse the connection and close it manually.

## Requirements

- Go 1.24 or later
- MetaTrader 4 server with TCP access

## Maintainer & Project Info

- Vanity import path: `go.popov.link/metatrader4`
- Source mirror (read-only): [code.popov.link](https://code.popov.link/valentineus/go-metatrader4)
- Issues and contributions: [GitHub](https://github.com/valentineus/go-metatrader4/issues)

Maintained by [Valentin Popov](mailto:valentin@popov.link).

## License

This project is licensed under the [MIT License](LICENSE.txt).