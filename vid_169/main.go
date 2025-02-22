// write a program that
// ○ puts 100 numbers to a channel
// ○ pull the numbers off the channel and print them

package main

import (
	"fmt"
)

func main() {
	q := make(chan int)

	gen(q)
	print(q)
	fmt.Println("about to exit")
}

func print(c <-chan int) {
	fmt.Println("in print ")

	// the below implementation is wrong, as channel is read twicr
	// for range c {
	// 	val, ok := <-c
	// 	if !ok {
	// 		fmt.Print("not ok")
	// 	}
	// 	fmt.Println("got value ", val)
	// }

	// approach 1: this is correct implementation
	// for val := range c { // range automatically consumes values from the channel
	// 	fmt.Println("got value", val)
	// }

	// approach 2 : reqading channel using select
	for {
		select {
		case val, ok := <-c:

			if !ok {
				fmt.Print("not ok")
				return
			}

			fmt.Println("got value", val)

		}
	}
}

func gen(q chan<- int) {

	go func() {
		for i := 0; i < 12; i++ {
			fmt.Println("sending on channel q ")
			q <- i
		}
		fmt.Println("closing on channel")
		close(q)
	}()

	fmt.Println("RETURN from gen channel")

}
