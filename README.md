# MT4 Client Library [![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE.txt)

A small Go package providing an interface to a MetaTrader4 server.

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

`Execute` sends a command with parameters, encoding them using the MT4 protocol.
Use `WithAutoClose(false)` and `client.Close()` for session reuse.

---

**License**

This project is licensed under the [MIT License](LICENSE.txt).

**Contact**

Valentin Popov <valentin@popov.link>

For issues please visit <https://github.com/valentineus/go-metatrader4/issues>.
