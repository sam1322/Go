package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	start := time.Now()

	timer1 := time.NewTimer(2 * time.Second)

	// <-timer1.C
	// fmt.Println("timer 1 fired")
	wg.Add(2)
	go func() {
		<-timer1.C
		fmt.Println("Timer 1 fired")
		wg.Done()
	}()

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
		wg.Done()
	}()

	stop2 := timer2.Stop()
	if stop2 {
		wg.Done()
		fmt.Println("Timer 2 stopped")
	}
	// waitingTime := 2 * time.Second
	// fmt.Printf("Waiting for %s\n", waitingTime)
	// time.Sleep(waitingTime)
	wg.Wait()
	timeElapsed := time.Since(start)
	fmt.Printf("The `for` loop took %s\n", timeElapsed)
}
