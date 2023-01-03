package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	m := map[int]User{
		1: {Name: "user1", Age: 10},
		2: {Name: "user2", Age: 20},
		3: {Name: "user3", Age: 30},
	}
	fmt.Println(m)
	fmt.Println(m[1].Name)

	m3 := make(map[int]User)
	fmt.Println(m3)

	m3[1] = User{Name: "user1", Age: 10}
	m3[2] = User{Name: "user2", Age: 20}
	fmt.Println(m3)
}