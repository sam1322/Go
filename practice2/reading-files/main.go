package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fileName := "/home/sriram/projects/Go/defer.txt"
	dat, err := os.ReadFile(fileName)
	check(err)
	fmt.Println(string(dat))
}
