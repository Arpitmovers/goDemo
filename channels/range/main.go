package main

import "fmt"

func main() {

	c := make(chan int)

	// send
	go func() {

		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()

	// recieve -ranger is pulling values from channel ,until it is closed
	// if we dont close the channle , it keeps waiting for more values to come
	for v := range c {
		fmt.Println("value is ", v)
	}

	fmt.Println("about to exit")
}
