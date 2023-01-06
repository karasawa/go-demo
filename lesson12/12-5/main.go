package main

import (
	"flag"
	"fmt"
)

func main() {
	//go run main.go -n 20 -m message -x

	var (
		max int
		msg string
		x   bool
	)

	flag.IntVar(&max, "n", 32, "処理数の最大値")
	flag.StringVar(&msg, "m", "", "処理メッセージ")
	flag.BoolVar(&x, "x", false, "拡張オプション")

	flag.Parse()

	fmt.Printf("処理数の最大値 =", max)
	fmt.Printf("処理メッセージ =", msg)
	fmt.Printf("拡張オプション =", x)
}