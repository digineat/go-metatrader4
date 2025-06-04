# Example: INFO Command

This example demonstrates how to use the [`go-metatrader4`](https://github.com/valentineus/go-metatrader4) library to send the `INFO` command to a MetaTrader 4 (MT4) server and retrieve server information.

The `INFO` command requests basic server details such as build version and company name.

## Usage

To run this example:

```bash
go run main.go
```

Make sure you are connected to an MT4 server that accepts TCP connections on the configured host and port.

## Code Overview

```go
client := mt4.NewClient("127.0.0.1", 443,
    mt4.WithDialTimeout(3*time.Second),
    mt4.WithReadTimeout(5*time.Second),
    mt4.WithWriteTimeout(5*time.Second),
)
ctx := context.Background()
resp, err := client.Execute(ctx, "INFO", nil)

```

This code creates an MT4 client, sends the INFO command without parameters, and prints the response to stdout.

## Expected Response Format

The response typically looks like this:

```text
MetaTrader 4 Server 4.00 build 1380
Some Broker Company Name
```

Where:

- `build 1380` — current server build number
- `Some Broker Company Name` — name of the White Label owner of the server

## Requirements

- Go 1.24 or later
- Access to a running MetaTrader 4 server

## License

This example is provided under the MIT License. See the [main project license](../../LICENSE.txt) for details.