package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)
	if err != nil {
		log.Fatalln("err ", err)
	}
	fmt.Println(string(bs))

}

// toJSON needs to return an error also
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {

		//return []byte{}, fmt.Errorf("there was an error %v", err)
		//	return []byte{}, fmt.Errorf("there was an error %v", err)
		//return []byte{}, fmt.Println("there was an error %v", err) // cant use this as it prints & returns no of bytes written,error

		return []byte{}, errors.New(fmt.Sprintf("error in marshal", err))
	}

	return bs, nil

}
