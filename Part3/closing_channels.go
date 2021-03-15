package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			val, more := <-jobs

			if more {
				fmt.Println("Received job", val)
			} else {
				fmt.Println("Received all jobs")
				done <- true
				break
			}
		}
	}()

	for i:=1; i<=3; i++ {
		fmt.Println("Sending Job", i)
		jobs <- i
		time.Sleep(time.Second)
	}

	fmt.Println("Sent all jobs")
	close(jobs)
	<-done
}