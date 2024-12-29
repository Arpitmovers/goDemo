package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string
	Age  int
}

func main() {

	p1 := person{
		Name: "arpit",
		Age:  24,
	}

	p2 := person{
		Name: "money penny",
		Age:  24,
	}

	people := []person{p1, p2}

	fmt.Println(people)

	byteslice, error := json.Marshal(people)
	if error != nil {
		fmt.Print("error", error)
	}

	fmt.Print(string(byteslice))
}
