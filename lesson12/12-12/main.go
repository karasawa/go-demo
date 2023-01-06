package main

import (
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("A", "ABC")
	fmt.Println(match)

	re1, _ := regexp.Compile("A")
	m1 := re1.MatchString("ABC")
	fmt.Println(m1)

	re2 := regexp.MustCompile("A")
	m2 := re2.MatchString("ABC")
	fmt.Println(m2)

}