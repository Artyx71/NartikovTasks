package main

import (
	"fmt"
	"math"
)

func main() {
	var countUint8, countUint16 int
	maxUint8 := uint32(math.MaxUint8)
	maxUint16 := uint32(math.MaxUint16)
	maxUint32 := uint32(math.MaxUint32)

	for i := uint32(0); i <= maxUint32; i++ {
		if i > maxUint8 && (i % (maxUint8 + 1)) == 0 {
			countUint8++
		}
		if i > maxUint16 && (i % (maxUint16 + 1)) == 0 {
			countUint16++
		}
	}

	fmt.Printf("Количество переполнений для uint8: %d\n", countUint8)
	fmt.Printf("Количество переполнений для uint16: %d\n", countUint16)
}
