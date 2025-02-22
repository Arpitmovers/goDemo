package main

import (
	"fmt"
	"sync"
)

var count int

var increment = func() {
	count++
}

func main() {
	var once sync.Once

	var increments sync.WaitGroup

	increments.Add(50)

	for i := 0; i < 50; i++ {

		go func() {
			defer increments.Done()
			once.Do(increment)
		}()
	}

	increments.Wait()
	fmt.Println(" count is", count)
}
