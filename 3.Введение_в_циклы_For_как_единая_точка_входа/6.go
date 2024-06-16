package main

import (
	"fmt"
)

func main() {
	passengers := []int{4, 7, 10}
	liftCapacity := 2
	liftPosition := 0
	totalFloors := 24
	passengersInLift := 0
	passengersDelivered := 0
	
	for passengersDelivered < len(passengers) {
		// Лифт поднимается вверх
		for liftPosition < totalFloors {
			liftPosition++
		}
		fmt.Println("Лифт на верхнем этаже")

		// Лифт спускается вниз
		for liftPosition > 0 {
			liftPosition--
			for i := 0; i < len(passengers); i++ {
				if passengers[i] == liftPosition {
					if passengersInLift < liftCapacity {
						passengersInLift++
						fmt.Printf("Пассажир забран с %d этажа\n", liftPosition)
						passengers[i] = -1 // Пассажир забран
					}
				}
			}
			if liftPosition == 1 {
				passengersDelivered += passengersInLift
				passengersInLift = 0
				fmt.Println("Пассажиры доставлены на первый этаж")
			}
		}
	}
	fmt.Println("Все пассажиры доставлены на первый этаж")
}
