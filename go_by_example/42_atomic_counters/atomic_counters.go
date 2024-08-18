package main

// the primary mechanism for managing state in Go is communication over channels
// we saw this for example with worker pools
// there are a few other options for managing state though
// here we'll look at using the sync/atomic package for atomic counters accessed by multiple goroutines

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// we'll use an atomic integer type to represent our (always-positive) counter
	var ops atomic.Uint64

	// a WaitGroup will help us wait for all goroutines to finish their work
	var wg sync.WaitGroup

	// we'll start 50 goroutines that each increment the counter exactly 1000 times
	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				// to atomically increment the counter we use Add
				ops.Add(1)
			}

			wg.Done()
		}()
	}

	// wait until all the goroutines are done
	wg.Wait()

	// here no goroutines are writing to 'ops'
	// but using Load it's safe to atomically read a value
	// even while other goroutines are (atomically) updating it
	fmt.Println("ops:", ops.Load())
}

// we expect to get exactly 50,000 operations
// had we used a non-atomic integer and incremented it with ops++
// we'd likely get a different number, changing between runs
// because the goroutines would interfere with each other
// moreover, we'd get data race failures when running the -race flag
