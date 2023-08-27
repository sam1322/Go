package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

func main() {
	// basically we are composing or mixing two structs as one
	// and therefore all the methods of base struct become accessible in the derived or child struct
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	fmt.Println("co:", co)
	fmt.Printf("co={num:%v,str:%v}", co.num, co.str)
	fmt.Println("also num:", co.base.num)
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	co.base.num = 2
	var d describer = co
	fmt.Println("Container:", co)
	fmt.Println("describer:", d.describe())
}
