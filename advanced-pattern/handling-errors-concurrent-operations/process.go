package handling_errors_concurrent_operations

import (
	"context"
	"fmt"
	"time"
)

func process(ctx context.Context, id int, errCh chan<- error) {
	select {
	case <-time.After(2 * time.Second):
		if id%2 == 0 {
			errCh <- fmt.Errorf("error in task %d", id)
		} else {
			fmt.Printf("Task %d completed successfully\n", id)
		}
	case <-ctx.Done():
		fmt.Println("Task Context cancelled", ctx.Err())
	}
}

func ExecProcess() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	errCh := make(chan error, 3)
	for i := range 3 {
		go process(ctx, i, errCh)
	}

	for range 3 {
		select {
		case err := <-errCh:
			if err != nil {
				fmt.Println("Received error:", err.Error())
			}
		case <-ctx.Done():
			fmt.Println("ExecProcess cancelled", ctx.Err())
			return
		}
	}
}
