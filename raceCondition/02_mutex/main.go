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

	var mu sync.Mutex

	for i := 0; i < gs; i++ {

		go func() {
			// v is local variable
			fmt.Println("cpus:", runtime.NumCPU())
			// when some is reading /writing , no other process can read /write it
			mu.Lock()
			v := counter

			//  Yields the processor to allow other goroutines to run. This increases the chance of race conditions by introducing a delay.
			runtime.Gosched()
			v++
			counter = v

			mu.Unlock()
			wg.Done()
		}()

		fmt.Println("Go routines:", runtime.NumGoroutine())
	}
	wg.Wait()

	fmt.Println("!!!!!!!!!!!!!!!!!!Go routines:", runtime.NumGoroutine())
	fmt.Println("!!!!!!!!!!!!!!!!!!counter:", counter)

}
