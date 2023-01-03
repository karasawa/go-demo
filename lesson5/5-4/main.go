package main

import "fmt"

func main() {
	n := 1
	switch n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("I don't know")
	}

	switch n := 2; n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("I don't know")
	}

	m := 6
	switch {
	case m > 0 && m < 4:
		fmt.Println("0 < m < 4")
	case m > 3 && m < 7:
		fmt.Println("3 < m < 7")
	}

}