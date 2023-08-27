package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	messages := make(chan string)
	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("goroutine")
		messages <- "ping"
	}()

	go func() {
		defer wg.Done()
		fmt.Println("goroutine")
		// msg := <-messages
		// fmt.Println("goroutine26", msg)
		// messages <- "ping"
	}()

	msg := <-messages
	fmt.Println(msg)
	wg.Wait()

	// channel buffering
	messages = make(chan string, 2)
	messages <- "buffered"
	messages <- "channel"
	// messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)

	// creates capacity of 2
	ch := make(chan int, 2)
	go write(ch)
	time.Sleep(2 * time.Second)
	for v := range ch {
		fmt.Println("read value", v, "from ch")
		time.Sleep(2 * time.Second)

	}

	done := make(chan bool, 1)
	go worker(done)

	fmt.Println(<-done)

}

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func write(ch chan int) {
	for i := 0; i < 4; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}
