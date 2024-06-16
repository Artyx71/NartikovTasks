package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("readonly.txt")
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)
		return
	}
	defer file.Close()

	err = file.Chmod(0400) // Установка режима только для чтения
	if err != nil {
		fmt.Println("Ошибка установки режима файла:", err)
		return
	}

	file, err = os.Open("readonly.txt")
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("Эта строка не должна быть записана")
	if err != nil {
		fmt.Println("Ошибка записи в файл:", err)
	} else {
		fmt.Println("Данные успешно записаны в файл")
	}
}
