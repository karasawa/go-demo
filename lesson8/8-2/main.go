package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u User) SayName() {
	fmt.Println(u.Name)
}

func (u User) SetName(name string) {
	u.Name = name
}

func (u *User) SetName2(name string) {
	u.Name = name
}

func main() {
	user1 := User{Name: "user1", Age: 10}
	user1.SayName()

	user1.SetName("new_user1")
	fmt.Println(user1)

	user1.SetName2("new_user1")
	fmt.Println(user1)

	user2 := &User{Name: "user2"}
	user2.SetName2("new_user2")
	fmt.Println(user2)
}