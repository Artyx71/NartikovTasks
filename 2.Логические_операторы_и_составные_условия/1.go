package main

import (
	"fmt"
)

func main() {
	var x, y float64
	fmt.Println("Введите координату X:")
	fmt.Scan(&x)
	fmt.Println("Введите координату Y:")
	fmt.Scan(&y)
	
	if x > 0 && y > 0 {
		fmt.Println("Точка находится в первой четверти.")
	} else if x < 0 && y > 0 {
		fmt.Println("Точка находится во второй четверти.")
	} else if x < 0 && y < 0 {
		fmt.Println("Точка находится в третьей четверти.")
	} else if x > 0 && y < 0 {
		fmt.Println("Точка находится в четвертой четверти.")
	} else if x == 0 && y == 0 {
		fmt.Println("Точка находится в начале координат.")
	} else if x == 0 {
		fmt.Println("Точка находится на оси Y.")
	} else if y == 0 {
		fmt.Println("Точка находится на оси X.")
	}
}
