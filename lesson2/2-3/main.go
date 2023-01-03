package main

import "fmt"

func main() {
	var fl64 float64 = 2.4
	fmt.Println(fl64)

	fl := 3.2
	fmt.Println(fl)
	fmt.Println(fl + fl64)
	fmt.Printf("%T, %T\n", fl64, fl)

	zero := 0.0
	pinf := 1.0 / zero
	fmt.Println(pinf)

	ninf := -1.0 / zero
	fmt.Println(ninf)

	nan := 0.0 / zero
	fmt.Println(nan)
}