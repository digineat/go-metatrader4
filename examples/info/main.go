package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.popov.link/metatrader4/mt4"
)

func main() {
	client := mt4.NewClient("127.0.0.1", 443,
		mt4.WithDialTimeout(3*time.Second),
		mt4.WithReadTimeout(5*time.Second),
		mt4.WithWriteTimeout(5*time.Second),
	)
	ctx := context.Background()
	// INFO does not require parameters
	resp, err := client.Execute(ctx, "INFO", nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}
