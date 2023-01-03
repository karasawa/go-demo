package main

import "fmt"

func main() {
	var s string = "Hello Go"
	fmt.Println(s)

	fmt.Println(`test
	test
	        test`)

	fmt.Println("\"")
	fmt.Println(`"`)
	fmt.Println(s[0])
	fmt.Println(string(s[0]))
}