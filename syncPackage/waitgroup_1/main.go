package main

import (
	"fmt"
	"sync"
)

var count int

var hello = func(value int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("value is ", value)
}

func main() {

	var increments sync.WaitGroup

	increments.Add(5)

	for i := 0; i < 5; i++ {
		go hello(i, &increments)

	}
	increments.Wait()

}
