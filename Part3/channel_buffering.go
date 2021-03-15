package main

import "fmt"

func main() {
	messages := make(chan string, 2)

	messages <- "channel"
	messages <- "buffered"

	receiver := []string{<-messages, <-messages}
	fmt.Println(receiver)
}