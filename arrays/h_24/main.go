package main

import (
	"fmt"
	"math/rand"
)

func main() {

	a := rand.Intn(250)
	fmt.Printf("a = %d \t \n", a)

	if a <= 100 {
		fmt.Printf("case 1")
	} else if a <= 200 {
		fmt.Printf("case 2")
	} else if a < 250 {
		fmt.Printf("case 3")
	}

	fmt.Printf("%d  \n", rand.Intn(3))
	fmt.Printf("%d  \n", rand.Intn(3))
	fmt.Printf("%d  \n", rand.Intn(3))

	fmt.Printf("%d  \n", rand.Intn(3))
	fmt.Printf("%d  \n", rand.Intn(3))
	fmt.Printf("%d  \n", rand.Intn(3))
	fmt.Printf("%d  \n", rand.Intn(3))

}
