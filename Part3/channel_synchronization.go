package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("Work started")
	time.Sleep(time.Second)
	fmt.Println("Word DONE")
	done <- true
}

func main() {
	done := make(chan bool, 1)

	go worker(done)

	fmt.Println("Waiting for worker to finish")
	<-done
	fmt.Println("Next Worker >>")
}