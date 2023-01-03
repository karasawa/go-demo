package main

import (
	"fmt"
)

func reciever(c chan int) {
	for {
		x := <-c
		fmt.Println(x)
	}
}

func reciever2(name string, ch chan int) {
	for {
		y, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(name, y)
	}
	fmt.Println(name, "END")
}

func main() {
	ch1 := make(chan int, 5)
	ch1 <- 1
	ch1 <- 2
	fmt.Println(len(ch1))

	i := <- ch1
	fmt.Println(i)
	fmt.Println(<- ch1)

	// ch2 := make(chan int)
	// ch3 := make(chan int)

	// go reciever(ch2)
	// go reciever(ch3)

	// x := 0
	// for x < 100 {
	// 	ch2 <- x
	// 	ch3 <- x
	// 	time.Sleep(50 * time.Millisecond)
	// 	x++
	// }

	// ch4 := make(chan int)

	// go reciever2("goroutine1", ch4)
	// go reciever2("goroutine2", ch4)
	// go reciever2("goroutine3", ch4)

	// y := 0
	// for y < 100 {
	// 	ch4 <- y
	// 	y++
	// }
	// close(ch4)
	// time.Sleep(3 * time.Second)

	// ch5 := make(chan int, 3)
	// ch5 <- 1
	// ch5 <- 2
	// ch5 <- 3
	// close(ch5)
	// for v := range ch5 {
	// 	fmt.Println(v)
	// }

	ch6 := make(chan int, 2)
	ch7 := make(chan string, 2)

	ch7 <- "A"

	select {
	case v1 := <-ch6:
		fmt.Println(v1 + 1000)
	case v2 := <-ch7:
		fmt.Println(v2 + "!!")
	}
}