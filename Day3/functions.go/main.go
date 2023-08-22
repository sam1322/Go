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

	a := []int{1, 2, 3}

	var v []int
	v = make([]int, 3)
	v = append(v, 1, 2, 3, 4)
	v[0] = 1
	x := v[:1]
	a = append(a, 1, 2, 3, 4)

	fmt.Println("array:", a)
	fmt.Println("Slice:", v, x, v[0:3], a[0:2], "sum:", plusPlus(90, 20, -10))
}
