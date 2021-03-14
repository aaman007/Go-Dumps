package main

import "fmt"

type Rect struct {
	width, height int
}

func (r *Rect) area() int {
	return r.width * r.height
}

func (r Rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	rect := Rect{width: 5, height: 4}
	fmt.Println(rect.area())
	fmt.Println(rect.perim())

	rp := &rect
	fmt.Println(rp.area(), rp.perim())
}
