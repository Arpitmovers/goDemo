package main

import "fmt"

func main() {

	f := incrementor()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	g := incrementor()
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
}

// func (r receinver) identifier (p paramsetrs )  (r returns) {code}

func incrementor() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
