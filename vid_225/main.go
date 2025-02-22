//  pull the values off the channel using a for range loop
// https://go.dev/play/p/sfyu4Is3FG


package main

import (
	"fmt"
)

func main() {
	c := gen()
	receive(c)
	fmt.Println("about to exit")
}

// returns recive
func gen() chan int {
	c := make(chan int) // bi directional channels

	go func() {
		for i := 0; i < 12; i++ {
			c <- i
		}
		close(c)
	}()

	return c // returns recive channel
}

func receive(c chan int) {

	for v := range c {
		fmt.Print(v, "\n")
	}
	for v := range c {
		fmt.Print(v, "\n")
	}
}
