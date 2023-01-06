package main

import (
	"fmt"
	"strconv"
)

func main() {
	bt := true
	fmt.Printf("%T\n", strconv.FormatBool(bt))

	i := strconv.FormatInt(-100, 10)
	fmt.Printf("%v, %T\n", i, i)

	i2 := strconv.Itoa(100)
	fmt.Printf("%v, %T\n", i2, i2)

	i3, _ := strconv.Atoi("200")
	fmt.Printf("%v, %T\n", i3, i3)

	b1, _ := strconv.ParseBool("true")
	fmt.Printf("%v, %T\n", b1, b1)
	b2, _ := strconv.ParseBool("false")
	fmt.Printf("%v, %T\n", b2, b2)
}