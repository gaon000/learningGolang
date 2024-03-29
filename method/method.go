package main

import "fmt"

type Rect struct {
	width, height int
}

func (r Rect) area() int {
	return r.width * r.height
}

func (r *Rect) area2() int {
	r.width++
	return r.width * r.height
}

func main() {
	rect := Rect{10, 20}
	area := rect.area()
	fmt.Println(area)
	area2 := rect.area2()
	fmt.Println(rect.width, area2)
}
