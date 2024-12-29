package main

import (
	"fmt"
	"time"
)

func printMessage(msg string) {
	for i := 0; i < 4; i++ {
		fmt.Println(msg, "a")
		time.Sleep(2000 * time.Millisecond)
		fmt.Println(msg, "b")
	}
}

func main() {
	go printMessage("test1")
	//go printMessage("test2")

	time.Sleep(19 * time.Second) // Allow goroutines to finish
	fmt.Println("Main function finished")
}
