package main

import "fmt"

func main() {

	even := make(chan int)

	odd := make(chan int)
	quit := make(chan int)

	// send
	go send(even, odd, quit)

	// recieve
	recieve(even, odd, quit)
	fmt.Println("about to exit")
}

func send(e, o, q chan<- int) {
	for i := 0; i < 3; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
	close(q)
	q <- 11111
}

func recieve(e, o, q <-chan int) {

	for {

		select {
		case v := <-e:
			fmt.Println("from the even channel", v)

		case v := <-o:
			fmt.Println("from the odd channel", v)

		case v := <-q:
			fmt.Println("from teh quit channel", v)
			return
		}
	}
}
