package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	// solution  as go routine
	go func() {
		c <- 42
	}()

	fmt.Println(<-c)

	// solition as a buffered channel
	d := make(chan int, 1)
	d <- 22
	fmt.Println(<-d)

}
