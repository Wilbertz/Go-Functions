package main

import "fmt"

func main() {
	a := func(name string) string {
		fmt.Printf("my first %s function\n", name)
		return name
	}
	value := a("function 1")
	println(value)
}
