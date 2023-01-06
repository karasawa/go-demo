package main

import (
	"fmt"
	"sort"
)

type Entry struct {
	Name string
	Value int
}

func main() {
	i := []int{2, 4, 1, 6, 5, 7, 3}
	s := []string{"c", "q", "a"}

	fmt.Println(i, s)

	sort.Ints(i)
	sort.Strings(s)

	fmt.Println(i, s)

	el := []Entry{
		{"A", 10},
		{"s", 20},
		{"O", 70},
		{"n", 10},
		{"k", 40},
		{"W", 60},
		{"d", 20},
		{"b", 90},
	}

	fmt.Println(el)

	sort.Slice(el, func(i, j int) bool {return el[i].Name < el[j].Name})

	fmt.Println("----------")
	fmt.Println(el)
	fmt.Println("----------")

	sort.SliceStable(el, func(i, j int) bool {return el[i].Name < el[j].Name})

	fmt.Println("----------")
	fmt.Println(el)
	fmt.Println("----------")


}