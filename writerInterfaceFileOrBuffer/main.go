package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type person struct {
	name string
}

// implements io.Writer interface
func (p person) writeOut(w io.Writer) {
	w.Write([]byte(p.name))
}

func main() {

	p1 := person{
		name: "Jenny",
	}

	ptr, err := os.Create("op.txt")
	if err != nil {
		log.Fatalf("error is %s ", err)
	}
	defer ptr.Close()

	var b bytes.Buffer

	p1.writeOut(ptr)
	p1.writeOut(&b)

	fmt.Println(b.String())
}
