package main

import (
	"fmt"
)

func main() {
	var ticket int
	fmt.Println("Введите четырёхзначный номер билета:")
	fmt.Scan(&ticket)
	
	d1 := ticket / 1000
	d2 := (ticket / 100) % 10
	d3 := (ticket / 10) % 10
	d4 := ticket % 10
	
	if d1 == d4 && d2 == d3 {
		fmt.Println("Зеркальный билет")
	} else if d1 + d2 == d3 + d4 {
		fmt.Println("Счастливый билет")
	} else {
		fmt.Println("Обычный билет")
	}
}
