# To-Do List Application (Go)

A simple Go-based to-do list program that demonstrates how to manage tasks in memory: **view tasks**, **add a new task**, and **mark a task as completed**.

---

## Features

-  Display an initial list of tasks
-  Add a new task to the list
-  Mark a task as completed using its index
-  Show task due dates (formatted as `YYYY-MM-DD`)


---

## How It Works

This program defines a `task` struct:

- `name` (string): The task title
- `dueDate` (time.Time): The deadline
- `isCompleted` (bool): Completion status


## Project Structure
to-do-list/
└── main.go
└── README.md
└──go.mod


## Requirements
- `Go installed on your PC (version 1.18+ recommended)`
### Check your Go version:
```bash
go version
```

### Run the program
```bash
go run main.go
```


