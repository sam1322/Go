package main

import (
	"fmt"
)

func main() {
	// Arrays
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	fmt.Println("len", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	var twoD [2][3]int

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d : ", twoD)

	//Slices
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	fmt.Println(s)
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "t"}
	fmt.Println("dcl:", t)

	TwoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		TwoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			TwoD[i][j] = i + j

		}

	}
	fmt.Println("2d:", TwoD)

	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 10
	m["k4"] = 7
	m["k5"] = 7
	fmt.Println("map", m)

	v1 := m["k1"]

	fmt.Println("v1", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2")

	fmt.Println("map", m)

	fmt.Println("len:", len(m))

	A, B := m["k2"]
	fmt.Println("a:", A, "b:", B)
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	// Range
	nums := []int{1, 2, 3}
	sum := 0
	for _, num := range nums {
		sum += num
		fmt.Println(num, sum)
	}

	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index", i)

		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	fmt.Println("kvs:", kvs)
	for k, v := range kvs {
		fmt.Printf("%s -> %s \n", k, v)
	}

	for k := range kvs {
		fmt.Println("keys", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}

}
