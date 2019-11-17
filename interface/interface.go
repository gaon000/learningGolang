package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type Rect struct {
	width, height float64
}
type Circle struct {
	radius float64
}

func (r Rect) area() float64 {
	return r.width * r.height
}
func (r Rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	r := Rect{10., 20.}
	c := Circle{10}
	showArea(r, c)
	var x interface{}
	x = 1
	x = "Tom" //dynamic type
	printIt(x)
	var a interface{} = 1
	i := a       // dynamic type
	j := a.(int) // type int, value is 1
	fmt.Println(i)
	fmt.Println(j)
}
func showArea(shapes ...Shape) {
	for _, s := range shapes {
		a := s.area()
		fmt.Println(a)
	}
}

func printIt(v interface{}) {
	fmt.Println(v)
}
