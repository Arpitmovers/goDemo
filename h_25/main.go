package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("inside init block")
}
func main() {

	x := rand.Intn(250)

	fmt.Printf("x is %d \t 	", x)
	switch {
	case x < 100:
		fmt.Printf("case1 ")

	case x <= 200:
		fmt.Printf("case2 ")
	default:
		fmt.Printf("case 3")
	}
}
