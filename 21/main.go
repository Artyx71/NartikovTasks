package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

type Student struct {
    name  string
    age   int
    grade int
}

func newStudent(name string, age int, grade int) *Student {
    return &Student{name: name, age: age, grade: grade}
}

func put(storage map[string]*Student, student *Student) {
    storage[student.name] = student
}

func get(storage map[string]*Student, name string) *Student {
    return storage[name]
}

func main() {
    storage := make(map[string]*Student)
    
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

        student := newStudent(name, age, grade)
        put(storage, student)
    }

    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "Ошибка чтения ввода:", err)
    }

    fmt.Println("Список студентов из хранилища:")
    for _, student := range storage {
        fmt.Println(student.name, student.age, student.grade)
    }
}