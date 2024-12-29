package main

import "fmt"

// learn range

func main() {

	// slice of integer
	xi := [5]int{}

	for i := 0; i < 5; i++ {
		xi[i] = i
	}

	for _, v := range xi {
		fmt.Printf("value is %d , type is %T  \n", v, v)

	}

}
