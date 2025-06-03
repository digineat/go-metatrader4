# MT4 Client Library

A lightweight Go client library for interacting with a MetaTrader 4 (MT4) trading server over TCP.

## Example Usage

```go
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

## Requirements

- Go 1.21+
- MetaTrader 4 server with TCP access

## Maintainer

Created and maintained by [Valentin Popov](mailto:valentin@popov.link).

For issues, visit the [GitHub Issues Page](https://github.com/valentineus/go-metatrader4/issues).

## License

This project is licensed under the [MIT License](LICENSE.txt).