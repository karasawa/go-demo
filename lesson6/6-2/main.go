package main

import "fmt"

func main() {
	var m = map[string]int{
		"A": 100,
		"B": 200,
	}
	fmt.Println(m)

	m2 := make(map[int]string)
	fmt.Println(m2)

	m2[1] = "JAPAN"
	m2[2] = "USA"
	fmt.Println(m2)

	s, ok := m2[3]
	if !ok {
		fmt.Println("err!!!")
	}
	fmt.Println(s, ok)

	delete(m2, 2)
	fmt.Println(m2)

	m3 := map[string]int{
		"Apple": 100,
		"Banana": 200,
	}

	for k, v := range m3 {
		fmt.Println(k, v)
	}
}