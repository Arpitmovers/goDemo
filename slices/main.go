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
	fmt.Printf("yi-->%v \n", yi[0:2])

	fmt.Printf("yi-->%v \n", yi[1:])

	fmt.Printf("yi-->%v \n", yi[:2])

	fmt.Printf("yi-->%v \n", yi[:])

	// deleting from slice
	// lets delete element at index 5 (4) from yi
	yi = append(yi[0:5], yi[6:]...)
	fmt.Printf("yi after removing element at index 4%v \n", yi)

	fmt.Print("_______________________________ \n")
	// create slice using make function with length 0 and capacity 10
	zi := make([]int, 0, 10)

	fmt.Printf("zi has length and capacity %v %v \n", len(zi), cap(zi))

	zi = append(zi, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	fmt.Printf("zi has length and capacity %v %v \n", len(zi), cap(zi))

	zi = append(zi, 11)
	fmt.Printf("zi has length and capacity %v %v \n", len(zi), cap(zi))

	fmt.Print("_____COPY COMMAND___________ \n")
	ai := []int{1, 2, 3, 4}
	ai[2] = 122
	fmt.Println(" ai is ", ai)
	bi := make([]int, 7)
	copy(bi, ai)

	fmt.Println(" ai is ", ai)
	fmt.Println(" bi is ", bi)

}
