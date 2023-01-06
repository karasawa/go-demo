package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%s\n", "Hello")
	fmt.Printf("%#v\n", "Hello")
	fmt.Printf("%T\n", "Hello")

	s := fmt.Sprint("Hello")
	s1 := fmt.Sprintf("%v\n", 100)
	s2 := fmt.Sprintln("Hello")

	fmt.Println(s)
	fmt.Println(s1)
	fmt.Printf("%T\n", s1)
	fmt.Println(s2)

	fmt.Fprint(os.Stdout, "Hello")
	fmt.Fprint(os.Stdout, "Hello")
	fmt.Fprint(os.Stdout, "Hello")

	f, _ := os.Create("test.txt")
	defer f.Close()

	fmt.Fprint(f, "Hello")
}