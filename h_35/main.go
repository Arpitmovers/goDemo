package main

import (
	"fmt"
	"math/rand"
)

// learn range

func main() {
	// xi := []int{42, 43, 44, 45, 46, 47}

	// for idx, val := range xi {
	// 	fmt.Printf("index %v , val %v \n", idx, val)
	// }

	fmt.Println("--------------------")

	mp := map[string]int{
		"arpit":  30,
		"saloni": 23,
	}

	// for key, value := range mp {
	// 	fmt.Println(key, value)
	// }

	fmt.Println("--------------------")

	m1 := mp["arpit"]

	fmt.Printf("age of arpit is %v \n", m1)

	m1 = mp["test"]

	fmt.Println("age of test is ", m1)

	if v, ok := mp["test"]; !ok {
		fmt.Println("there is no  test in m1  and this is the zero value for int ", v, ok)
	}

	obj := map[string]string{
		"name": "arpit",
		"age":  "34",
	}
	// comma ok idiom
	if value, ok := obj["address"]; !ok {

		fmt.Println("value not present for address ", value, ok)
	}

	if name, ok := obj["name"]; ok {
		fmt.Println("name present", name, ok)
	}

	cnt := 0
	fmt.Println("--------------------")

	// statement ; statement idiom
	for i := 0; i < 100; i++ {
		if h := rand.Intn(5); h == 3 {
			cnt++
			fmt.Printf(" i = %d ,h is 3 \n", i)
		}
	}
	fmt.Printf(" total cnt is %v", cnt)

}
