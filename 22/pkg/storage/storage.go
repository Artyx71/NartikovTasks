package storage

import (
    "22/pkg/student"
)

type Storage struct {
	students map[string]*student.Student
}

// NewStorage создаёт новое хранилище
func NewStorage() *Storage {
    return &Storage{students: make(map[string]*student.Student)}
}

func (s *Storage) Put(student *student.Student) {
    s.students[student.Name] = student
}

// GetAllNames возвращает имена всех студентов в хранилище
func (s *Storage) GetAllNames() []string {
    var names []string
    for name := range s.students {
        names = append(names, name)
    }
    return names
}