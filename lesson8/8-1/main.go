package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	var user1 User
	fmt.Println(user1)
	user1.Name = "user1"
	user1.Age = 10
	fmt.Println(user1)

	user2 := User{}
	fmt.Println(user2)
	user2.Name = "user2"
	user2.Age = 20
	fmt.Println(user2)

	user3 := User{Name: "user3", Age: 30}
	fmt.Println(user3)

	user4 := User{"user4", 40}
	fmt.Println(user4)

	user5 := new(User)
	fmt.Println(user5)

	user6 := &User{}
	fmt.Println(user6)
}