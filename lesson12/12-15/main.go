package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type A struct{}

type User struct {
	ID      int			`json:"id"`
	Name    string		`json:"name"`
	Email   string		`json:"email"`
	Created time.Time	`json:"created"`
	A A					`json:"a"`
}

func main() {
	u := new(User)
	u.ID = 1
	u.Name = "test"
	u.Email = "example@example.com"
	u.Created = time.Now()

	bs, err := json.Marshal(u)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(bs))

	u2 := new(User)

	if err := json.Unmarshal(bs, u2); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(u2)
}