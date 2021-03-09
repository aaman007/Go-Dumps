package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("initialized array:", a)

	a[4] = 100
	fmt.Println("array:", a)
	fmt.Println("value on index 4:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4} // Last index will contain 0 as no values are provided
	fmt.Println("array with initial values:", b)

	var twoD [2][3]int
	fmt.Println(len(twoD))
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
