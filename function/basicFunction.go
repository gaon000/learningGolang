package main

import "fmt"

func main() {
	fun1()
	var a = fun2()
	fmt.Println(a)
	var b = fun3(999)
	fmt.Println(b)
}
func fun1() {
	fmt.Println("This is function1")
}
func fun2() string {
	return "It return string"
}
func fun3(a int) int {
	return a
}
