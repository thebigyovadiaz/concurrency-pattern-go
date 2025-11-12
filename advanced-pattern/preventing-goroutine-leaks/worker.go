package preventing_goroutine_leaks

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, ch chan<- string) {
	select {
	case <-time.After(1 * time.Second):
		ch <- "Work completed!"
	case <-ctx.Done():
		fmt.Println("Worker context cancelled: ", ctx.Err())
	}
}

func ExecWorker() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	ch := make(chan string)
	go worker(ctx, ch)

	select {
	case msg := <-ch:
		fmt.Println(msg)
	case <-ctx.Done():
		fmt.Println("ExecWorker cancelled: ", ctx.Err())
	}
}
