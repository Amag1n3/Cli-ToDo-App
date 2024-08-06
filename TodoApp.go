package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "This command runs the CLI ToDo list app",
	Long:  "This command has 3 functionalities which are passed as arguemtns to the command.(add, list, edit tasks)",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to my Command Line based ToDo List App made in Golang")
	},
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Use this command to add tasks",
	Long:  "This command is used to add tasks to the list of your ToDo tasks",
	Run: func(cmd *cobra.Command, args []string) {
		userHome, _ := os.UserHomeDir()
		filepath := userHome + "/Desktop/CliTodoApp/todolist.txt"
		task := writeTask()
		file, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		defer file.Close()
		file.WriteString(task)
		fmt.Println("Task added successfully")
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Use this command to list all the tasks",
	Long:  "This command is used to list all the tasks currently in your ToDo list",
	Run: func(cmd *cobra.Command, args []string) {
		listTasks()
	},
}

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Use this command to edit any task",
	Long:  "This command can be used to edit any task as long as you know it's name. (Case Sensitive)",
	Run: func(cmd *cobra.Command, args []string) {
		editTask()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(editCmd)
}

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

	creationTime := time.Now().Format("02/01/2006 1504")

	task := fmt.Sprintf("%s,%s,%s,%s,%s\n", taskName, taskDate, taskTime, taskStatus, creationTime)
	return task
}

func listTasks() {
	var choice int
	userHome, _ := os.UserHomeDir()
	filepath := userHome + "/Desktop/CliTodoApp/todolist.txt"
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
	now := time.Now()
	var timeLeft string
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
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}
			fields := strings.Split(line, ",")
			if len(fields) != 5 {
				continue
			}
			taskName := fields[0]
			taskDate := fields[1]
			taskTime := fields[2]
			taskStatus := fields[3]
			//creationTime := fields[4]

			deadlineStr := taskDate + " " + taskTime
			deadlineTime, err := time.Parse("02/01/2006 1504", deadlineStr)
			if err != nil {
				fmt.Println("error parsing time:", err)
			}
			duration := deadlineTime.Sub(now)
			if duration < 0 {
				timeLeft = "Task overdue"
			} else {
				days := int(duration.Hours()) / 24
				hours := int(duration.Hours()) % 24
				minutes := int(duration.Minutes()) % 60
				timeLeft = fmt.Sprintf("Time left to complete: %d days, %d hours, %d minutes", days, hours, minutes)
			}
			fmt.Printf("Name: %s, Date: %s, Time: %s, Status: %s, %s\n", taskName, taskDate, taskTime, taskStatus, timeLeft)

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
			if len(fields) != 5 {
				fmt.Printf("Error: Invalid task format or missing fields in line: %s", line)
				continue
			}
			taskName := fields[0]
			taskDate := fields[1]
			taskTime := fields[2]
			taskStatus := fields[3]

			if date == taskDate {
				deadlineStr := taskDate + " " + taskTime
				deadlineTime, _ := time.Parse("02/01/2006 1504", deadlineStr)
				duration := deadlineTime.Sub(now)
				timeLeft := "Time left to complete: " + duration.String()
				if duration < 0 {
					timeLeft = "Task Overdue"
				} else {
					days := int(duration.Hours()) / 24
					hours := int(duration.Hours()) % 24
					minutes := int(duration.Minutes()) % 60
					timeLeft = fmt.Sprintf("Time left to complete: %d days, %d hours, %d minutes", days, hours, minutes)
				}
				fmt.Printf("Name: %s,Date: %s,Time: %s,Status: %s, %s\n", taskName, taskDate, taskTime, taskStatus, timeLeft)
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
				deadlineStr := taskDate + " " + taskTime
				deadlineTime, err := time.Parse("02/01/2006 1504", deadlineStr)
				if err != nil {
					fmt.Println("error parsing time:", err)
				}
				duration := deadlineTime.Sub(now)
				timeLeft := "Time left to complete: " + duration.String()
				if duration < 0 {
					timeLeft = "Task overdue"
				} else {
					days := int(duration.Hours()) / 24
					hours := int(duration.Hours()) % 24
					minutes := int(duration.Minutes()) % 60
					timeLeft = fmt.Sprintf("Time left to complete: %d days, %d hours, %d minutes", days, hours, minutes)
				}
				fmt.Printf("Name: %s, Date: %s, Time: %s, Status: %s, %s\n", taskName, taskDate, taskTime, taskStatus, timeLeft)
			}
		}
	}
}

func editTask() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter name of task to edit")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	userHome, _ := os.UserHomeDir()
	filepath := userHome + "/Desktop/CliTodoApp/todolist.txt"
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return // or handle the error appropriately
	}
	defer file.Close()
	var tasks []string
	filereader := bufio.NewReader(file)
	found := false
	anyEdited := false
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
			fmt.Println("Enter field to edit: 1. Name, 2. Status, 3. Time, 4. Date")
			fmt.Scanln(&num)
			if num == 1 {
				fmt.Println("Enter new name for the task: ")
				newName, _ := reader.ReadString('\n')
				newName = strings.TrimSpace(newName)
				if newName == "" {
					newName = fields[0]
				}
				fields[0] = newName
				anyEdited = true
			} else if num == 2 {
				fmt.Println("Enter new status for the task: ")
				newStatus, _ := reader.ReadString('\n')
				newStatus = strings.TrimSpace(newStatus)
				fields[3] = newStatus
				anyEdited = true
			} else if num == 3 {
				fmt.Println("Enter new deadline time in HHMM(24 HOURS) format.")
				newTime, _ := reader.ReadString('\n')
				newTime = strings.TrimSpace(newTime)
				fields[2] = newTime
				anyEdited = true
			} else if num == 4 {
				fmt.Println("Enter new date in DD/MM/YYYY format")
				newDate, _ := reader.ReadString('\n')
				newDate = strings.TrimSpace(newDate)
				fields[1] = newDate
				anyEdited = true
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

	if !anyEdited {
		fmt.Println("No Tasks edited")
		return
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
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}
