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
	
	if num1 > 5 || num2 > 5 || num3 > 5 {
		fmt.Println("Среди введённых чисел есть число больше 5.")
	} else {
		fmt.Println("Среди введённых чисел нет числа больше 5.")
	}
}
