package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "Value1"
	queue <- "Value2"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}