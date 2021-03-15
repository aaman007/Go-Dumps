package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigs := make(chan os.Signal)
	done := make(chan bool)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT) 
	// SIGQUIT: ctrl + \
	// SIGINT: ctrl + c

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println("Signal:", sig)
		done <- true
	}()

	fmt.Println("Wait for signal")
	<-done
	fmt.Println("Exit")
}