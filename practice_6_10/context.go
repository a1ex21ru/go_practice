package main

import (
	"context"
	"fmt"
	"time"
)

const (
	userId = 0
)

func contextFunc() {

	ctx := context.WithValue(context.Background(), userId, 1)

	ctx, cancel := context.WithCancel(ctx)

	go func() {
		fmt.Scanln()
		cancel()
	}()
	fmt.Println(processLongTask(ctx))
}

func processLongTask(ctx context.Context) string {

	id := ctx.Value(userId)
	select {
	case <-time.After(2 * time.Second):
		return fmt.Sprintln("done processing ", id)
	case <-ctx.Done():
		return fmt.Sprintln("request cancelled")
	}

}
