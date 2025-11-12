package essential_pattern

import (
	"fmt"
	"sync"
)

const (
	PRODUCERS  int = 5
	INCREMENT  int = 2
	CAPCHANNEL int = 10
)

func producer(id int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range PRODUCERS {
		ch <- i
		fmt.Printf("Producer %d created %d\n", id, i)
	}
}

func consumer(id int, in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := range in {
		out <- i * INCREMENT
		fmt.Printf("Consumer %d created %d\n", id, i)
	}
}

func ExecFanInFanOut() {
	var wg sync.WaitGroup
	numProducers := INCREMENT
	numConsumers := INCREMENT
	input, output := make(chan int, CAPCHANNEL), make(chan int, CAPCHANNEL)

	// Fan-Out: start producers
	for i := range numProducers {
		wg.Add(1)
		go producer(i, input, &wg)
	}
	wg.Wait()
	close(input)

	// Fan-Out: start consumers
	for i := range numConsumers {
		wg.Add(1)
		go consumer(i, input, output, &wg)
	}
	wg.Wait()
	close(output)

	// Fan-In: aggregate results
	for result := range output {
		fmt.Println("Result:", result)
	}
}
