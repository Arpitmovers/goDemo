package main

import "fmt"

func intDelta(n *int) {

	*n = 43

}

func main() {

	x := 42

	fmt.Println("x =", x)
	fmt.Println("address of x", &x)

	y := &x
	fmt.Println("y =", y)
	fmt.Println("*y =", *y)

	fmt.Printf("%v \t %T ", y, y)
	// y is a pointer to int

	a := 42
	// pass
	intDelta(&a)
}
