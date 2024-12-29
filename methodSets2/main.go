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

type youngin interface {
	walk()
	run()
}

func youungRun(y youngin) {
	y.run()
}

func youngWalk(y youngin) {
	y.walk()
}

func main() {

	d1 := dog{"henry"}

	d2 := &dog{"Padget"}

	// this wont run as
	// for value that is not a pointer , we cant use pointer reciever
	youungRun(d1)

	// for value that is a pointer , we can use both pointer and
	youungRun(d2)

	youngWalk(d1)
	youngWalk(d2)
}
