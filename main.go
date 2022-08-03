package main

import (
	"Go-Functions/simplemath"
	"fmt"
)

func main() {
	answer, err := simplemath.Divide(6, 3)
	if err != nil {
		fmt.Printf("An errors occurred %s\n", err.Error())
	} else {
		fmt.Printf("%f\n", answer)
	}
	total := simplemath.Sum(12.2, 14, 16, 22.4)
	fmt.Printf("total of sum: %f\n", total)
}
