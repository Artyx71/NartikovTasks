package main

import (
	"fmt"
)

func main() {
	a, b, c := 0, 0, 0
	
	for a < 10 || b < 100 || c < 1000 {
		if a < 10 {
			a++
		}
		if b < 100 {
			b++
		}
		if c < 1000 {
			c++
		}
		fmt.Printf("a: %d, b: %d, c: %d\n", a, b, c)
	}
}
