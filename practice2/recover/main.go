package main

import "fmt"

func mayPanic() {
	panic("a problem")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered Successfully! Error encountered:", r)
		}
	}()
	mayPanic()
	fmt.Println("After mayPanic()") //this function won't run after panic
	// as function aborts execution
}
