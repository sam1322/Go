package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	fmt.Println("hello")
	res := plus(1, 2)

	res2 := plusPlus(1, 2, 3)
	fmt.Println("1 + 2 =", res)
	fmt.Println("1 + 2 + 3 =", res2)
}
