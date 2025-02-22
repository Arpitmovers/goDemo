// Starting with this code, pull the values off the channel using a select statement
// solution: https://play.golang.org/p/FulKBY5JNj

package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func receive(c, q <-chan int) {
	fmt.Println("in receive ")
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-q:

			fmt.Println("gor from  q channel")
			return
		}

	}
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 12; i++ {
			fmt.Println("sending on channel c ")
			c <- i
		}
		q <- 1
		fmt.Println("closing on channel")
		close(c)
	}()
	fmt.Println("RETURN from gen channel")
	return c
}
