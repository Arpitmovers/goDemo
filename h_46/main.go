/*
To DELETE from a slice, we use APPEND along with SLICING. For this hands-on exercise,
follow these steps:
● start with this slice
○ x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
● use APPEND & SLICING to get these values here which you should ASSIGN to the
variable “x” and then print:
○ [42, 43, 44, 48, 49, 50, 51]
*/

package main

import (
	"fmt"
)

// learn range

func main() {

	// slice of integer
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Printf("%#v \n", x)
	fmt.Print("---------------------------------------\n")

	x = append(x, 52)
	fmt.Printf("%#v \n", x)

	x = append(x, 53, 54, 55)
	fmt.Printf("%#v \n", x)

	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Printf("%#v \n", x)

	// 	To DELETE from a slice, we use APPEND along with SLICING. For this hands-on exercise,
	// follow these steps:
	// ● start with this slice
	// ○ x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	// ● use APPEND & SLICING to get these values here which you should ASSIGN to the
	// variable “x” and then print:
	// [42, 43, 44, 48, 49, 50, 51]

	z := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	z = append(z[:3], z[6:]...)
	fmt.Println(" array after append and slice is ", z)

	// 47th exerise
	z2 := make([]string, 0, 50)
	fmt.Println("z2 len is ", len(z2), " capacity is ", cap(z2))

	// z2 has capacity of 50 , but length 0
	z2 = append(z2, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, `
			Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, `
			Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, `
			Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`,
		` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`,
		` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, `
			Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)

	fmt.Printf("z2 is %#v", z2)

	fmt.Println("z2 len is ", len(z2), "capacity is ", cap(z2))

	z3 := make([]string, 12)
	fmt.Println("z3 len is ", len(z3), "capacity is ", cap(z3))

	for i := 0; i < len(z2); i++ {
		fmt.Println(z2[i], i)
	}

}
