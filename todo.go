package main

import (
	"fmt"
	"time"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) Add(title string) {
	todo := Todo{
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	}
	*todos = append(*todos, todo)
}

func (todos *Todos) Delete(index int) error {
	t := *todos
	if index < 0 || index >= len(t) {
		return fmt.Errorf("index not found")
	}
	*todos = append(t[:index], t[index+1:]...)
	return nil
}

func (todos *Todos) Complete(index int) error {
	t := *todos
	if index < 0 || index >= len(t) {
		return fmt.Errorf("index not found")
	}
	completedTime := time.Now()
	t[index].Completed = !t[index].Completed
	t[index].CompletedAt = &completedTime
	return nil
}

func (todos Todos) List() {
	for _, todo := range todos {
		fmt.Printf("%s (Completed: %t)\n", todo.Title, todo.Completed)
	}
}
