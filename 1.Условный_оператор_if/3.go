package main

import (
	"fmt"
)

func main() {
	var amount int
	fmt.Println("Введите сумму снятия со счёта:")
	fmt.Scan(&amount)
	
	if amount%100 != 0 {
		fmt.Println("Сумма должна быть кратна 100.")
	} else if amount > 100000 {
		fmt.Println("Максимальная сумма снятия 100 000 рублей.")
	} else {
		fmt.Println("Операция успешно выполнена.")
		fmt.Printf("Вы сняли %d рублей.\n", amount)
	}
}
