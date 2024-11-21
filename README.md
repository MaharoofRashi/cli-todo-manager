
# To-Do List Manager (CLI)

A simple command-line interface (CLI) application built in Go to manage your to-do list. This tool allows you to add, view, delete, and mark tasks as completed. Tasks are stored persistently in a JSON file.

---

## Features

1. **Add a New Task**:
   Add a new to-do item with a description.
   ```bash
   go run main.go -add "Buy groceries"
   ```

2. **View All Tasks**:
   Display all tasks, showing their status (completed or incomplete).
   ```bash
   go run main.go -view
   ```

3. **Mark a Task as Completed**:
   Mark a specific task as completed using its ID.
   ```bash
   go run main.go -complete <id>
   ```

4. **Delete a Task**:
   Delete a task from the list using its ID.
   ```bash
   go run main.go -delete <id>
   ```

5. **Persistent Storage**:
   Tasks are saved in a `todo.json` file, ensuring your data remains available between program runs.

---

## Usage

### 1. Running the Program
Ensure you have Go installed on your system. Clone this repository and run the program as follows:
```bash
go run main.go <flags>
```

### 2. Available Commands
| Flag            | Description                                    | Example Usage                         |
|------------------|------------------------------------------------|---------------------------------------|
| `-add`          | Add a new task with a description              | `go run main.go -add "Learn Go"`      |
| `-view`         | View all tasks                                 | `go run main.go -view`                |
| `-delete <id>`  | Delete a task by ID                            | `go run main.go -delete 2`            |
| `-complete <id>`| Mark a task as completed by ID                 | `go run main.go -complete 1`          |

---

## Example Workflow

1. Add Tasks:
   ```bash
   go run main.go -add "Learn Go"
   go run main.go -add "Build a To-Do List Manager"
   ```

2. View Tasks:
   ```bash
   go run main.go -view
   ```

   Output:
   ```plaintext
   Your tasks:
   [ ] 1: Learn Go
   [ ] 2: Build a To-Do List Manager
   ```

3. Mark Task 1 as Completed:
   ```bash
   go run main.go -complete 1
   ```

4. View Tasks Again:
   ```bash
   go run main.go -view
   ```

   Output:
   ```plaintext
   Your tasks:
   [✔] 1: Learn Go
   [ ] 2: Build a To-Do List Manager
   ```

5. Delete Task 2:
   ```bash
   go run main.go -delete 2
   ```

---

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/todo-list-manager
   cd todo-list-manager
   ```

2. Build the project:
   ```bash
   go build -o todo
   ```

3. Run the program:
   ```bash
   ./todo <flags>
   ```

---

## Project Structure

```plaintext
todo-list-manager/
├── main.go          # Main entry point for the CLI application
├── todo.json        # JSON file to persist tasks (created on first run)
└── README.md        # Project documentation
```

---

## Future Features

1. **Task Sorting**:
   - Allow tasks to be sorted by ID or completion status.

2. **Search Functionality**:
   - Add the ability to search for tasks by keywords (e.g., `-search "keyword"`).

3. **Export Tasks**:
   - Add a flag to export tasks to other formats like CSV or plain text (e.g., `-export <filename>`).

4. **Undo Last Action**:
   - Implement a feature to undo the last modification (e.g., delete, complete).

5. **Recurring Tasks**:
   - Support for recurring tasks with configurable intervals.

6. **Enhanced CLI Output**:
   - Use colored output to improve readability (e.g., `github.com/fatih/color`).

7. **Unit Tests**:
   - Add unit tests to validate functionality and improve code reliability.

---

## Contributing

Contributions are welcome! If you'd like to add features or fix bugs:
1. Fork the repository.
2. Create a feature branch.
3. Submit a pull request with detailed information about your changes.

---

## License

This project is licensed under the [MIT License](LICENSE).
