package main
import (
  "fmt"
  "os"
  "io"
  "bufio"
  "strings"
)

func writeTask() string{
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
  if taskStatus == ""{
    taskStatus = "Ongoing"
  }

  task := fmt.Sprintf("%s,%s,%s,%s\n", taskName, taskDate, taskTime, taskStatus)
  return task
}



func listTasks(){
  var choice int
  filepath := "/home/amogh/Desktop/Cli_ToDo_List/todo_list.txt"
  file, err := os.Open(filepath)
  if err != nil {
    fmt.Println("Error opening file:", err)
    return // or handle the error appropriately
  }
  defer file.Close()
  _, err = file.Seek(0, io.SeekStart) // Ensure we are at the start of the filepath
  if err != nil {
    fmt.Println("error seeking to the start of file:", err)
    return
  }
  inputreader := bufio.NewReader(os.Stdin)
  reader := bufio.NewReader(file)
  fmt.Println("How would you like to list the tasks?(1: All, 2: Date, 3: Status)")
  fmt.Scanln(&choice)

  if choice == 1{
    for{
      line,err := reader.ReadString('\n')
      if err!=nil{
        if err == io.EOF{
          break
        }
        fmt.Println("Error:",err)
        return
      }
      fmt.Println(line) // Print the line
    }
  }
  if choice == 2{
    fmt.Println("Enter date in DD/MM/YYYY format")
    date, _ := inputreader.ReadString('\n')
    date = strings.TrimSpace(date)
    for{
      line, err := reader.ReadString('\n')
      if err!=nil{
        if err == io.EOF{
          break
        }
        fmt.Println("Error:",err)
        return
      }
      line = strings.TrimSpace(line)
      if line == ""{
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

      if date == taskDate{
        fmt.Printf("Name: %s,Date: %s,Time: %s,Status: %s\n", taskName, taskDate, taskTime, taskStatus)
      }

    }
  }
  if choice == 3{
    fmt.Println("Enter Ongoing, Completed or Incomplete")
    status, _ := inputreader.ReadString('\n')
    status = strings.TrimSpace(status)
    for{
      line, err := reader.ReadString('\n')
      if err!=nil{
        if err == io.EOF{
          break
        }
        fmt.Println("Error:",err)
        return
      }
      line = strings.TrimSpace(line)
      if line == ""{
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
      if status == taskStatus{
        fmt.Printf("Name: %s,Date: %s,Time: %s,Status: %s\n", taskName, taskDate, taskTime, taskStatus)
      }
    }
  }



  
}




func main(){
  filepath := "/home/amogh/Desktop/Cli_ToDo_List/todo_list.txt"
  var input int
  fmt.Println("Enter your choice: 1.Add task, 2. List task(s)")
  fmt.Scanln(&input)

  if input == 1{
    task := writeTask()
    file, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
    if err!=nil{
      fmt.Println("error:", err)
      return
    }
    defer file.Close()
    _, err = file.Seek(0,0) // Ensure we are at the end of the file
    if err != nil {
      fmt.Println("error seeking to the end of file:", err)
      return
    }
    file.WriteString(task)
    fmt.Println("Task added successfully")
  }
  if input == 2{
    listTasks()
  }
}
