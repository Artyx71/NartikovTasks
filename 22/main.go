// main.go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"

    "22/pkg/storage"
    "22/pkg/student"
)

func main() {
    store := storage.NewStorage()
    
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Println("Введите данные студентов в формате: name age grade")
    
    for scanner.Scan() {
        line := scanner.Text()
        if line == "" {
            continue
        }

        parts := strings.Fields(line)
        if len(parts) != 3 {
            fmt.Println("Неверный формат ввода. Попробуйте еще раз.")
            continue
        }

        name := parts[0]
        var age, grade int
        _, err1 := fmt.Sscanf(parts[1], "%d", &age)
        _, err2 := fmt.Sscanf(parts[2], "%d", &grade)
        if err1 != nil || err2 != nil {
            fmt.Println("Ошибка чтения возраста или оценки. Попробуйте еще раз.")
            continue
        }

        student := student.NewStudent(name, age, grade)
        store.Put(student)
    }

    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "Ошибка чтения ввода:", err)
    }

    fmt.Println("Список студентов:")
    for _, name := range store.GetAllNames() {
        fmt.Println(name)
    }
}