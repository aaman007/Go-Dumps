package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func adder(a, b, c int) int {
	return a + b + c
}

func main() {
	fmt.Println(add(5, 7))
	fmt.Println(adder(10, 12, 23))
}
