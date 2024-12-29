package main

import "fmt"

func main() {
	x := doMath(1, 3, add)
	fmt.Println("x is ", x)

	fmt.Printf("%T \n", add)
	fmt.Printf("%T \n", substract)
	fmt.Printf("%T \n", doMath)
}

func doMath(a int, b int, f func(int, int) int) int {

	return f(a, b)

}

func add(a int, b int) int {
	return a + b
}

func substract(a int, b int) int {
	return a - b
}

// func (r receinver) identifier (p paramsetrs )  (r returns) {code}
