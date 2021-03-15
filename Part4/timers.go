package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	time.Sleep(2 * time.Second)
	stop := timer2.Stop()
	if stop {
		fmt.Println("Timer 2 stopped")
	} else {
		fmt.Println("Could not stop timer as it executed already", stop)
	}
}