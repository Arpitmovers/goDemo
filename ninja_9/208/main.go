/*
●
 in addition to the main goroutine, launch two additional goroutines
○ each additional goroutine should print something out
● use waitgroups to make sure each goroutine finishes before your program exists
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	fmt.Printf("Worker %d starting\n", id)
	defer wg.Done()

	time.Sleep(time.Second) // Simulate some work with sleep
	fmt.Printf("Worker %d done\n", id)

}

func main() {

	var wg sync.WaitGroup

	noOfWorkers := 5
	wg.Add(noOfWorkers)

	for i := 1; i <= noOfWorkers; i++ {
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Print("DONE ALL")
}
