package main

import (
	"log"
	"os"
)

func main() {
	log.SetOutput(os.Stdout)

	log.Print("Log\n")
	log.Print("Log2")
	log.Printf("Log%d\n", 3)

	// log.Fatal("Log")
	// log.Fatalln("Log2")
	// log.Fatalf("Log%d\n", 3)

	// f, err := os.Create("sample.txt")
	// if err != nil {
	// 	log.Fatal("err!!")
	// }
	// log.SetOutput(f)
	// log.Print("Log")

	log.SetFlags(log.LstdFlags)
	log.Println("A")

	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	log.Println("B")

	log.SetFlags(log.Ltime | log.Lshortfile)
	log.Println("C")

	log.SetFlags(log.LstdFlags)
	log.SetPrefix("[LOG]")
	log.Println("D")

	file , _ := os.Open("text.txt")
	logger := log.New(file, "[LOG]", log.LstdFlags)
	logger.Println("message1")
	logger.Println("message2")
	log.Println("message3")

}