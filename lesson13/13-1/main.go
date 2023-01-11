package main

import (
	"fmt"
	"strconv"
)

func PrintSlice[T any](sl []T) {
	for _, v := range sl {
		fmt.Println(v)
	}
}

func f[T fmt.Stringer](sx []T) {
	result := []string{}
	for _, v := range sx {
		result = append(result, v.String())
	}
	fmt.Println(result)
}

type MyInt int

func(i MyInt) String() string {
	return strconv.Itoa(int(i))
}

type Vector[T any] []T

type T[A any, B []C, C *A] struct {
	a A
	b B
	c C
}

func Double(i *int) {
	fmt.Println(*i)
	*i = *i * 2
}

func main() {
	// PrintSlice([]int{1,2,3,4,5})
	// PrintSlice[string]([]string{"a", "b", "c", "d"})

	// f([]MyInt{1,2,3,4,5})

	// var v Vector[int] = Vector[int]{10,20,30}
	// fmt.Println(v)

	// var t T[int, []*int, *int]
	// fmt.Printf("A: %T, B: %T, C: %T", t.a, t.b, t.c)

	var n int = 300
	Double(&n)
	fmt.Println(n)
}