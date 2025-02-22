package main

import (
	"fmt"
	"runtime"
	"sync"
)

func memConsumed() uint64 {

	runtime.GC()
	var s runtime.MemStats
	runtime.ReadMemStats(&s)
	return s.Sys
}

var wg sync.WaitGroup
var c <-chan interface{}

var noop = func() {
	wg.Done()
	<-c

}

func main() {
	var noOfRoutines = 10000
	before := memConsumed()
	wg.Add(noOfRoutines)

	for i := noOfRoutines; i > 0; i-- {
		go noop()
	}
	wg.Wait()

	afterMem := memConsumed()
	fmt.Println("before ", before)
	fmt.Println("after", afterMem)
}
