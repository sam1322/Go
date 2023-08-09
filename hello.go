package main

import (
	"fmt"
	"math"
	"math/rand"
)

const s string = "constant"

var r = "st"

func main() {
	fmt.Println(s, r)
	fmt.Println("Hello, World!")
	// fmt.Println(quote.Go())

	const n = 50000000000000000

	const d = 3e20 / n

	fmt.Println(d, n, int64(d), math.Sin(n))

	fmt.Println(math.Pi)
	for i := 0; i < 10; i++ {
		fmt.Print(rand.Intn(10), " ")
	}
	fmt.Println()
	f := "Apple"
	fmt.Println(f)

}
