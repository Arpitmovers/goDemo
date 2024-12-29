package main

import "fmt"

func main() {
	// RECIEVE only channel
	c := make(<-chan int, 2)

	// sending to a channel
	c <- 12
	c <- 33

	fmt.Print(<-c)
	fmt.Print(<-c)

}
