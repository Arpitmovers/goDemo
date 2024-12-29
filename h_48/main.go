/*
Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional
slice:
"James", "Bond", "Shaken, not stirred"
"Miss", "Moneypenny", "I'm 008."
Range over the records, then range over the data in each record.
https://go.dev/play/p/2jxGMyXiZrE
*/

package main

import (
	"fmt"
)

// learn range

func main() {

	fmt.Print("---------------------------------------\n")
	a := []string{"James", "Bond", "Shaken, not stirred"}
	b := []string{"Miss", "Moneypenny", "I'm 008."}

	matrix := [][]string{a, b}

	for index, value := range matrix {
		fmt.Println(index, value)
	}

}
