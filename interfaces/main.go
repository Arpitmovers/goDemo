package main

import (
	"fmt"
)

type person struct {
	first string
}

type secretAgent struct {
	person
}

func (p person) speak() {
	fmt.Println("I am", p.first)
}

func (sa secretAgent) speak() {
	fmt.Println("i am a secret agent \n", sa.first)
}

// allows to have polymorphism
type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {

	// embedding one struct to other
	p1 := person{
		first: "Arpit",
	}

	p2 := person{
		first: "Kashvi",
	}
	// p1.speak()
	// p2.speak()

	james := secretAgent{
		person: person{
			first: "James",
		},
	}

	// james.speak()
	// p1.speak()
	fmt.Print("demonstate polymorhism")
	saySomething(james)
	saySomething(p1)
	saySomething(p2)

}
