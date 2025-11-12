package managing_goroutine_context

import (
	"context"
	"fmt"
	"time"
)

func processRequest(ctx context.Context, id int) {
	select {
	case <-time.After(2 * time.Second):
		fmt.Printf("Processed request %d\n", id)
	case <-ctx.Done():
		fmt.Printf("Request %d cancelled: %v\n", id, ctx.Err())
	}
}

func ExecProcessRequest() {
	// Context timeout in 1 second
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	for v := range 3 {
		go processRequest(ctx, v)
	}

	// Main sleep for 3 second
	time.Sleep(3 * time.Second)
}
