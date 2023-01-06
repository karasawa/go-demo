package main

import (
	"20230102/go-demo/lesson11/alib"
	"fmt"
)

func IsOne(i int) bool {
	if i == 1 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(IsOne(1))

	s := []int{1, 2, 3, 4, 5}
	fmt.Println(alib.Avarage(s))
}