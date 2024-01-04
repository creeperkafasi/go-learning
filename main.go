package main

import (
	"fmt"
	"bufio"
	"os"
)

type Task struct {
	Completed bool
	Text      string
}

func main() {
	var tasks []Task
	running := true
	scanner := bufio.NewScanner(os.Stdin)

	for running {
		fmt.Println("------- Tasks -------")
		if len(tasks) == 0 {
			fmt.Print("You don't have any tasks, try to add some!")
		} else {
			for i := 0; i < len(tasks); i++ {
				completed := ' '
				if tasks[i].Completed {
					completed = 'X'
				}
				fmt.Printf(
					"%d# [%c] %s\n",
					i+1,
					completed,
					tasks[i].Text,
				)
			}
		}

		var action byte
		fmt.Print("\n(A)dd, (D)elete, (E)dit, (C)ompleted, (Q)uit: ")
		_, err := fmt.Scanf("%c\n", &action)
		fmt.Println()
		if err != nil {
			fmt.Println("Error: ", err)
		}
		switch action {
		case 'A' | 'a':
			fmt.Print("Task Text: ")
			var taskText string
			if scanner.Scan(){
				taskText = scanner.Text()
			}
			tasks = append(tasks, Task{Completed: false, Text: taskText})
		case 'D' | 'd':
			fmt.Print("Enter the index of the task you want to delete: ")
			var index int
			_, err = fmt.Scan(&index)
			if err != nil {
				fmt.Print("Error", err)
			}
			if index > len(tasks) {
				fmt.Println("Please enter a valid index!")
				continue
			}
			index--
			tasks = append(tasks[:index], tasks[index+1:]...)
		case 'E' | 'e':
			fmt.Print("Enter the index of the task you want to edit: ")
			var index int
			_, err = fmt.Scan(&index)
			if err != nil {
				fmt.Print("Error", err)
			}
			if index > len(tasks) {
				fmt.Println("Please enter a valid index!")
				continue
			}
			index--

			fmt.Print("Text: ")
			var taskText string
			if scanner.Scan() {
				taskText = scanner.Text()
			}
			tasks =
				append(
					tasks[:index],
					append(
						[]Task{Task{Text: taskText, Completed: tasks[index].Completed}},
						tasks[index+1:]...)...,
				)
		case 'C' | 'c':
			var index int
			fmt.Print("Which task should be marked as completed? ")
			_, err := fmt.Scan(&index)
			if err != nil {
				fmt.Print("Error", err)
			}
			if index > len(tasks) {
				fmt.Println("Please enter a valid index!")
				continue
			}
			index--

			tasks =
				append(
					tasks[:index],
					append(
						[]Task{Task{Text: tasks[index].Text, Completed: true}},
						tasks[index+1:]...)...,
				)

		case 'Q' | 'q':
			fmt.Print("Are you sure you want to quit? (y/N)")
			var confirm rune
			fmt.Scanf("%c", &confirm)
			if confirm == 'Y' || confirm == 'y' {
				running = false
			}
		}

	}
}

