package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mutex   sync.Mutex // to protect shared resource access (queue).
	queue   []int
	cond    = sync.NewCond(&mutex)
	maxSize = 5
)

func produce(id int) {

	for i := 0; i < 10; i++ {
		mutex.Lock()

		if len(queue) == maxSize {
			fmt.Print("produer is waiting ")
			cond.Wait() // releases the mutex and suspends the producer until another goroutine calls
		}

		queue = append(queue, i)
		fmt.Printf("producer %d produced %d \n", id, i)

		cond.Signal() /// Signal the consumer that an item is available
		mutex.Unlock()
		time.Sleep(time.Millisecond * 100)
	}
}

func consume(id int) {
	mutex.Lock()

	if len(queue) == 0 {
		fmt.Print("consumer is waiting ")
		cond.Wait() // wait if queue is empty
	}
	item := queue[0]
	queue = queue[1:]
	fmt.Printf("consumer %d , consumed %d \n", id, item)
	cond.Signal()
	mutex.Unlock()
	time.Sleep(time.Millisecond * 500)

}
func main() {
	go produce(1)
	go consume(1)
	//go consume(2)
	time.Sleep(10 * time.Second) // wait for 10 seconds
}
