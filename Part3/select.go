package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)
	c3 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "From channel 1 after 2 secs"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		c2 <- "From channel 2 after 1 secs"
	}()

	go func() {
		time.Sleep(3 * time.Second)
		c3 <- "From channel 3 after 3 secs"
	}()

	for i:=1; i<=3; i++ {
		select {
		case message := <-c1:
			fmt.Println(message)
		case message := <-c2:
			fmt.Println(message)
		case message := <-c3:
			fmt.Println(message)
		}
	}
}