package main

import "fmt"

func main() {
	sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	sum(1, 2)
}

func sum(values ...int) {
	res := 0
	for _, v := range values {
		res = res + v
	}
	fmt.Println(res)
}
