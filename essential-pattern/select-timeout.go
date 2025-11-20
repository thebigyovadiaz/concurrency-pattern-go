package essential_pattern

import (
	"fmt"
	"time"
)

func ExecSelectTimeout() {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		ch <- "result"
	}()

	select {
	case res := <-ch:
		fmt.Println("Received:", res)
	case <-time.After(time.Second * 3):
		fmt.Println("Timeout")
	}
}
