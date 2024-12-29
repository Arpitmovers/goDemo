// Take the code from the previous exercise, then store the VALUES of type person in a map with
// the KEY of last name. Access each value in the map. Print out the values, ranging over the
// slice.

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

	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	// fmt.Println("p1 = ", p1)
	// fmt.Println("p2 = ", p2)

	// for _, value := range p1.favIc {
	// 	fmt.Println(p1.firstName, " ", value)
	// }

	// for _, value := range p2.favIc {
	// 	fmt.Println(p2.firstName, " ", value)
	// }

	for _, val := range m {
		fmt.Println("val is ", val)

		for _, iceCream := range val.favIc {

			fmt.Println("iceCream is ", iceCream)

		}
	}
}
