// Starting with this code, marshal the []user to JSON. There is a little bit of a curve ball in this
// hands-on exercise - remember to ask yourself what you need to do to EXPORT a value from a
// package.
package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	first string
	age   int
}

func main() {
	u1 := user{
		first: "James",
		age:   32,
	}

	// u2 := user{
	// 	First: "Moneypenny",
	// 	Age:   27,
	// }

	// u3 := user{
	// 	First: "M",
	// 	Age:   54,
	// }

	users := []user{u1}

	// fmt.Println(users)
	// fmt.Printf("%T", users)
	// your code goes here

	byteString, err := json.Marshal(users)
	if err != nil {
		fmt.Print("error ", err)
	}

	fmt.Println("bytes", string(byteString))
}
