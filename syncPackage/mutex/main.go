package main

import (
	"fmt"
	"sync"
)

var count int
var lock sync.Mutex

// . If the lock is already in use, the calling goroutine blocks until the mutex is available.
var increment = func() {
	lock.Lock()
	defer lock.Unlock()
	count++
	fmt.Println("count is ", count)
}

var decrement = func() {
	lock.Lock()
	defer lock.Unlock()
	count--
	fmt.Println("decrement count is ", count)
}

// order of go routines can be anything
func main() {
	var arithmetic sync.WaitGroup
	arithmetic.Add(10)

	for i := 0; i < 5; i++ {
		go func() {
			defer arithmetic.Done()
			increment()
		}()
	}

	for i := 0; i < 5; i++ {
		go func() {
			defer arithmetic.Done()
			decrement()
		}()
	}

	arithmetic.Wait()
	fmt.Println("arithmetic complete")
}
