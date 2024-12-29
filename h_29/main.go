package main

import (
	"fmt"
)

func init() {
	fmt.Println("inside init block")
}
func main() {

	// for i := 0; i <= 99; i++ {
	// 	fmt.Println(i)
	// }

	// for i := 0; i <= 99; i++ {
	// 	fmt.Println(rand.Intn(10))
	// }

	// for i := 0; i < 42; i++ {
	// 	x := rand.Intn(5)
	// 	fmt.Printf(" iteration is %d no is %d  \n", i, x)
	// }

	//z := 0
	// for z < 10 {
	// 	fmt.Printf(" z is %d \n", z)
	// 	z++
	// }

	zz := 5

	for {
		if zz == 3 {
			break
		}
		fmt.Printf("zz is %d \n", zz)
		zz--
	}

	for i := 1; i < 100; i++ {

		if i%2 != 0 {
			fmt.Printf("%d  is odd no \n", i)
		}
	}
}
