package main

import "fmt"

func main() {

	var x interface{}
	fmt.Println(x)

	x = 2
	fmt.Println(x)
	
	x = 3.14
	fmt.Println(x)

	x = "A"
	fmt.Println(x)

	x = [3]int{5,6,7}
	fmt.Println(x)

}