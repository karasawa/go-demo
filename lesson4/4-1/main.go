package main

import "fmt"

func Plus(x int, y int) int {
	i := x + y
	return i
}

func Div(x int, y int) (int, int) {
	o := x / y
	s := x % y
	return o, s
}

func Double(x int) (result int){
	result = x * 2
	return 
}

func NoReturn(){
	fmt.Println("No Return")
	return 
}

func main() {
	i := Plus(2, 3)
	fmt.Println(i)

	o, s := Div(9, 3)
	fmt.Println(o, s)

	result := Double(5)
	fmt.Println(result)

	NoReturn()
}