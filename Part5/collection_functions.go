package main

import (
	"fmt"
	"strings"
)

func Index(s []string, t string) int {
	for i, v := range(s) {
		if v == t {
			return i
		}
	}
	return -1
}

func Include(s []string, t string) bool {
	return Index(s, t) >= 0
}

func Any(s []string, f func(string) bool) bool {
	for _, v := range(s) {
		if f(v) {
			return true
		}
	}
	return false
}

func All(s []string, f func(string) bool) bool {
	for _, v := range(s) {
		if !f(v) {
			return false
		}
	}
	return true
}

func Filter(s []string, f func(string) bool) []string {
	newS := make([]string, 0)
	for _, v := range(s) {
		if f(v) {
			newS = append(newS, v)
		}
	}
	return newS
}


func Map(s []string, f func(string) string) []string {
	newS := make([]string, len(s))
	for i, v := range(s) {
		newS[i] = f(v)
	}
	return newS
}

func main() {
	var strs = []string{"peach", "apple", "pear", "plum"}

		
    fmt.Println(Index(strs, "pear"))
	fmt.Println(Include(strs, "grape"))
	fmt.Println(Include(strs, "apple"))
    fmt.Println(Any(strs, func(v string) bool {
        return strings.HasPrefix(v, "p")
    }))
    fmt.Println(All(strs, func(v string) bool {
        return strings.HasPrefix(v, "p")
    }))
    fmt.Println(Filter(strs, func(v string) bool {
        return strings.Contains(v, "e")
    }))

	fmt.Println(Map(strs, strings.ToUpper))
}grape