package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("cpus:", runtime.NumCPU())

	fmt.Println("Go routines:", runtime.NumGoroutine())

	// this is shared counter
	counter := 0

	// no of go routines to spawn
	gs := 100

	var wg sync.WaitGroup

	wg.Add(gs)
	for i := 0; i < gs; i++ {

		go func() {
			// v is local variable
			v := counter

			//  Yields the processor to allow other goroutines to run. This increases the chance of race conditions by introducing a delay.
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()

		fmt.Println("Go routines:", runtime.NumGoroutine())
	}
	wg.Wait()

	fmt.Println("!!!!!!!!!!!!!!!!!!Go routines:", runtime.NumGoroutine())
	fmt.Println("!!!!!!!!!!!!!!!!!!counter:", counter)

}
