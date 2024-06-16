package main

import (
	"fmt"
	"os"
	"time"
	"bufio"
)

func main() {
	file, err := os.Create("log.txt")
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)
		return
	}
	defer file.Close()

	fmt.Println("Введите сообщения для записи в файл (для завершения введите 'exit'):")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if input == "exit" {
			break
		}
		now := time.Now().Format("2006-01-02 15:04:05")
		_, err := file.WriteString(fmt.Sprintf("%s %s\n", now, input))
		if err != nil {
			fmt.Println("Ошибка записи в файл:", err)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка чтения ввода:", err)
	}
}
