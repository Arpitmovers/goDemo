// Starting with this code, sort the []user by
// ● age
// ● last
// Todd McLeod - Learn To Code Go - Page 129
// Also sort each []string “Sayings” for each user
// ● print everything in a way that is pleasant

package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []user

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByLast []user

func (last ByLast) Len() int           { return len(last) }
func (last ByLast) Swap(i, j int)      { last[i], last[j] = last[j], last[i] }
func (last ByLast) Less(i, j int) bool { return last[i].Last < last[j].Last }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	//fmt.Println(users)

	// your code goes here
	sort.Sort(ByAge(users))
	for _, userV := range users {
		fmt.Println(userV)
	}

	fmt.Println("_____________________________________-")
	sort.Sort(ByLast(users))
	for _, userV := range users {
		fmt.Println(userV)
	}
}
