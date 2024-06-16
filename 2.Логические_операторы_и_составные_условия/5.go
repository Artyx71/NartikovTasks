package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Println("Введите коэффициент a:")
	fmt.Scan(&a)
	fmt.Println("Введите коэффициент b:")
	fmt.Scan(&b)
	fmt.Println("Введите коэффициент c:")
	fmt.Scan(&c)
	
	D := b*b - 4*a*c
	fmt.Printf("Дискриминант D = %.2f\n", D)
	
	if D > 0 {
		x1 := (-b + math.Sqrt(D)) / (2 * a)
		x2 := (-b - math.Sqrt(D)) / (2 * a)
		fmt.Printf("Уравнение имеет два различных корня: x1 = %.2f, x2 = %.2f\n", x1, x2)
	} else if D == 0 {
		x := -b / (2 * a)
		fmt.Printf("Уравнение имеет один корень: x = %.2f\n", x)
	} else {
		fmt.Println("Корней нет.")
	}
}
