package main

import (
	"flag"
	"fmt"
	"os"
)

type ToDo struct {
	ID        int
	Task      string
	Completed bool
}

var todos []ToDo

// add task
func addTask(task string) { //Defining function. Takes the string parameter task
	id := len(todos) + 1                                              //Generating a new ID for the task It's one more than the current number of tasks[incrementing ]
	todos = append(todos, ToDo{ID: id, Task: task, Completed: false}) //appending a new Todo struct to the todos slice with the generated ID , the task description , and a default completion
	fmt.Printf("Added \"%s\" to Todo list .\n", task)                 //printing a confirmation message indicating that the task was added
}

func listTasks() {
	for _, t := range todos {
		status := "Pending"
		if t.Completed {
			status = "Completed"
		}
		fmt.Printf("%d. %s [%s]\n", t.ID, t.Task, status)
	}
}

func completeTask(id int) {
	for i, t := range todos {
		if t.ID == id {
			todos[i].Completed = true
			fmt.Printf("Marked \"%s\" as completed.\n", t.Task)
			return
		}
	}
	fmt.Println("Task not found.")
}

func deleteTask(id int) {
	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			fmt.Printf("Deleted \"%s\" from your ToDo list.\n", t.Task)
			return
		}
	}
	fmt.Println("Task not found.")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Expected 'add', 'list', 'complete' or 'delete' subcommands")
		os.Exit(1)
	}

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addTaskArg := addCmd.String("task", "", "Task to add")

	completeCmd := flag.NewFlagSet("complete", flag.ExitOnError)
	completeTaskArg := completeCmd.Int("id", 0, "Task ID to mark as completed")

	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteTaskArg := deleteCmd.Int("id", 0, "Task ID to delete")

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		if *addTaskArg == "" {
			fmt.Println("Please provide a task to add.")
			os.Exit(1)
		}
		addTask(*addTaskArg)

	case "list":
		listTasks()

	case "complete":
		completeCmd.Parse(os.Args[2:])
		if *completeTaskArg == 0 {
			fmt.Println("Please provide a valid task ID to complete.")
			os.Exit(1)
		}
		completeTask(*completeTaskArg)

	case "delete":
		deleteCmd.Parse(os.Args[2:])
		if *deleteTaskArg == 0 {
			fmt.Println("Please provide a valid task ID to delete.")
			os.Exit(1)
		}
		deleteTask(*deleteTaskArg)

	default:
		fmt.Println("Expected 'add', 'list', 'complete' or 'delete' subcommands")
		os.Exit(1)
	}
}
