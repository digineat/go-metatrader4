# MT4 Client Library

[![Go Reference](https://pkg.go.dev/badge/github.com/digineat/go-metatrader4.svg)](https://pkg.go.dev/github.com/digineat/go-metatrader4)
[![Go Report Card](https://goreportcard.com/badge/github.com/digineat/go-metatrader4)](https://goreportcard.com/report/github.com/digineat/go-metatrader4)

A lightweight Go client library for interacting with a MetaTrader 4 (MT4) trading server over TCP.

## Example Usage

```go
import "github.com/digineat/go-metatrader4/mt4"

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

- `WithDialTimeout(d time.Duration)`: Sets the timeout for establishing a TCP connection. Default: 5s.
- `WithReadTimeout(d time.Duration)`: Sets the maximum time to wait for a server response. Default: 5s.
- `WithWriteTimeout(d time.Duration)`: Sets the maximum time to complete sending a request. Default: 5s.
- `WithAutoClose(enabled bool)`: If `true`, closes the connection after each `Execute` (default). Use `false` to reuse the session manually via `client.Close()`.

## Requirements

- Go 1.24 or later
- MetaTrader 4 server with TCP access

## Project Info

- Vanity import path: `github.com/digineat/go-metatrader4/mt4`
- Issues and contributions: [GitHub](https://github.com/digineat/go-metatrader4/issues)

Maintained by [Valentin Popov](mailto:valentin.p@digineat.com).

## License

This project is licensed under the [MIT License](LICENSE.txt).
