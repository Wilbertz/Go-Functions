package main

import (
	"Go-Functions/simplemath"
	"math"
)

type MathExpr = string

const (
	AddExpr      = MathExpr("add")
	SubtractExpr = MathExpr("subtract")
	MultiplyExpr = MathExpr("multiply")
)

func main() {
	var funcs []func() int64
	for i := 0; i < 10; i++ {
		cleanI := i
		funcs = append(funcs, func() int64 {
			return int64(math.Pow(float64(cleanI), 2))
		})
	}
	for _, f := range funcs {
		println(f())
	}
}

func powerOfTwo() func() int64 {
	x := 1.0
	return func() int64 {
		x += 1
		return int64(math.Pow(x, 2))
	}
}

func mathExpression(expr MathExpr) func(float64, float64) float64 {
	switch expr {
	case AddExpr:
		return simplemath.Add
	case SubtractExpr:
		return simplemath.Subtract
	case MultiplyExpr:
		return simplemath.Multiply
	default:
		return func(f1 float64, f2 float64) float64 {
			return 0
		}
	}
}

func double(f1, f2 float64, mathExpr func(float64, float64) float64) float64 {
	return mathExpr(f1, f2) * 2
}
