package main

import "fmt"

func main() {
	m1 := make(map[string]string)
	m2 := make(map[string]string, 1)
	m3 := make(map[string]string, 1)
	m4 := make(map[string]string, 1)
	fmt.Printf("%v\n", m1)
	fmt.Printf("%v\n", &m2)
	fmt.Printf("%v\n", &m3)
	fmt.Printf("%v\n", &m4)
}
