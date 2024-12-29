package main

import "fmt"

type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Println("my name is ", d.first, "and i am walking")
}

func (d *dog) run() {
	fmt.Println("my name is ", d.first, "and i am running")
}

func main() {

	d1 := dog{"henry"}

	d2 := &dog{"Padget"}

	// d1 has value semantics we can use both value & pointer RECIEVER
	d1.walk()
	d1.run()

	// d2 has value semantics we can use both value & pointer RECIEVER
	d2.walk()
	d2.run()
}
