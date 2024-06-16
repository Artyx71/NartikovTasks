package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("log.txt")
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Ошибка получения информации о файле:", err)
		return
	}

	if fileInfo.Size() == 0 {
		fmt.Println("Файл пуст")
		return
	}

	data := make([]byte, fileInfo.Size())
	_, err = file.Read(data)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	fmt.Println(string(data))
}
