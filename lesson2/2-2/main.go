package main

import "fmt"

func main() {
	//整数型

	var i int = 100
	fmt.Println(i + 50)

	var i2 int64 = 200
	fmt.Println(i2 + 50)

	fmt.Printf("%T\n", i2)
	fmt.Printf("%T\n", int32(i2))
	fmt.Println(i + int(i2))
}