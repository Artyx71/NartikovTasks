package main

import (
	"fmt"
)

func main() {
	var width, height int
	fmt.Println("Введите ширину:")
	fmt.Scan(&width)
	fmt.Println("Введите высоту:")
	fmt.Scan(&height)
	
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if (x+y)%2 == 0 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
