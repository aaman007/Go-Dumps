package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2

	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Friday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekdays")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	myFunc := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("Boolean")
		case string:
			fmt.Println("String")
		case int:
			fmt.Println("Integer")
		default:
			fmt.Println("Unknown")
		}
	}

	myFunc(1)
	myFunc("Hello")
	myFunc(true)
	myFunc(1.4)

}
