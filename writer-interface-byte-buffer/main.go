package main

import (
	"bytes"
	"fmt"
)

func main() {

	b := bytes.NewBufferString("Hello")
	fmt.Println("b->", b)

	b.WriteString("Gophes")
	fmt.Println("b->", b)

	b.Reset()

	fmt.Println("b->", b)

	b.WriteString("its dec")

	fmt.Println("b->", b)

	b.Write([]byte(" happy happy"))

	fmt.Println("b->", b)
}
