package main

import (
	"fmt"
)

func main() {
	var guess, response int
	
	for attempts := 1; attempts <= 4; attempts++ {
		guess = (attempts - 1) * 2 + 1
		fmt.Printf("Попытка %d: ваше число %d?\n", attempts, guess)
		fmt.Println("Введите 1, если угадала, 2 если больше, 3 если меньше:")
		fmt.Scan(&response)
		
		if response == 1 {
			fmt.Println("Программа угадала ваше число!")
			return
		}
	}
	
	fmt.Println("Программа не смогла угадать ваше число за 4 попытки.")
}
