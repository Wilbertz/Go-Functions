package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	answer, err := divide(6, 3)
	if err != nil {
		fmt.Printf("An errors occurred %s\n", err.Error())
	} else {
		fmt.Printf("%f\n", answer)
	}
	total := sum(12.2, 14, 16, 22.4)
	fmt.Printf("total of sum: %f\n", total)
}

func sum(values ...float64) float64 {
	total := 0.0
	for _, value := range values {
		total += value
	}
	return total
}

func divide(p1, p2 float64) (float64, error) {
	if p2 == 0 {
		return math.NaN(), errors.New("cannot divide by zero")
	}
	return p1 / p2, nil
}

func add(p1, p2 float64) float64 {
	return p1 + p2
}

func subtract(p1, p2 float64) float64 {
	return p1 - p2
}

func multiply(p1, p2 float64) float64 {
	return p1 * p2
}
