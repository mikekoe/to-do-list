package main

import (
	"fmt"
	"time"
	"bufio"
	"os"
	"strings"
)

type task struct{
	name string
	dueDate time.Time
	isCompleted bool
}

func showOptions(){
	fmt.Print("Choose an option: 1. View Tasks 2. Add Task 3. Mark Task as Completed 4. Delete Task 5. Exit\n")
}

func main() {
   taskList := []*task{}

   fmt.Print("Enter your name: ")
   reader := bufio.NewReader(os.Stdin)

   inputedName, err := reader.ReadString('\n')
   if err != nil {
	   fmt.Println("Error reading input:", err)
	   return
   }
   formattedName := strings.TrimSpace(inputedName)

   fmt.Printf("Hello, %v Welcome to the To-do List Application! What do you want to do now?\n", formattedName)


    for  {
	   showOptions()
		inputOption, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		formattedOption := strings.TrimSpace(inputOption)
		fmt.Printf("You selected option: %v\n", formattedOption)

		if formattedOption == "exit" || formattedOption == "5" {
			fmt.Println("Goodbye!")
			break
		}

	   switch formattedOption {
			case "1":
				viewTasks(taskList)
				
			case "2":
				fmt.Print("Enter task name: ")
				taskName, _ := reader.ReadString('\n')
				taskName = strings.TrimSpace(taskName)

				fmt.Print("Enter due date (YYYY-MM-DD): ")
				dueDateStr, _ := reader.ReadString('\n')
				dueDateStr = strings.TrimSpace(dueDateStr)
				dueDate, err := time.Parse("2006-01-02", dueDateStr)
				if err != nil {
					fmt.Println("Invalid date format:", err)
					break
				}
				taskList = addTask(taskList, taskName, dueDate)
				viewTasks(taskList)
			case "3":
				fmt.Print("Enter task name to be marked as completed: ")
				markIndexInput, _ := reader.ReadString('\n')
				markIndex := strings.TrimSpace(markIndexInput)
				markTaskCompleted(taskList, markIndex)
				viewTasks(taskList)
			case "4":
				fmt.Print("Enter task name to be deleted: ")
				
				delIndexInput, _ := reader.ReadString('\n')
				delIndex := strings.TrimSpace(delIndexInput)
				taskList = deleteTask(taskList, delIndex)
				viewTasks(taskList)
			default:
				fmt.Println("Invalid option. Please try again.")
		}
	}
}

func addTask(taskList []*task, name string, dueDate time.Time) []*task {
	newTask := &task{name: name, dueDate: dueDate, isCompleted: false}
	fmt.Printf("The task - %s - has been added\n", newTask.name)
	return append(taskList, newTask)
}

func markTaskCompleted(taskList []*task, name string) {
	for i, t := range taskList {
		if t.name == name {
			taskList[i].isCompleted = true
			return
		}
	}
	fmt.Println("Task not found")
}

func viewTasks(taskList []*task) {
	if len(taskList) <= 0 {
		fmt.Println("No tasks available.")
		return
	}
	fmt.Println("######## Your Tasks ########")
	
	for _, t := range taskList {
		status := "Pending"
		if t.isCompleted {
			status = "Completed"
		}
		
		fmt.Printf("%s - Due: %s - Status: %s\n", t.name, t.dueDate.Format("2006-01-02"), status)
	}
}

func deleteTask (taskList []*task, name string) []*task {
	
	for i, t := range taskList {
		if t.name == name {
			fmt.Println("You deleted the task:", t.name)
			return append(taskList[:i], taskList[i+1:]...)
		}
	}
	fmt.Println("Task not found")
	return taskList
}



