package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range(jobs) {
		fmt.Println("worker", id, "received job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j*j
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for i:=1; i<=5; i++ {
		fmt.Println("Started worker", i)
		go worker(i, jobs, results)
	}

	for i:=1; i<=numJobs; i++ {
		fmt.Println("Sent job", i)
		jobs <- i
	}
	close(jobs)

	for i:=1; i<=numJobs; i++ {
		<-results
	}
}