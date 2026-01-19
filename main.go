package main

import (
	"fmt"
	"time"
)

type task struct{
	name string
	dueDate time.Time
	isCompleted bool
	
}


func main() {
   taskList := []*task{
		{"Read books", time.Date(2026, time.January, 10, 11, 0, 0, 0, time.UTC), false},
		{"Finish leetcode questions", time.Date(2026, time.January, 15, 15, 0, 0, 0, time.UTC), false},
		{"Write code", time.Date(2026, time.January, 12, 7, 0, 0, 0, time.UTC), false},
   }

   fmt.Println("Welcome to the To-do List Application! What do you want to do now?")
   for _, t := range taskList {
	    fmt.Printf("%s Task:, Due: %s, is-Completed: %t\n", t.name, t.dueDate, t.isCompleted)
   }
   addTasks := addTask(taskList, "Go to Gym", time.Date(2026, time.January, 20, 3, 30, 0, 0, time.UTC))
   markTaskCompleted(addTasks, 2)
   fmt.Println("\nUpdated Task List:")
   viewTasks(addTasks)
}

func addTask(taskList []*task, name string, dueDate time.Time) []*task {
	newTask := &task{name: name, dueDate: dueDate, isCompleted: false}
	return append(taskList, newTask)
}

func markTaskCompleted(taskList []*task, index int) {
	if index >= 0 && index < len(taskList) {
		taskList[index].isCompleted = true
	} else {
		fmt.Println("Invalid task index")
	}
}

func viewTasks(taskList []*task) {
	for _, t := range taskList {
		status := "Pending"
		if t.isCompleted {
			status = "Completed"
		}
		fmt.Printf("%s - Due: %s - Status: %s\n", t.name, t.dueDate.Format("2006-01-02"), status)
	}
}


