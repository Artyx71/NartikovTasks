package main

import (
	"fmt"
)

func main() {
	var num1, num2, num3 int
	fmt.Println("Введите первое число:")
	fmt.Scan(&num1)
	fmt.Println("Введите второе число:")
	fmt.Scan(&num2)
	fmt.Println("Введите третье число:")
	fmt.Scan(&num3)
	
	count := 0
	if num1 >= 5 {
		count++
	}
	if num2 >= 5 {
		count++
	}
	if num3 >= 5 {
		count++
	}
	
	fmt.Printf("Среди введённых чисел %d больше или равны 5.\n", count)
}
