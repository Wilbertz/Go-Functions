package main

import "fmt"

func main() {
	a := func(name string) {
		fmt.Printf("my first %s function", name)
	}
	a("function 1")
}
