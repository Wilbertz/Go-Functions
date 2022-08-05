package main

func main() {
	addExpr := mathExpression()

	println(addExpr(2, 3))
}

func mathExpression() func(float64, float64) float64 {
	return func(f1 float64, f2 float64) float64 {
		return f1 + f2
	}
}
