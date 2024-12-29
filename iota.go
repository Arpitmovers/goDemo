package main

import "fmt"

const (
	a = iota
	b
	c
)

const (
	_ = iota
	y
	z
	z1
	z2
)

func main() {

	fmt.Printf("a  %d, b %d ,  c %d \n", a, b, c)
	///	fmt.Printf("x  %d, y %d ,  z %d", x, y, z)

	fmt.Printf("%d \t %d", y, 1<<y)

	fmt.Printf("%d \t %d", z, 2<<z)

	fmt.Printf("%d \t %d", z1, 2<<z1)

}
