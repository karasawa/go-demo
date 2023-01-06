package main

import "fmt"

type MyError struct {
	Message string
	ErrCode int
}

func (e *MyError) Error() string {
	return e.Message
}

func RaiseError() error {
	return &MyError{Message: "カスタムエラーが発生しました", ErrCode: 1234}
}

func main() {
	e := RaiseError()
	fmt.Println(e.Error())

	err, ok := e.(*MyError)
	if ok {
		fmt.Println(err.ErrCode)
	}
	
}