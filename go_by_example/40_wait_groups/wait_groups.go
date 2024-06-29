package main

import (
	"fmt"
	"sync"
	"time"
)

// to wait for multiple goroutines to finish, we can use a wait group

func worker(id int) {
	fmt.Printf("worker %d starting\n", id)

	// sleep to simulate and expressive task
	time.Sleep(time.Second)
	fmt.Printf("worker %d done\n", id)
}

func main() {
	// this WaitGroup is used to wait for all the goroutines launched her to finish
	// note: if a WaitGroup is explicitly passed into functions, it should be done by pointer
	var wg sync.WaitGroup

	// launch several goroutines and increment the WaitGroup counter for each
	for i := 1; i <= 5; i++ {
		wg.Add(1)

		// wrap the worker call in a closure that makes sure to tell the WaitGroup that this worker is done
		// this way the worker itself does not have to be aware of the concurrency primitives involved in tis execution
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	// block until the WaitGroup counter goes back to 0; all the workers notified they're done
	wg.Wait()

	// note that this approach has no straightforward way to propagate errors from workers
	// for more advanced use cases, consider the errgroup package
	// https://pkg.go.dev/golang.org/x/sync/errgroup
}

// the order of workers starting up and finishing is likely to be different for each invocation
