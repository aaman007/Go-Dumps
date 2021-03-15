package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "From channel 1 after 2 secs"
	}()


	select {
	case message := <-c1:
		fmt.Println(message)
	case <-time.After(1 * time.Second):
		fmt.Println("1sec timeout for c1")
	}


	c2 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "From channel 2 after 2 secs"
	}()

	select {
	case message := <-c2:
		fmt.Println(message)
	case <-time.After(3 * time.Second):
		fmt.Println("3sec timeout for c2")
	}

}