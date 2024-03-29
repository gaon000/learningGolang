package main

import (
	"fmt"
	"os"
)

func main() {
	openFile("asdf.txt")
	fmt.Println("Done")
}
func openFile(fn string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("OPEN ERROR", r)
		}
	}()
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	defer f.Close()
}
