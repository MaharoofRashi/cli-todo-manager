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
	add  = flag.String("add", "", "Add a new task")
	view = flag.Bool("view", false, "View all tasks")
	del  = flag.Int("delete", 0, "Delete a task by ID")
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

func main() {
	flag.Parse()

	if *add != "" {
		addTask(*add)
	} else if *view {
		viewTasks()
	} else if *del != 0 {
		deleteTask(*del)
	} else {
		fmt.Println("Please provide a valid command: -add, -view, -delete")
	}
}
