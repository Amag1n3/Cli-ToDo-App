package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func writeTask() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Name of the task: ")
	taskName, _ := reader.ReadString('\n')
	taskName = strings.TrimSpace(taskName)

	fmt.Println("Enter deadline date for the task: ")
	taskDate, _ := reader.ReadString('\n')
	taskDate = strings.TrimSpace(taskDate)

	fmt.Println("Enter deadline time for the task: ")
	taskTime, _ := reader.ReadString('\n')
	taskTime = strings.TrimSpace(taskTime)

	fmt.Println("Enter status of the task: ")
	taskStatus, _ := reader.ReadString('\n')
	taskStatus = strings.TrimSpace(taskStatus)
	if taskStatus == "" {
		taskStatus = "Ongoing"
	}

	task := fmt.Sprintf("%s,%s,%s,%s\n", taskName, taskDate, taskTime, taskStatus)
	return task
}

func listTasks() {
	var choice int
	filepath := "/home/amogh/Desktop/Cli_ToDo_List/todo_list.txt"
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return // or handle the error appropriately
	}
	defer file.Close()
	inputreader := bufio.NewReader(os.Stdin)
	reader := bufio.NewReader(file)
	fmt.Println("How would you like to list the tasks?(1: All, 2: Date, 3: Status)")
	fmt.Scanln(&choice)

	if choice == 1 {
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break
				}
				fmt.Println("Error:", err)
				return
			}
			fmt.Println(line) // Print the line
		}
	}
	if choice == 2 {
		fmt.Println("Enter date in DD/MM/YYYY format")
		date, _ := inputreader.ReadString('\n')
		date = strings.TrimSpace(date)
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break
				}
				fmt.Println("Error:", err)
				return
			}
			line = strings.TrimSpace(line)

			if line == "" {
				continue
			}
			fields := strings.Split(line, ",")
			//if len(fields) != 4 {
			//  fmt.Printf("Error: Invalid task format or missing fields in line: %s", line)
			//  continue
			//}
			taskName := fields[0]
			taskDate := fields[1]
			taskTime := fields[2]
			taskStatus := fields[3]

			if date == taskDate {
				fmt.Printf("Name: %s,Date: %s,Time: %s,Status: %s\n", taskName, taskDate, taskTime, taskStatus)
			}

		}
	}
	if choice == 3 {
		fmt.Println("Enter Ongoing, Completed or Incomplete")
		status, _ := inputreader.ReadString('\n')
		status = strings.TrimSpace(status)
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break
				}
				fmt.Println("Error:", err)
				return
			}
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}
			fields := strings.Split(line, ",")
			//if len(fields) != 4 {
			//  fmt.Printf("Error: Invalid task format or missing fields in line: %s", line)
			// continue
			//}
			taskName := fields[0]
			taskDate := fields[1]
			taskTime := fields[2]
			taskStatus := fields[3]
			if status == taskStatus {
				fmt.Printf("Name: %s,Date: %s,Time: %s,Status: %s\n", taskName, taskDate, taskTime, taskStatus)
			}
		}
	}
}

func editTask() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter name of task to edit")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	filepath := "/home/amogh/Desktop/Cli_ToDo_List/todo_list.txt"
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return // or handle the error appropriately
	}
	defer file.Close()
	var tasks []string
	filereader := bufio.NewReader(file)
	found := false
	for {
		line, err := filereader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("error:", err)
			return

		}
		line = strings.TrimSpace(line)
		fields := strings.Split(line, ",")
		if name == fields[0] {
			found = true
			var num int
			fmt.Println("Enter field to edit: 1. Name, 2. Status")
			fmt.Scanln(&num)
			if num == 1 {
				fmt.Println("Enter new name for the task: ")
				newName, _ := reader.ReadString('\n')
				newName = strings.TrimSpace(newName)
				if newName == "" {
					newName = fields[0]
				}
				fields[0] = newName

			} else if num == 2 {
				fmt.Println("Enter new status for the task: ")
				newStatus, _ := reader.ReadString('\n')
				newStatus = strings.TrimSpace(newStatus)
				if newStatus == "" {
					newStatus = fields[3]
				}
				fields[3] = newStatus
				editedTask := fmt.Sprintf("%s,%s,%s,%s", fields[0], fields[1], fields[2], fields[3])
				tasks = append(tasks, editedTask)
			} else {
				tasks = append(tasks, line)
				continue
			}

			editedTask := strings.Join(fields, ",")
			tasks = append(tasks, editedTask)
		} else {
			tasks = append(tasks, line)
		}

	}
	if !found {
		fmt.Println("Task name entered wrong or doesn't exist")
	}
	file, err = os.OpenFile(filepath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("error:", err)
	}
	defer file.Close()

	// Write all tasks back to the file
	for _, task := range tasks {
		_, err := file.WriteString(task + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	fmt.Println("Task edited successfully.")

}

func main() {
	filepath := "/home/amogh/Desktop/Cli_ToDo_List/todo_list.txt"
	var input int
	fmt.Println("Enter your choice: 1.Add task, 2. List task(s), 3. Edit Tasks")
	fmt.Scanln(&input)

	if input == 1 {
		task := writeTask()
		file, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		defer file.Close()
		file.WriteString(task)
		fmt.Println("Task added successfully")
	}
	if input == 2 {
		listTasks()
	}
	if input == 3 {
		editTask()
	}
}
