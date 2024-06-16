package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int16
	fmt.Println("Введите два числа (int16):")
	fmt.Scan(&a, &b)

	result := int32(a) * int32(b)
	fmt.Printf("Результат умножения: %d\n", result)

	switch {
	case result >= 0 && result <= math.MaxUint8:
		fmt.Println("Минимальный тип данных для хранения результата: uint8")
	case result >= math.MinInt8 && result <= math.MaxInt8:
		fmt.Println("Минимальный тип данных для хранения результата: int8")
	case result >= 0 && result <= math.MaxUint16:
		fmt.Println("Минимальный тип данных для хранения результата: uint16")
	case result >= math.MinInt16 && result <= math.MaxInt16:
		fmt.Println("Минимальный тип данных для хранения результата: int16")
	case result >= 0 && result <= math.MaxUint32:
		fmt.Println("Минимальный тип данных для хранения результата: uint32")
	case result >= math.MinInt32 && result <= math.MaxInt32:
		fmt.Println("Минимальный тип данных для хранения результата: int32")
	default:
		fmt.Println("Результат слишком велик для стандартных целочисленных типов.")
	}
}
