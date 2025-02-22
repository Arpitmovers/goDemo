package main

import (
	"fmt"
)

// wrong code
// func main() {
// 	cs := make(chan<- int) // this is defined as send only channel

// 	go func() {
// 		cs <- 42
// 	}()
// 	fmt.Println(<-cs)   // this is defined as reciveve only channel

// 	fmt.Printf("------\n")
// 	fmt.Printf("cs\t%T\n", cs)
// }

// corrected code
func main() {
	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs) // this is blocking line

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)

	cr := make(chan int)

	go func() {
		cr <- 42
	}()
	fmt.Println(<-cr)

	fmt.Printf("------\n")
	fmt.Printf("cr\t%T\n", cr)
}
