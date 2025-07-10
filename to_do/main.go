package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	todos := Todos{}

	menu(&todos)

}

func menu(todos *Todos) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Hello!")

	goodbye := false
	for {
		fmt.Print("1. Add task\t")
		fmt.Print("2. Edit task\t\n")
		fmt.Print("3. Delete task\t")
		fmt.Println("4. Complete task")
		fmt.Print("5. Output todolist\t")
		fmt.Println("0. Exit")

		input, _ := reader.ReadString('\n')
		choise := strings.TrimSpace(input)

		switch choise {

		case "1":
			fmt.Println("Enter title fo task: ")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)
			todos.add(title)
		case "2":
			if !check_empty(todos) {
				break
			}
			fmt.Println("Enter id for edit task:")
			input, _ := reader.ReadString('\n')
			index, _ := strconv.Atoi(strings.TrimSpace(input))

			todos.edit(index)

		case "3":
			if !check_empty(todos) {
				break
			}
			fmt.Println("Enter id for delete task")
			input, _ := reader.ReadString('\n')
			index, _ := strconv.Atoi(strings.TrimSpace(input))

			todos.delete(index)

		case "4":
			if !check_empty(todos) {
				break
			}
			todos.print()
			fmt.Println("Enter id for completed task:")
			input, _ := reader.ReadString('\n')
			index, _ := strconv.Atoi(strings.TrimSpace(input))

			todos.complete(index)

		case "5":
			todos.print()
		case "0":
			fmt.Println("Exit, goodbye!")
			goodbye = true
		default:
			fmt.Println("Unknown number menu\nPlease enter number again!")
		}

		if goodbye {
			break
		}
	}
}

func check_empty(todos *Todos) bool {
	if len(*todos) == 0 {
		fmt.Println("Todolist is empty!")
		return false
	} else {
		return true
	}
}
