package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "abc"
	s[1] = "def"
	s[2] = "ghi"
	fmt.Println("after assignment:", s)
	fmt.Println("index 2:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "xyz")
	s = append(s, "123", "678")
	fmt.Println("append:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("slice from to:", l)

	l = s[:5]
	fmt.Println("slice upto:", l)

	l = s[2:]
	fmt.Println("slice from:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)
	t = append(t, "sdjhf")
	fmt.Println("dcl:", t)

	twoD := make([][]int, 5)
	for i := 0; i < 5; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
