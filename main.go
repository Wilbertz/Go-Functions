package main

import (
	"Go-Functions/simplemath"
	"errors"
	"fmt"
	"io"
	"math"
)

type MathExpr = string

const (
	AddExpr      = MathExpr("add")
	SubtractExpr = MathExpr("subtract")
	MultiplyExpr = MathExpr("multiply")
)

func main() {
	ReadSomething()
}

func ReadSomething() error {
	var r io.Reader = BadReader{errors.New("my nonsense reader")}

	value, err := r.Read([]byte("test something"))

	if err != nil {
		fmt.Printf("an error occurred %s", err)
		return err
	}
	println(value)
	return nil
}

type BadReader struct {
	err error
}

func (br BadReader) Read(p []byte) (n int, err error) {
	return -1, br.err
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
