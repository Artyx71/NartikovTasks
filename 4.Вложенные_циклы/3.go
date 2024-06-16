package main

import (
	"fmt"
)

func main() {
	var height int
	fmt.Println("Введите высоту ёлочки:")
	fmt.Scan(&height)
	
	for i := 0; i < height; i++ {
		// Вывод пробелов
		for j := 0; j < height-i-1; j++ {
			fmt.Print(" ")
		}
		// Вывод звёздочек
		for j := 0; j < 2*i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
