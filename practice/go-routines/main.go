package main

import (
	"fmt"
	"sync"
)

func f(from string) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

var wg sync.WaitGroup

func main() {
	wg.Add(3)
	f("direct")
	go f("goroutine")

	// wg.Add(1)
	go func(msg string) {
		defer wg.Done()
		fmt.Println(msg)
	}("going")

	// time.Sleep(time.Second)

	fmt.Println("done")
	wg.Wait()
}
