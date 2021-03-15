package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() {
		messages <- "ping pong"
	}()

	message := <-messages
	fmt.Println(message)
}