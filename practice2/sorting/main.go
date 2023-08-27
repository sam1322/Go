package main

// import "golang.org/x/exp/slices"
import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	strs := []string{"c", "b", "a"}
	slices.Sort(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 5}
	slices.Sort(ints)
	fmt.Println("Ints:", ints)

	s := slices.IsSorted(ints)
	fmt.Println("sorted:", s)

	// Sorting by functions

	fruits := []string{"peach", "banana", "kiwi"}

	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits)

	type Person struct {
		name string
		age  int
	}

	people := []Person{
		Person{name: "Jax", age: 37},
		Person{name: "CJ", age: 25},
		Person{name: "Mike", age: 72},
	}

	slices.SortFunc(people,
		func(a, b Person) int {
			return cmp.Compare(a.age, b.age)
		})

	fmt.Println(people)
}
