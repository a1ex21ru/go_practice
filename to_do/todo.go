package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"time"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}

	*todos = append(*todos, todo)
}

func (todos *Todos) delete(index int) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	*todos = remove(t, index)

	return nil
}

func (todos *Todos) edit(index int) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter title")

	title, _ := reader.ReadString('\n')

	t[index] = Todo{title, false, time.Now(), nil}

	return nil
}

func (todos *Todos) complete(index int) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	new := time.Now()

	t[index].Completed = true
	t[index].CompletedAt = &new

	return nil
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("Invalid index")
		fmt.Println(err)
		return err
	}
	return nil
}

func remove(slice Todos, index int) Todos {
	result := append(slice[:index], slice[index+1:]...)
	return result
}

func (todos *Todos) print() {

	fmt.Printf("%-4s | %-20s | %10s | %15s | %20s |\n", "id", "Title", "Completed", "Created time", "Completed time")
	for index, val := range *todos {
		completed := "✅"
		uncompleted := "❌"
		createdStr := val.CreatedAt.Format("2006-01-02 15:04")

		var completedStr string
		if val.CompletedAt == nil {
			completedStr = "---"
		} else {
			completedStr = val.CompletedAt.Format("2006-01-02 15:04")
		}
		if val.Completed {
			printLine(index, val.Title, completed, createdStr, completedStr)
		} else {
			printLine(index, val.Title, uncompleted, createdStr, completedStr)
		}
	}
	fmt.Println()
}

func printLine(index int, title string, compl string, createdTime string, completedTime string) {

	fmt.Printf("%-4d | %-20s | %8v | %15s | %20s |\n", index, title, compl, createdTime, completedTime)
}
