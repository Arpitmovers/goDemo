package main

import (
	"fmt"

	"honnef.co/go/tools/analysis/callcheck"
)

func outer() func() int {
	return func() int {
		return 122

	}
}


func square(n int) int{
	return n * n;
}

func main() {

	a := outer()
	fmt.Println("a", a())



	// example of callback
	func printSquare( f func (int) int) int{
		 
		return fmt.Sprintf("area is ",%f(2));
	}	

	x:=printSquare(square)
}
