package main

import "fmt"

type T struct {
	User User
}

type User struct {
	Name string
	Age  int
}

func (u *User) SetName(name string) {
	u.Name = name
}

func main() {
	var user1 = T{User: User{Name: "user1", Age: 10}}
	fmt.Println(user1)
	fmt.Println(user1.User)
	fmt.Println(user1.User.Name)

	user1.User.SetName("new_user1")
	fmt.Println(user1)
	fmt.Println(user1.User)
	fmt.Println(user1.User.Name)
}