package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := strings.Join([]string{"A", "B", "C"}, ",")
	s2 := strings.Join([]string{"A", "B", "C"}, "")
	fmt.Println(s1, s2)

	i1 := strings.Index("ABCDE", "A")
	i2 := strings.Index("ABCDE", "D")
	fmt.Println(i1, i2)

	i3 := strings.LastIndex("ABCDABCD", "BC")
	fmt.Println(i3)

	i4 := strings.IndexAny("ABC", "ABC")
	i5 := strings.LastIndexAny("ABC", "ABC")
	fmt.Println(i4, i5)

	b1 := strings.HasPrefix("ABC", "A")
	b2 := strings.HasSuffix("ABC", "C")
	fmt.Println(b1, b2)

	b3 := strings.Contains("ABC", "B")
	b4 := strings.ContainsAny("ABCDE", "BD")
	fmt.Println(b3, b4)

	i6 := strings.Count("ABCDEABC", "B")
	i7 := strings.Count("ABCDEABC", "")
	fmt.Println(i6, i7)

	s3 := strings.Repeat("ABC", 4)
	fmt.Println(s3)

	s4 := strings.Replace("AAAAA", "A", "B", 1)
	s5 := strings.Replace("AAAAA", "A", "B", -1)
	fmt.Println(s4, s5)

	s6 := strings.Split("A,B,C,D", ",")
	fmt.Println(s6)
	s7 := strings.SplitAfter("A,B,C,D", ",")
	fmt.Println(s7)
	s8 := strings.SplitN("A,B,C,D", ",", 2)
	fmt.Println(s8)
	s9 := strings.SplitAfterN("A,B,C,D", ",", 3)
	fmt.Println(s9)

	s10 := strings.ToLower("ABC")
	s11 := strings.ToUpper("ab")
	fmt.Println(s10, s11)

	s12 := strings.TrimSpace("   -   A   -   ")
	s13 := strings.TrimSpace("　　　a　　　")
	fmt.Println(s12, s13)

	s14 := strings.Fields("a b c")
	fmt.Println(s14)
	fmt.Println(s14[1])

}