package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 1
	fmt.Println(i)

	fl64 := float64(i)
	fmt.Println(fl64)

	fmt.Printf("%T\n", i)
	fmt.Printf("%T\n", fl64)

	var s string = "100"
	fmt.Printf("%T\n", s)

	i2, _ := strconv.Atoi(s)
	fmt.Printf("%T\n", i2)

	var i3 int = 300
	s2 := strconv.Itoa(i3)
	fmt.Printf("%T, %T\n", i3, s2)

	var h string = "Hello World"
	b := []byte(h)
	fmt.Println(b)

	h2 := string(b)
	fmt.Println(h2)
}