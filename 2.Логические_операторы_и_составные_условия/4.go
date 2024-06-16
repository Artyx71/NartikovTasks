package main

import (
	"fmt"
)

func main() {
	var cost, coin1, coin2, coin3 int
	fmt.Println("Введите стоимость товара:")
	fmt.Scan(&cost)
	fmt.Println("Введите номинал первой монеты:")
	fmt.Scan(&coin1)
	fmt.Println("Введите номинал второй монеты:")
	fmt.Scan(&coin2)
	fmt.Println("Введите номинал третьей монеты:")
	fmt.Scan(&coin3)
	
	if cost == coin1 || cost == coin2 || cost == coin3 || 
		cost == coin1 + coin2 || cost == coin1 + coin3 || cost == coin2 + coin3 || 
		cost == coin1 + coin2 + coin3 {
		fmt.Println("Можно заплатить без сдачи.")
	} else {
		fmt.Println("Нельзя заплатить без сдачи.")
	}
}
