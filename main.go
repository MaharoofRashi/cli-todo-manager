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
	tasks := []Task{
		{ID: 1, Task: "Buy groceries", Completed: false},
		{ID: 2, Task: "Clean the house", Completed: true},
	}

	err := WriteTasks(tasks)
	if err != nil {
		fmt.Println("Error writing tasks: ", err)
		return
	}

	readTasks, err := ReadTasks()
	if err != nil {
		fmt.Println("Error in reading tasks: ", err)
	}

	fmt.Println("Tasks from file: ", readTasks)
}
