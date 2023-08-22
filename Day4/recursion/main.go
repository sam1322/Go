package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func zeroVal(ival int) {
	ival = 0
}

func zeroPtr(iptr *int) {
	*iptr = 0
}

func main() {
	fmt.Println(fact(5))
	i := 1
	fmt.Println("i:", i)

	zeroVal(i)
	fmt.Println("i:", i)

	zeroPtr(&i)
	fmt.Println("i:", i)

	fmt.Println("pointer:", &i)

}
