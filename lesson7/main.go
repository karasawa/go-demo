package main

import "fmt"

func Double(i int) {
	i = i * 2
	fmt.Println(&i)
}

func Double2(i *int) {
	*i = *i * 2
}

func Double3(s []int) {
	for k, v := range s {
		s[k] = v * 2
	}
}

func main() {
	var n int = 2
	fmt.Println(n)
	fmt.Println(&n)

	Double(n)

	fmt.Println(n)
	fmt.Println(&n)

	var p *int = &n
	fmt.Println(p)
	fmt.Println(*p)

	Double2(&n)
	fmt.Println(n)

	Double2(p)
	fmt.Println(*p)

	var sl = []int{1,2,3}
	Double3(sl)
	fmt.Println(sl)

}