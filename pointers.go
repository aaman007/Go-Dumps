package main

import "fmt"

func zeroval(n int) {
	n = 0
}

func zeroptr(ptr *int) {
	*ptr = 0
}

func main() {
	val := 10

	zeroval(val)
	fmt.Println(val)

	zeroptr(&val)
	fmt.Println(val)

	fmt.Println(&val)
}
