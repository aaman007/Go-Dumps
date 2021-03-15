package main

import (
	"os"
	"fmt"
)

func createFile(s string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(s)

	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File, data string) {
	fmt.Println("writing")
	fmt.Fprintln(f, data)
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
	}
}

func main() {
	f := createFile("data.txt")
	defer closeFile(f)
	writeFile(f, "my data")
}