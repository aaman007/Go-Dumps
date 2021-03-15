package main

import "fmt"

func addSub(a int, b int) (int, int) {
	return a + b, a - b
}

func main() {
	sum, sub := addSub(5, 7)
	fmt.Println(sum, sub)
	fmt.Println(addSub(10, 5))

}
