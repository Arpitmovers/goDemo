package main

import (
	"fmt"
)

// learn range

func main() {

	a1 := [5]int{1, 3}

	fmt.Println("array a1 ", a1)

	a2 := [...]string{"arpit", "garg"}

	fmt.Println("array a2 ", a2)

	// only declaration
	var c [2]int

	c[0] = 11
	c[1] = 12
	fmt.Println("array c ", c)

	// printing types
	fmt.Printf("type of a1 is  %T \n", a1)
	fmt.Printf("type of a2 is %T \n", a2)

}
