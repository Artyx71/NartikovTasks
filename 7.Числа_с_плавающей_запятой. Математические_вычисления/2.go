package main

import (
	"fmt"
	"math"
)

func main() {
	var initialAmount, monthlyRate float64
	var years int

	fmt.Println("Введите начальную сумму вклада (в рублях):")
	fmt.Scan(&initialAmount)
	fmt.Println("Введите ежемесячный процент капитализации:")
	fmt.Scan(&monthlyRate)
	fmt.Println("Введите количество лет вклада:")
	fmt.Scan(&years)

	months := years * 12
	rate := monthlyRate / 100
	totalAmount := initialAmount
	lostAmount := 0.0

	for i := 0; i < months; i++ {
		newAmount := totalAmount * (1 + rate)
		roundedAmount := math.Floor(newAmount * 100) / 100
		lostAmount += newAmount - roundedAmount
		totalAmount = roundedAmount
	}

	fmt.Printf("Сумма на счете по окончании вклада: %.2f руб.\n", totalAmount)
	fmt.Printf("Сумма, потерянная за счет округления: %.10f руб.\n", lostAmount)
}
