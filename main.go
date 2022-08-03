package main

import "Go-Functions/simplemath"

func main() {
	sv := simplemath.NewSemanticVersion(1, 2, 3)
	println(sv.String())
}
