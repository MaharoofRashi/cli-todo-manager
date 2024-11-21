package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type Task struct {
	ID        int    `json:"id"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

var (
	add      = flag.String("add", "", "Add a new task")
	view     = flag.Bool("view", false, "View all tasks")
	del      = flag.Int("delete", 0, "Delete a task by ID")
	complete = flag.Int("complete", 0, "Mask a task as completed by ID")
)

const fileName = "todo.json"

func WriteTasks(tasks []Task) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(tasks)
}

func ReadTasks() ([]Task, error) {
	file, err := os.Open(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var tasks []Task
	err = json.NewDecoder(file).Decode(&tasks)
	return tasks, err
}

func addTask(task string) {
	tasks, err := ReadTasks()
	if err != nil {
		fmt.Println("Error reading tasks: ", err)
		return
	}

	newID := 1
	if len(tasks) > 0 {
		newID = tasks[len(tasks)-1].ID + 1
	}

	newTask := Task{
		ID:        newID,
		Task:      task,
		Completed: false,
	}
	tasks = append(tasks, newTask)

	err = WriteTasks(tasks)
	if err != nil {
		fmt.Println("Error writing tasks: ", err)
		return
	}

	fmt.Printf("Added task: %s.\n", task)
}

func viewTasks() {
	tasks, err := ReadTasks()
	if err != nil {
		fmt.Println("Error reading tasks: ", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	fmt.Println("Your tasks: ")
	for _, task := range tasks {
		status := " "
		if task.Completed {
			status = "✔"
		}
		fmt.Printf("[%s] %d: %s\n", status, task.ID, task.Task)
	}
}

func deleteTask(id int) {
	tasks, err := ReadTasks()
	if err != nil {
		fmt.Println("Error in reading tasks: ", err)
		return
	}

	found := false
	newTasks := []Task{}
	for _, task := range tasks {
		if task.ID == id {
			found = true
		} else {
			newTasks = append(newTasks, task)
		}
	}

	if !found {
		fmt.Printf("Task with ID: %d, not found.\n", id)
		return
	}

	err = WriteTasks(newTasks)
	if err != nil {
		fmt.Println("Error writing tasks: ", err)
		return
	}
	fmt.Printf("Deleted task with ID: %d.\n", id)
}

func completeTask(id int) {
	tasks, err := ReadTasks()
	if err != nil {
		fmt.Println("Error reading tasks: ", err)
		return
	}

	found := false
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Completed = true
			found = true
			break
		}
	}
	if !found {
		fmt.Printf("Task with ID: %d, not exists.\n", id)
		return
	}

	err = WriteTasks(tasks)
	if err != nil {
		fmt.Println("Error in writing tasks: ", err)
		return
	}

	fmt.Printf("Marked task with ID: %d as completed.\n", id)
}

func main() {
	flag.Parse()

	if *add != "" {
		addTask(*add)
	} else if *view {
		viewTasks()
	} else if *del != 0 {
		deleteTask(*del)
	} else if *complete != 0 {
		completeTask(*complete)
	} else {
		fmt.Println("Please provide a valid command: -add, -view, -delete, -complete")
	}
}
