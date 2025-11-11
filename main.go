package main

import (
	"fmt"

	heco "github.com/thebigyovadiaz/concurrency-pattern-go/advance-pattern/handling-errors-concurrent-operations"
	mgc "github.com/thebigyovadiaz/concurrency-pattern-go/advance-pattern/managing-goroutine-context"
	pgl "github.com/thebigyovadiaz/concurrency-pattern-go/advance-pattern/preventing-goroutine-leaks"
)

func main() {
	fmt.Println("Concurrency Pattern Go")

	// Managing goroutine with context
	mgc.ExecProcessRequest()

	fmt.Printf("\n--------------------------------\n\n")

	// Preventing goroutine leaks
	pgl.ExecWorker()

	fmt.Printf("\n--------------------------------\n\n")

	// Handling errors concurrent operations
	heco.ExecProcess()
}
