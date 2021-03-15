package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)

	for i:=1; i<=5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(1 * time.Second)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)

	for i:=0; i<3; i++ {
		burstyLimiter <- time.Now()
	}
	go func() {
		for t := range time.Tick(1 * time.Second) {
			burstyLimiter <- t
		}
	}()

	burstRequests := make(chan int, 9)
	for i:=1; i<=9; i++ {
		burstRequests <- i
	}
	close(burstRequests)

	for req := range(burstRequests) {
		<- burstyLimiter
		fmt.Println("request", req, time.Now())
		/**
		go func() {
			time.Sleep(time.Second)
			burstyLimiter <- time.Now()
		}()
		**/
	}
}