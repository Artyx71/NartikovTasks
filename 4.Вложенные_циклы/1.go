package main

import (
	"fmt"
	"strconv"
)

func isMirror(number int) bool {
	str := strconv.Itoa(number)
	return str == reverse(str)
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	count := 0
	for i := 100000; i <= 999999; i++ {
		if isMirror(i) {
			count++
		}
	}
	fmt.Printf("Количество зеркальных билетов среди всех шестизначных номеров от 100000 до 999999: %d\n", count)
}
