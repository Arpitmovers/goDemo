/*
●
 Using a COMPOSITE LITERAL:
● create a SLICE of TYPE int
● assign these 10 VALUES
42, 43, 44, 45, 46, 47, 48, 49, 50, 51
● Range over the slice and print the values out.
○ print out the VALUE and the TYPE
*/

package main

import (
	"fmt"
)

// learn range

func main() {

	// slice of integer
	xi := [10]int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	// for _, v := range xi {
	// 	fmt.Printf("value is %d , type is %T  \n", v, v)
	// }

	fmt.Printf("%v \n", xi)
	fmt.Printf("%#v \n", xi)
	fmt.Print("---------------------------------------\n")

	/*
		● [42 43 44 45 46]
		● [47 48 49 50 51]
		● [44 45 46 47 48]
		● [43 44 45 46 47]
	*/
	s1i := xi[:5]
	s2i := xi[5:]
	s3i := xi[2:7]
	s4i := xi[1:6]
	fmt.Printf("%v \n", s1i)
	fmt.Printf("%v \n", s2i)
	fmt.Printf("%v \n", s3i)
	fmt.Printf("%v \n", s4i)
}
