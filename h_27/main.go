package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("inside init block")
}
func main() {

	x := rand.Intn(10)
	y := rand.Intn(10)

	fmt.Printf("x is %d y is %d \n", x, y)

	if x < 4 && y < 4 {
		fmt.Println("x and y are less than 4")
	} else if x > 6 && y > 6 {
		fmt.Println("x and y are greater than 6")
	} else if x >= 4 && x <= 6 {
		fmt.Println("x is >=4 and less than 6")
	} else if y != 5 {
		fmt.Println("y is not 5")
	} else {
		fmt.Println("none cases match")
	}

	switch {
	case x < 4 && y < 4:
		fmt.Println("x and y are less than 4")

	case x > 6 && y > 6:
		fmt.Println("x and y are greater than 6")

	case x >= 4 && x <= 6:
		fmt.Println("x is >=4 and less than 6")

	case y != 5:
		fmt.Println("y is not 5")

	default:
		fmt.Println("none cases match")
	}

}
