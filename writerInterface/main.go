package main

import (
	"log"
	"os"
)

type person struct {
	name string
}

// w implements io.Writer interface
// func (p person) writeOut(w io.Writer) {
// 	_, err := w.Write([]byte(p.person))

// 	if err != nil {

// 	}

// }

func main() {

	ptr, err := os.Create("op.txt")

	if err != nil {
		log.Fatalf("error is %s ", err)
	}
	defer ptr.Close()

	// slice of byttes
	s := []byte("hello gophers")

	_, error := ptr.Write(s)
	if error != nil {
		log.Fatalf("error is %s ", error)
	}
}
