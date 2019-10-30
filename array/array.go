package main

import "fmt"

func main() {
	var a = [3]int{1, 2, 3}
	for _, v := range a {
		fmt.Println(v)
	}
	var b = [...]int{1, 2, 3} //compiler figures out array length
	for _, v := range b {
		fmt.Println(v)
	}
}
