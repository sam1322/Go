package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2} // has type Vertex
	v2 = Vertex{X: 1} // Y:0 is implicit
	v3 = Vertex{}     // X:0 and Y:0
	// p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	// p := &Vertex{1, 2}
	// fmt.Println(v1, p, v2, v3)

	// array literal
	p := [3]int{1, 2, 3}
	fmt.Println(p)
	// can't append in array literal
	// p = append(p[0:1], 4)

	// slice literal
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)
	q = append(q, 17)
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)
	r = append(r, false)
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}
