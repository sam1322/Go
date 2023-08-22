package main

import (
	"fmt"
	"math/rand"
)

// multiple return values
func vals() (int, int) {
	return rand.Intn(10), rand.Intn(5)
}

// variadic functions
func sum(a int, nums ...int) {
	fmt.Println("a:", a)
	fmt.Println("nums:", nums)
}

func main() {
	fmt.Println(vals())
	sum(1, 2, 3, 34, 4, 5)

}
