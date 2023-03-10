package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(256)

	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64())

	fmt.Println(time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())

	fmt.Println(rand.Intn(100))
}