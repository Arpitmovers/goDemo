package main

import "fmt"

// learn range

func main() {

	// slice of integer
	xi := []int{1, 3}

	fmt.Printf("type is  %T \n", xi)
	for k, v := range xi {
		fmt.Printf("key %v , value %v \n", k, v)
	}

	// slice of string
	xs := []string{"garden pizza", "alfredo parta", "lamb steak"}

	for k, v := range xs {
		fmt.Printf("key %v , value %v \n", k, v)
	}

	// blank identrifer
	for _, v := range xs {
		fmt.Printf(" blank idetifer value is %v \n", v)
	}

	// access slice using for loop

	for i := 0; i < len(xi); i++ {
		fmt.Printf("for loop %v \n", xi[i])
	}

	fmt.Print("___________________________")
	fmt.Printf("\n")
	// append to slice
	xi = append(xi, 5, 6)
	fmt.Println("slice after append is ", xi)

	// slice a slice --> cut parts of a slice

	yi := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// [inclusive:exclusuve]
	fmt.Printf("yi-->%v", yi[0:2])
}
