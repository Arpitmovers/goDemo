package main

import (
	"fmt"
	"strconv"
)

type book struct {
	title string
}

// attach this function as a method to the type using a reciver  (b book)
func (b book) String() string {
	return fmt.Sprint("title of book is ", b.title)
}

type count int

// implement string() interface
func (c count) String() string {
	return fmt.Sprint("title of book is ", strconv.Itoa(int(c)))

}

func main() {

	b := book{
		title: "West",
	}

	var n count = 42

	fmt.Println(b)
	fmt.Println(n)

}
