package main

import (
	"fmt"
)

func main() {
	var price, discount float64
	fmt.Println("Введите цену товара:")
	fmt.Scan(&price)
	fmt.Println("Введите скидку в процентах:")
	fmt.Scan(&discount)
	
	if discount > 30 {
		discount = 30
	}
	
	discountAmount := price * discount / 100
	if discountAmount > 2000 {
		discountAmount = 2000
	}
	
	fmt.Printf("Сумма скидки: %.2f рублей\n", discountAmount)
}
