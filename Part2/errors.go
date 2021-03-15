package main

import (
	"errors"
	"fmt"
)

func f1(val int) (int, error) {
	if val == 5 {
		return -1, errors.New("Value cannot be 5")
	}
	return val * val, nil
}

type valError struct {
	val  int
	prob string
}

func (e *valError) Error() string {
	return fmt.Sprintf("%d - %s", e.val, e.prob)
}

func f2(val int) (int, error) {
	if val == 5 {
		return -1, &valError{val, "Value cannot be 5"}
	}
	return val * val, nil
}

func main() {
	for _, num := range []int{5, 10} {
		if r, e := f1(num); e != nil {
			fmt.Println("failed", e)
		} else {
			fmt.Println("worked", r)
		}
	}

	for _, num := range []int{5, 10} {
		if r, e := f2(num); e != nil {
			fmt.Println("failed", e)
		} else {
			fmt.Println("worked", r)
		}
	}

	_, e := f2(5)
	if ae, ok := e.(*valError); ok {
		fmt.Println(ae)
		fmt.Println(ae.val)
		fmt.Println(ae.prob)
	}
}
