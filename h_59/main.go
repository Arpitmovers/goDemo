package main

import "fmt"

// unfurl the p1
func foo(p1 ...int) int {

	s := 0
	for _, val := range p1 {
		s += val
	}
	return s
}

func bar(p1 []int) int {

	s := 0
	for _, val := range p1 {
		s += val
	}
	return s
}

func main() {

	a1 := []int{1, 2, 3, 4}
	fmt.Println(foo(a1...))
	fmt.Println(bar(a1))

}

// create a func with the identifier foo that
// ○ takes in a variadic parameter of type int
// ○ pass in a value of type []int into your func (unfurl the []int)
// ○ returns the sum of all values of type int passed in

// ● create a func with the identifier bar that
// ○ takes in a parameter of type []int
// ○ returns the sum of all values of type int passed in
