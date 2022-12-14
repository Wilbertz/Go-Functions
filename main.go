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
	if err := ReadFullFile(); err != nil {
		println("something bad occurred")
	}
}

func ReadFullFile() (err error) {
	var r io.ReadCloser = &SimpleReader{}
	defer func() {
		r.Close()
		if p := recover(); p == errCatastrophicReader {
			println(p)
			err = errors.New("a panic occurred but it is ok")
		}
	}()
	for {
		value, readerErr := r.Read([]byte("text that does nothing"))
		if readerErr == io.EOF {
			println("finished reading file, breaking out of loop")
			break
		} else if readerErr != nil {
			err = readerErr
			return
		}
		println(value)
	}
	return nil
}

func ReadSomethingBad() error {
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

type SimpleReader struct {
	count int
}

var errCatastrophicReader = errors.New("something catastrophic occurred in the reader")

func (br *SimpleReader) Read(p []byte) (n int, err error) {
	if br.count == 2 {
		panic(errCatastrophicReader)
	}
	if br.count > 3 {
		return 0, io.EOF
	}
	br.count += 1
	return br.count, nil
}

func (br *SimpleReader) Close() error {
	println("closing reader")
	return nil
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
