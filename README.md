# MT4 Client Library

This repository contains a small Go package `mt4` providing access to a MetaTrader4 server.

```go
client := mt4.NewClient("127.0.0.1", 443)
ctx := context.Background()
res, err := client.BlockAccount(ctx, mt4.BlockAccountOptions{
    Login:  1000,
    Master: "master",
    Name:   "name",
})
```

The library exposes a low level `Execute` method for sending custom commands.
