package main

import (
	"fmt"
)

func main() {
	var n, k, studentNum int
	fmt.Println("Введите количество студентов (N):")
	fmt.Scan(&n)
	fmt.Println("Введите количество групп (K):")
	fmt.Scan(&k)
	fmt.Println("Введите порядковый номер студента:")
	fmt.Scan(&studentNum)
	
	group := (studentNum - 1) % k + 1
	fmt.Printf("Студент с номером %d попадает в группу %d.\n", studentNum, group)
}
