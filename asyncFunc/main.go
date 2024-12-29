package main

import "fmt"

func main() {

	func() {
		fmt.Println("anonmous func")
	}()

	func(name string) {
		fmt.Println("anonmous func", name)
	}("arpit")

	// func expresssion
	x := func() {
		fmt.Println("x func")
	}

	x()
}

// func (r receinver) identifier (p paramsetrs )  (r returns) {code}
