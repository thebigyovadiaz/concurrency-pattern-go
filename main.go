package main

import (
	"fmt"

	ep "github.com/thebigyovadiaz/concurrency-pattern-go/essential-pattern"
)

func main() {
	fmt.Println("Concurrency Pattern Go")

	/* fmt.Printf("\nAdvance Pattern\n")
	// Managing goroutine with context
	mgc.ExecProcessRequest()
	fmt.Printf("\n--------------------------------\n\n")

	// Preventing goroutine leaks
	pgl.ExecWorker()
	fmt.Printf("\n--------------------------------\n\n")

	// Handling errors concurrent operations
	heco.ExecProcess()
	fmt.Printf("\n--------------------------------\n\n") */

	fmt.Printf("\nEssential Pattern\n")

	// Worker Pool Pattern
	ep.ExecWorkerPool()
	fmt.Printf("\n--------------------------------\n\n")

	// Fan-In & Fan-Out Pattern
	ep.ExecFanInFanOut()
	fmt.Printf("\n--------------------------------\n\n")

	// Pipeline Pattern
	ep.ExecPipeline()
}
