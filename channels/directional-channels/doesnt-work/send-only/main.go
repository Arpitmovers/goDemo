package main

import "fmt"

func main() {
	// send only channel
	c := make(chan<- int, 2)

	c <- 12
	c <- 33

	// invalid operation: cannot receive from send-only channel c (variable of type chan<- int)
	fmt.Print(<-c)
	fmt.Print(<-c)

}
