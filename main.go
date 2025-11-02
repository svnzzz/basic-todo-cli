package main

import (
	"fmt"
)

func main() {
	todos := &Todos{}

	for {
		fmt.Println("\n=== TODO LIST ===")
		fmt.Println("1. Add task")
		fmt.Println("2. Delete task")
		fmt.Println("3. Complete task")
		fmt.Println("4. Show all tasks")
		fmt.Println("5. Exit")

		var scelta int
		fmt.Print("Choose an option: ")
		fmt.Scan(&scelta)

		switch scelta {
		case 1:
			var title string
			fmt.Print("Enter task title: ")
			fmt.Scan(&title)
			todos.Add(title)

		case 2:
			var index int
			fmt.Print("Enter the index of the task to delete: ")
			fmt.Scan(&index)
			todos.Delete(index)

		case 3:
			var index int
			fmt.Print("Enter the index of the task to complete: ")
			fmt.Scan(&index)
			todos.Complete(index)

		case 4:
			todos.List()

		case 5:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid option.")

		}
	}
}
