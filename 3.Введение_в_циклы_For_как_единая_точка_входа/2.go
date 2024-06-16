package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Println("Введите первое число:")
	fmt.Scan(&a)
	fmt.Println("Введите второе число:")
	fmt.Scan(&b)
	
	sum := a
	for sum < a + b {
		sum++
		fmt.Println(sum)
	}
}
