package main

import (
	"fmt"
	"math"
	"math/rand"
)

const s string = "constant"

var r = "st"

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func zeroval(ival int) {
	ival = 0
}
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	testing3()
}

func testing3() {
	i := 1
	fmt.Println("initial ", i)

	zeroval(i)
	fmt.Println("zeroval ", i)

	zeroptr(&i)
	fmt.Println("zeroptr ", i, &i)
}

func testing2() {
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())

}

func testing() {
	fmt.Println(s, r)
	fmt.Println("Hello, World!")
	// fmt.Println(quote.Go())

	const n = 50000000000000000

	const d = 3e20 / n

	fmt.Println(d, n, int64(d), math.Sin(n))

	fmt.Println(math.Pi)
	str := make([]int, 3)
	str2 := make([]int, 3)
	// var ctr int = 0
	ctr := 0
	for i := 0; i < 10; i++ {
		j := rand.Intn(10)
		ctr += j
		fmt.Print(j, " ")
		str = append(str, j)
		str2 = append(str2, ctr)
	}
	fmt.Println()
	f := "Apple"
	fmt.Println(f)

	total := sum(str...)
	fmt.Println("total sum = ", total)
	fmt.Println(str)
	fmt.Println(str2)
}
