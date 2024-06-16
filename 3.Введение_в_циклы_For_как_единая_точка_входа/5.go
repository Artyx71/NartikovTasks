package main

import (
	"fmt"
)

func main() {
	var cap1, cap2, cap3 int
	fmt.Println("Введите ёмкость первой корзины:")
	fmt.Scan(&cap1)
	fmt.Println("Введите ёмкость второй корзины:")
	fmt.Scan(&cap2)
	fmt.Println("Введите ёмкость третьей корзины:")
	fmt.Scan(&cap3)
	
	basket1, basket2, basket3 := 0, 0, 0
	
	for basket1 < cap1 || basket2 < cap2 || basket3 < cap3 {
		if basket1 < cap1 {
			basket1++
			fmt.Printf("Корзина 1: %d/%d\n", basket1, cap1)
		}
		if basket2 < cap2 {
			basket2++
			fmt.Printf("Корзина 2: %d/%d\n", basket2, cap2)
		}
		if basket3 < cap3 {
			basket3++
			fmt.Printf("Корзина 3: %d/%d\n", basket3, cap3)
		}
	}
}
