package main

import "fmt"

type person struct { //public -> Person
	name string
	age  int
}
type dict struct {
	data map[int]string
}

func newDict() *dict {
	d := dict{}
	d.data = map[int]string{}
	return &d
}
func main() {
	p := person{}
	p.name = "Lee"
	p.age = 10
	var p2 person
	p2 = person{"Kim", 20}
	p3 := person{name: "Jean", age: 32}
	p4 := new(person)
	p4.name = "Dean"

	fmt.Println(p)
	fmt.Println(p2)
	fmt.Println(p3)
	fmt.Println(*p4)
	dic := newDict()
	dic.data[1] = "A"
	fmt.Println(*dic)
}
