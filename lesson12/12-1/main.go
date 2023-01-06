package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// os.Exit(1)
	// fmt.Println("start")

	// _, err := os.Open("A.txt")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// fmt.Println(os.Args[0])
	// fmt.Println(os.Args[1])
	// fmt.Println(os.Args[2])
	// fmt.Println(os.Args[3])

	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	// s, _ := os.Create("foo.txt")
	// s.Write([]byte("Hello\n"))
	// s.WriteAt([]byte("Golang"), 6)
	// s.Seek(0, os.SEEK_END)
	// s.WriteString("Year")

	f, e := os.Open("foo.txt")
	if e != nil {
		log.Fatalln(e)
	}
	bs := make([]byte, 128)
	fmt.Println(string(bs))
	n, err := f.Read(bs)
	fmt.Println(n)
	fmt.Println(bs)
	fmt.Println(string(bs))
}