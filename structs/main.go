package main

import (
	"fmt"
)

type person struct {
	fName string
	lName string
	age   int
}

type secretagent struct {
	person
	ltk bool
}

func main() {

	p1 := person{
		fName: "arpit",
		lName: "garg",
		age:   33,
	}

	fmt.Printf("%T \n", p1)  // see type
	fmt.Printf("%#v \n", p1) // see type and value

	fmt.Println(p1.age)

	// embedded  structs
	jamesBond := secretagent{
		person: person{
			fName: "james",
			lName: "bond",
			age:   33,
		},
		ltk: true,
	}

	fmt.Printf("type for james bond: %T \n", jamesBond)        // see type
	fmt.Printf("type,value for jamesbond : %#v \n", jamesBond) // see type and value

	// anonymous struc
	buildingInstance := struct {
		floors int
		name   string
	}{
		floors: 14,
		name:   "Aparna Maple",
	}

	fmt.Print("p1 is \n", buildingInstance)
	fmt.Printf("p1 = %T \n	", buildingInstance)
}
