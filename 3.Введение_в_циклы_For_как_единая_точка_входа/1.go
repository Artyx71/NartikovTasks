package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Println("Введите число:")
	fmt.Scan(&n)
	
	for i := 0; i <= n; i++ {
		fmt.Println(i)
	}
}
