package main

import (
	"fmt"
)

func isLucky(number int) bool {
	str := fmt.Sprintf("%06d", number)
	sum1 := int(str[0]-'0') + int(str[1]-'0') + int(str[2]-'0')
	sum2 := int(str[3]-'0') + int(str[4]-'0') + int(str[5]-'0')
	return sum1 == sum2
}

func main() {
	previousLucky := -1
	maxDistance := 0
	
	for i := 100000; i <= 999999; i++ {
		if isLucky(i) {
			if previousLucky != -1 {
				distance := i - previousLucky
				if distance > maxDistance {
					maxDistance = distance
				}
			}
			previousLucky = i
		}
	}
	
	fmt.Printf("Минимальное количество билетов, которое нужно купить, чтобы среди них оказался счастливый: %d\n", maxDistance)
}
