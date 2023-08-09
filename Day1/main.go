package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hi")

	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("four")
	}

	if num := 9; num < 10 {
		fmt.Println(num, "is less than 10")
	}
	// fmt.Println(num)

	fmt.Println(time.Now().Weekday())

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")

	default:
		fmt.Println("It a weekday")
	}

	t := time.Now()
	fmt.Println(t.Hour())
	switch {
	case t.Hour() < 4 || t.Hour() > 18:
		fmt.Println("It's night time")
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'am a bool")
		case int:
			fmt.Println("I'am an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(2)
	whatAmI("Hey")
}
