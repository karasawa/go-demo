package main

import "fmt"

func Sum(s ...int) int  {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

func main() {
	var sl []int = []int{100, 200}
	fmt.Println(sl)

	sl2 := []string{"A", "B"}
	fmt.Println(sl2)

	sl3 := make([]int, 4)
	fmt.Println(sl3)

	sl3[1] = 1000
	fmt.Println(sl3)

	sl4 := []int{1,2,3,4,5}
	fmt.Println(sl4)

	sl4 = append(sl4, 6)
	fmt.Println(sl4)

	sl4 = append(sl4, 7,8)
	fmt.Println(sl4)

	sl5 := []int{1,2,3}
	sl6 := make([]int, 3)
	sl7 := copy(sl6, sl5)
	fmt.Println(sl7, sl6)

	fmt.Println(Sum(1,2,3,4,5))
	fmt.Println(Sum(sl5...))
	
}