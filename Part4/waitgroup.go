package main

import (
	"fmt"
	"time"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Worker", id, "started work")
	time.Sleep(time.Second)
	fmt.Println("Worker", id, "finished work")
}

func main() {
	var wg sync.WaitGroup

	for i:=1; i<=5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
}