package main

import "fmt"

var i5 int = 500

func outer(){
	var s4 string = "outer"
	fmt.Println(s4)
}

func main() {
	//明示的な型定義
	//var 変数 型 = 値 

	var i int = 100
	fmt.Println(i)

	var s string = "Hello Go"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int = 200
		s2 string = "Golang"
	)
	fmt.Println(i2, s2)

	var i3 int
	var s3 string
	fmt.Println(i3, s3)

	i3 = 300
	s3 = "Go"
	fmt.Println(i3, s3)

	i = 150
	fmt.Println(i)

	//暗黙的な型定義
	//変数 := 値

	i4 := 400
	fmt.Println(i4)

	i4 = 450
	fmt.Println(i4)

	fmt.Println(i5)

	outer()
}