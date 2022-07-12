package main

import "fmt"

func main() {
	s1 := "123"
	s2 := ""
	s2 += "123"

	fmt.Printf("%v\n", s1 == s2)
}
