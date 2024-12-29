/*
Create a map with a key of TYPE string which is a person’s “last_first” name, and a value of
TYPE []string which stores their favorite things. Store three records in your map. Print out all of
the values, along with their index position in the slice.
*/

package main

import (
	"fmt"
)

// learn range

func main() {

	fmt.Print("---------------------------------------\n")

	m1 := make(map[string][]string)

	m1[`bond_james`] = []string{`shaken, not stirred`, `martinis`, `fast cars`}

	m1[`moneypenny_jenny`] = []string{`intelligence`, `literature`, `computer science`}

	for k, v := range m1 {
		fmt.Println("key is ", k)
		for idx, value := range v {
			fmt.Println("index is ", idx, "value is ", value)
		}
	}

	delete(m1, `bond_james`)
	delete(m1, `moneypenny_jenny`)
	fmt.Println("AFTER DELETE OF ")
	for k, v := range m1 {
		fmt.Println("key is ", k)
		for idx, value := range v {
			fmt.Println("index is ", idx, "value is ", value)
		}
	}

	// matrix := [][]string{a, b}

	// for index, value := range matrix {
	// 	fmt.Println(index, value)
	// }

}
