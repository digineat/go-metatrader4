# MT4 Client Library

A small Go package providing an interface to a MetaTrader4 server.

```go
client := mt4.NewClient("127.0.0.1", 443,
    mt4.WithDialTimeout(3*time.Second),
    mt4.WithAutoClose(true),
)
ctx := context.Background()
params := map[string]string{
    "MASTER": "master",
    "LOGIN":  "1000",
    "NAME":   "name",
}
res, err := client.Execute(ctx, "WBLOCKLOGINSUSER", params)
```

`Execute` sends a command with parameters, encoding them using the MT4 protocol.
Use `WithAutoClose(false)` and `client.Close()` for session reuse.
