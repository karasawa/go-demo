package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)
	fmt.Println(math.Sqrt2)

	fmt.Println(math.MaxFloat32)
	fmt.Println(math.SmallestNonzeroFloat32)

	fmt.Println(math.MaxFloat64)
	fmt.Println(math.SmallestNonzeroFloat64)

	fmt.Println(math.MaxInt64)
	fmt.Println(math.MinInt64)

	fmt.Println(math.Abs(100))
	fmt.Println(math.Abs(-100))

	fmt.Println(math.Pow(0, 2))
	fmt.Println(math.Pow(3, 2))

	fmt.Println(math.Max(1, 2))
	fmt.Println(math.Min(1, 2 ))

	fmt.Println(math.Trunc(3.14))
	fmt.Println(math.Trunc(-3.14))

	fmt.Println(math.Floor(3.14))
	fmt.Println(math.Floor(-3.14))

	fmt.Println(math.Ceil(3.14))
	fmt.Println(math.Ceil(-3.14))

}