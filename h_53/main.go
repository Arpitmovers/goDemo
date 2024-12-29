// Create your own type “person” which will have an underlying type of “struct” so that it can store
// the following data:
// ● first name
// ● last name
// ● favorite ice cream flavors
// Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice
// which stores the favorite flavors.

package main

import "fmt"

func main() {

	type person struct {
		firstName string
		lastName  string
		favIc     []string
	}

	p1 := person{
		firstName: "Arpit",
		lastName:  "Garg",
		favIc:     []string{"passion fruit", "guava", "strawberry"},
	}

	p2 := person{
		firstName: "Jenny",
		lastName:  "Penny",
		favIc:     []string{"strawberry", "chocolate"},
	}

	fmt.Println("p1 = ", p1)
	fmt.Println("p2 = ", p2)

	for _, value := range p1.favIc {
		fmt.Println(p1.firstName, " ", value)
	}

	for _, value := range p2.favIc {
		fmt.Println(p2.firstName, " ", value)
	}
}
