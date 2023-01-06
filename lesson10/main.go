package main

import (
	"20230102/go-demo/lesson10/foo"
	"fmt"
	. "fmt"
	f "fmt"
)

func appName() string {
	var AppName string = "GoApp"
	var Version string = "1.0"
	return AppName + " " + Version
}

func main() {
	fmt.Println(foo.Max)
	//fmt.Println(foo.min)
	f.Println(foo.ReturnMin())
	Println(foo.ReturnMin())

	fmt.Println(appName())
}