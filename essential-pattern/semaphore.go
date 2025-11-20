package essential_pattern

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, sem chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	sem <- struct{}{}
	fmt.Printf("Worker %d started\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d finished\n", id)
	<-sem
}

func ExecSemaphore() {
	const numWorkers = 5
	const maxConcurrent = 2
	sem := make(chan struct{}, maxConcurrent)
	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, sem, &wg)
	}

	wg.Wait()
}
