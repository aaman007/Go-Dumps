package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"zz", "aa", "ccc"}
	sort.Strings(strs)
	fmt.Println(strs)

	ints := []int{5, 1, 3, 2}
	sort.Ints(ints)
	fmt.Println(ints)

	is_sorted := sort.IntsAreSorted(ints)
	fmt.Println("is sorted", is_sorted)
}