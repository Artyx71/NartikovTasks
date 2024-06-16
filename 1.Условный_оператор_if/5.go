package main

import (
	"fmt"
)

func main() {
	var dayOfWeek, guests int
	var checkAmount float64
	
	fmt.Println("Введите день недели:")
	fmt.Scan(&dayOfWeek)
	fmt.Println("Введите число гостей:")
	fmt.Scan(&guests)
	fmt.Println("Введите сумму чека:")
	fmt.Scan(&checkAmount)
	
	discount := 0.0
	serviceCharge := 0.0
	
	if dayOfWeek == 1 {
		discount = checkAmount * 0.10
		fmt.Printf("Скидка по понедельникам: %.2f\n", discount)
	}
	
	if dayOfWeek == 5 && checkAmount > 10000 {
		discount += checkAmount * 0.05
		fmt.Printf("Скидка по пятницам: %.2f\n", checkAmount * 0.05)
	}
	
	if guests > 5 {
		serviceCharge = checkAmount * 0.10
		fmt.Printf("Надбавка на обслуживание: %.2f\n", serviceCharge)
	}
	
	totalAmount := checkAmount - discount + serviceCharge
	fmt.Printf("Сумма к оплате: %.2f\n", totalAmount)
}
