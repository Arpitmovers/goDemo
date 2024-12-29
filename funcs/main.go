package main

import (
	"fmt"
)

// np params , no return
func foo() {
	fmt.Println("In foo")
}

// 1 param , 1 return
func bar(s string) {
	fmt.Println("name is ", s)
}

// 1 param , 1 return
func demon(s string) string {
	fmt.Println("In demon ", s)
	return "demon" + s
}

func dogBarks(name string, age int) (string, int) {
	age = age * 6

	return fmt.Sprint("todf is years olf"), age
}

func sumCalculte(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Printf("type of nums is %T", nums)

	return sum
}

type person struct {
	name string
}

func (s person) speak() {
	fmt.Println(" i am ", s.name)
}

func main() {

	a := person{
		name: "arpit",
	}
	a.speak()

	foo()

	bar("hello")

	s1 := demon("weird")

	fmt.Sprint("s1 is ", s1)

	s1, n := dogBarks("arpit", 34)
	fmt.Println("s1 ", s1, "n is ", n)

	result := sumCalculte(1, 3, 4)
	fmt.Println("result is ", result)
	//fmt.Printf("%T", result)
	result1 := sumCalculte()
	fmt.Println("result1 is ", result1)

}
