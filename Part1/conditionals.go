package main

import "fmt"

func main() {
	var a, b = 3, 7

	if a > b {
		fmt.Println("A is greater")
	} else {
		fmt.Println("B is greater")
	}

	if num := 10; num < 5 {
		fmt.Println("Less then 5")
	} else if num < 10 {
		fmt.Println("Less then 10")
	} else {
		fmt.Println("Greater than 9")
	}
}
