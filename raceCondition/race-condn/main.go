package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int

func main() {
	wg.Add(2)
	go incrementor("aaa")
	go incrementor("bbb")
	wg.Wait()
	fmt.Println("final counter ", counter)

}

func incrementor(str string) {
	for i := 0; i < 20; i++ {
		x := counter
		x++
		time.Sleep(time.Duration(133 * time.Millisecond))
		counter = x
		fmt.Println("counter ", counter, "str", str)
	}
	wg.Done()
}
