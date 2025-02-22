// write a program that
// ○ launches 10 goroutines
// ■ each goroutine adds 10 numbers to a channel
// ○ pull the numbers off the channel and print them

package main

import (
	"fmt"
	"runtime"
)

func main() {
	x := 10 // no of go routines to launch
	y := 2  // each go routines adds 10 no to channel
	c := gen(x, y)

	for i := 0; i < x*y; i++ {
		fmt.Println(i, <-c)
	}
	fmt.Println("ROUTINES", runtime.NumGoroutine())
}

func gen(x, y int) <-chan int {
	c := make(chan int)

	for i := 0; i < x; i++ {

		go func() {
			for j := 0; j < y; j++ {
				c <- j
			}
		}()

		fmt.Println("no of goroutines that exist right now", runtime.NumGoroutine())
	}

	return c
}
