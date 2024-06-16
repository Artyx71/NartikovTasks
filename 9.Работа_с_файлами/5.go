package main

import (
	"fmt"
)

func generateParenthesis(n int) []string {
	var result []string
	var backtrack func(left, right int, str string)
	backtrack = func(left, right int, str string) {
		if len(str) == n*2 {
			result = append(result, str)
			return
		}
		if left < n {
			backtrack(left+1, right, str+"(")
		}
		if right < left {
			backtrack(left, right+1, str+")")
		}
	}
	backtrack(0, 0, "")
	return result
}

func main() {
	var n int
	fmt.Println("Введите количество пар скобок:")
	fmt.Scan(&n)
	fmt.Println(generateParenthesis(n))
}
