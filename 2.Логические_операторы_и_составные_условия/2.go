package main

import (
	"fmt"
)

func main() {
	var num1, num2, num3 float64
	fmt.Println("Введите первое число:")
	fmt.Scan(&num1)
	fmt.Println("Введите второе число:")
	fmt.Scan(&num2)
	fmt.Println("Введите третье число:")
	fmt.Scan(&num3)
	
	if num1 > 0 || num2 > 0 || num3 > 0 {
		fmt.Println("Хотя бы одно из чисел положительное.")
	} else {
		fmt.Println("Все числа отрицательные или нулевые.")
	}
}
