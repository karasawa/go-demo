package main

import "fmt"

func Intergers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	i := Intergers()
	fmt.Println(i())
	fmt.Println(i())
	fmt.Println(i())

	s := Intergers()
	fmt.Println(s())
	fmt.Println(s())

	fmt.Println(i())
}