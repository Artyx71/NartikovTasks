package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "a10 10 20b 20 30c30 30 dd"

	fmt.Println("Исходная строка:")
	fmt.Println(input)

	parts := strings.Fields(input)
	numbers := []int{}

	for _, part := range parts {
		if num, err := strconv.Atoi(part); err == nil {
			numbers = append(numbers, num)
		}
	}

	fmt.Println("В строке содержатся числа в десятичном формате:")
	for _, num := range numbers {
		fmt.Print(num, " ")
	}
	fmt.Println()
}
