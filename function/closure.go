package main

import "fmt"

func main() {
	a, b := outer()
	fmt.Println("inner", a())
	fmt.Println("outer", b)
}
func outer() (func() int, int) {
	outer_var := 2
	inner := func() int {
		outer_var += 99
		return outer_var
	}
	inner()
	return inner, outer_var
}
