package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, finish := context.WithTimeout(context.Background(), time.Second*60)
	// ctx, finish := context.WithDeadline(context.Background(), time.Now().Add(time.Second*60))
	// ctx, finish := context.WithCancel(context.Background())

	// ctx - is a context instance
	// finish - is a finish callback for ctx, shutdown of all goroutines of this ctx
	finish()

	// infinity loop
LOOP:
	for {
		// listen for done event from ctx done channel
		select {
		case <-ctx.Done(): // read-only channel
			fmt.Println(ctx.Err()) // context canceled
			break LOOP
		default:
			fmt.Println("default case")
			break LOOP
		}
	}

	/*
		configKey := "propertyName"
		config := map[string]string{}
		ctx := context.WithValue(ctx, configKey, &config)
	*/
}
