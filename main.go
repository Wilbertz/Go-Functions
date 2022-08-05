package main

import "Go-Functions/simplemath"

func main() {
	addExpr := mathExpression()

	println(addExpr(2, 3))
}

func mathExpression() func(float64, float64) float64 {
	return simplemath.Add
}
