package main

import "fmt"

func main() {
	mp := make(map[string]bool)

	mp["Hello"] = true
	mp["NotHello"] = false

	fmt.Println("map", mp)

	val := mp["Hello"]
	fmt.Println("value", val)

	fmt.Println(len(mp))

	delete(mp, "Hello")
	fmt.Println(mp)

	val, isValid := mp["Hello"]
	fmt.Println(isValid, val)

	val2, isValid2 := mp["NotHello"]
	fmt.Println(isValid2, val2)

	newMap := map[string]int{"k1": 10, "k2": 20}
	fmt.Println(newMap)
}
