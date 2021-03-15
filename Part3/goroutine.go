package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("Going in GO")

	go f("goroutine v2")

	go func(msg string) {
		fmt.Println(msg)
	}("Going in GO v2")

	time.Sleep(time.Second)
	fmt.Println("Done")
}
