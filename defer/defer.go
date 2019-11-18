package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("asdf.txt")
	if err != nil {
		panic(err) // run all defer
	}
	defer fmt.Println("close the file")
	defer f.Close()

	bytes := make([]byte, 1024)
	f.Read(bytes)
	fmt.Println(len(bytes))
}
