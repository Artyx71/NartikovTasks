package main

import "fmt"

func lemonadeChange(bills []int) bool {
	five, ten := 0, 0

	for _, bill := range bills {
		switch bill {
		case 5:
			five++
		case 10:
			if five == 0 {
				return false
			}
			five--
			ten++
		case 20:
			if ten > 0 && five > 0 {
				ten--
				five--
			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println("Расчёт сдачи. Примеры:")
	example1 := []int{5, 5, 5, 10, 20}
	fmt.Printf("Ввод: %v\nВывод: %v\n", example1, lemonadeChange(example1))

	example2 := []int{10, 10}
	fmt.Printf("Ввод: %v\nВывод: %v\n", example2, lemonadeChange(example2))

	example3 := []int{5, 5, 10, 10, 20}
	fmt.Printf("Ввод: %v\nВывод: %v\n", example3, lemonadeChange(example3))
}
