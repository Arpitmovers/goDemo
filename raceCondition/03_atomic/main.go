package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	// fmt.Println("cpus:", runtime.NumCPU())

	// fmt.Println("INIT Go routines:", runtime.NumGoroutine())

	// this is shared counter
	var counter int64
	// no of go routines to spawn
	gs := 100

	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {

		go func() {

			atomic.AddInt64(&counter, 1)
			fmt.Println("Go routines:", runtime.NumGoroutine())

			fmt.Println("counter is ", atomic.LoadInt64(&counter)) // read from
			fmt.Println("cpus:", runtime.NumCPU())
			wg.Done()
		}()

	}

	wg.Wait()
	// fmt.Println("!!!!!!!!!!!!!!!!!!Go routines:", runtime.NumGoroutine())
	// fmt.Println("!!!!!!!!!!!!!!!!!!counter:", counter)

}
