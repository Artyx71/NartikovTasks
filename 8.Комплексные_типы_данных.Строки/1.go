package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	input := "Go is an Open source programming Language that makes it Easy to build simple, reliable, and efficient Software."

	words := strings.Fields(input)
	count := 0

	for _, word := range words {
		if len(word) > 0 && unicode.IsUpper(rune(word[0])) {
			count++
		}
	}

	fmt.Println("Определение количества слов, начинающихся с большой буквы в строке:")
	fmt.Println(input)
	fmt.Printf("Строка содержит %d слов с большой буквы.\n", count)
}
