package main

import "fmt"

func main() {
	hello := "Hello"
	ref(&hello)
	fmt.Println(hello)
}
func ref(a *string) {
	fmt.Println(*a)
	*a = "Bye"
}
