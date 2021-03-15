package main

import "fmt"

func main() {
	nums := []int{2, 3, 6}
	fmt.Println(nums)

	var sum int = 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)

	for i, num := range nums {
		fmt.Println(i, num)
	}

	mp := map[string]string{"key1": "val1", "key2": "val2"}
	for key, val := range mp {
		fmt.Println(key, val)
	}

	for key := range mp {
		fmt.Println(key)
	}

	for i, c := range "golang" {
		fmt.Println(i, c)
	}
}
