package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	r.width = r.width * 2
	return r.width * r.height
}
func (r rect) perim() int {
	r.width = r.width * 2
	return (r.height + r.width) * 2
}

func main() {
	r := rect{width: 10, height: 5}
	fmt.Println("area:", r.area(), r.width, r.height)
	fmt.Println("perim:", r.perim(), r.width, r.height)

	rp := &r
	fmt.Println("area:", rp.area(), r.width, r.height)
	fmt.Println("perim:", rp.perim(), r.width, r.height)
	fmt.Println("HI")

	userName := "Sriram"
	fmt.Printf("Variable is of type %T \n", userName)

	type Vertex struct {
		X int
		Y int
	}

	v := Vertex{1, 2}
	fmt.Println("vertex:", v)
	v.X = 4
	fmt.Println(v.X)
	fmt.Println("vertex:", v)

}
