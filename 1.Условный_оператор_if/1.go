package main

import (
	"fmt"
)

func main() {
	var score1, score2, score3 int
	fmt.Println("Введите результат первого экзамена:")
	fmt.Scan(&score1)
	fmt.Println("Введите результат второго экзамена:")
	fmt.Scan(&score2)
	fmt.Println("Введите результат третьего экзамена:")
	fmt.Scan(&score3)
	
	totalScore := score1 + score2 + score3
	fmt.Println("Сумма проходных баллов: 275")
	fmt.Printf("Количество набранных баллов: %d\n", totalScore)
	
	if totalScore >= 275 {
		fmt.Println("Вы поступили.")
	} else {
		fmt.Println("Вы не поступили.")
	}
}
