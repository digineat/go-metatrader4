# MT4 Client Library

This repository contains a small Go package `mt4` providing access to a MetaTrader4 server.

```go
client := mt4.NewClient("127.0.0.1", 443)
ctx := context.Background()
params := map[string]string{
    "MASTER": "master",
    "LOGIN":  "1000",
    "NAME":   "name",
}
res, err := client.Execute(ctx, "WBLOCKLOGINSUSER", params)
```

Use `Execute` to send arbitrary commands supported by your MT4 server.
