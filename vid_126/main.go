package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Agee    int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`

	fmt.Println(s)

	var people []person

	// convert string
	err := json.Unmarshal([]byte(s), &people)

	if err != nil {
		fmt.Print("people", people)
	}

	for i, person := range people {
		fmt.Println("i=", i, "\t", person.First, person.Last, person.Agee)

		for j, saying := range person.Sayings {
			fmt.Println("\t\t", j, saying)
		}
	}

}
