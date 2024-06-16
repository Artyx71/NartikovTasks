package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	var precision int

	fmt.Println("Введите значение x:")
	fmt.Scan(&x)
	fmt.Println("Введите точность (количество знаков после запятой):")
	fmt.Scan(&precision)

	epsilon := 1 / math.Pow(10, float64(precision))
	result := 1.0
	term := 1.0
	n := 1.0

	for {
		term *= x / n
		if math.Abs(term) < epsilon {
			break
		}
		result += term
		n++
	}

	fmt.Printf("e^%.10f с точностью до %d знаков после запятой: %.*f\n", x, precision, precision, result)
}
