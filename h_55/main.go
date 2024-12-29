// ●
//
//	Create a type engine struct, and include this field
//
// ○ electric bool
// ● Create a type vehicle struct, and include these fields
// ■ engine
// ■ make
// ■ model
// ■ doors
// ■ color
// ● Create two VALUES of TYPE vehicle
// ○ use a composite literal
// ● Print out each of these values.
// ● Print out a single field from each of these values.
package main

import (
	"fmt"
)

type engine struct {
	electric bool
}

type vehicle struct {
	engine
	make  string
	model string
	doors int
	color string
}

func main() {

	// embedding one struct to other
	v1 := vehicle{
		engine: engine{
			electric: true,
		},
		make:  "sdfds",
		model: "sdfsff",
		doors: 2,
		color: "red",
	}

	v2 := vehicle{
		engine: engine{
			electric: false,
		},
		make:  "sdfds",
		model: "sdfsff",
		doors: 2,
		color: "red",
	}

	fmt.Println("v1 is", v1)
	fmt.Println("v2 is", v2)

	fmt.Println("v2 is", v2.engine, v2.make)

	// 	Create and use an anonymous struct with these fields:
	// ● first string
	// ● friends map[string]int
	// ● favDrinks []string

	s1 := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "test",
		friends: map[string]int{
			"friedn1": 34,
			"frned2":  31,
		},
		favDrinks: []string{"water", "beeru"},
	}

	fmt.Println("s1 ", s1)

}
