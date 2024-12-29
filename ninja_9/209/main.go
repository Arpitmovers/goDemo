/*
●
 in addition to the main goroutine, launch two additional goroutines
○ each additional goroutine should print something out
● use waitgroups to make sure each goroutine finishes before your program exists
*/

package main

import "fmt"

type person struct {
	fname string
	lName string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println(p.fname, "  ", "__________")
}

func saySomeThing(h human) {
	h.speak()
}

func main() {

	p1 := person{
		fname: "arpit",
		lName: "garg",
	}

	saySomeThing(&p1)
	// works fine directly ,
	p1.speak()

}
