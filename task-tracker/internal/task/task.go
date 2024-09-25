package task

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)
type TaskStatus string

const (
  TASK_TODO         TaskStatus = "todo"
  TASK_IN_PROGRESS  TaskStatus = "in-progress"
  TASK_COMPLETE     TaskStatus = "done"
)

type Task struct {
  ID int    `json:"id"`
  Description string `json:"description"`
  Status TaskStatus `json:"status"`
  CreatedAt time.Time `json:"createdAt"`
  UpdatedAt time.Time `json:"UpdatedAt"`
}

type List []Task

func  NewTask(id int, description string) *Task  {
  return &Task{
    ID: id,
    Description: description,
    Status: TASK_TODO,
    CreatedAt: time.Now(),
    UpdatedAt: time.Now(),
  }
}

func (list *List) Add(task *Task) error {
  taskList := *list
  
  taskList = append(taskList, *task)

  data, err := json.Marshal(taskList)

  if err != nil {
    fmt.Println("Got error while marhaling tasks:\t", err)
  }

  err = WriteToFile(data)
  
  if err != nil {
    fmt.Println("Error while saving tasks to file:\t", err)
  }
  return nil
}

func (list *List) FilteredTask(status string) *List{
  tasksList := *list

  if(status == "all") {
    return &tasksList
  }

  var filteredList List 
  for _, task := range tasksList {
    if string(task.Status) == status {
      filteredList = append(filteredList, task)
    }
  }
  return &filteredList
}

func (list *List) DeleteTask(id int) bool{
  taskList := *list
  var mark bool
  for i, task := range taskList{
    mark = false
    if task.ID == id {
      taskList = append(taskList[:i], taskList[i+1:]...)
      mark = true
      break
    }
  }

  data, err := json.Marshal(taskList)
  
  if err != nil {
    panic(err)
  } 

  err = WriteToFile(data)
  if err != nil {
   fmt.Print("Content couldn't write to file:") 
   panic(err)
  }
  return mark
}


func (list *List) UpdateDescription(id int, description string) bool {
  taskList := *list
  mark := false
  for i, task := range taskList{
    mark = false
    if task.ID == id {
      taskList[i].Description = description
      taskList[i].UpdatedAt = time.Now()
      mark = true
      break
    }
  }

  data, err := json.Marshal(taskList)
  
  if err != nil {
    panic(err)
  } 

  err = WriteToFile(data)
  if err != nil {
   fmt.Print("Content couldn't write to file:") 
   panic(err)
  }
  return mark
}

func (list *List) UpdateStatus(id int, status string) bool {
  mark := false
  taskList := *list
  for i, task := range taskList{
    mark = false
    if task.ID == id {
      taskList[i].Status = TaskStatus(status)
      taskList[i].UpdatedAt = time.Now()
      mark = true
      break
    }
  }

  data, err := json.Marshal(taskList)
  
  if err != nil {
    panic(err)
  } 

  err = WriteToFile(data)
  if err != nil {
   fmt.Print("Content couldn't write to file:") 
   panic(err)
  }
  return mark
}

func (list *List) FormatTasks(status string) string {
	var output strings.Builder

	for _, task := range *list {
		// Determine the task's status
		taskStatus := string(task.Status)

		// Filter by status if necessary
		if status != "all" && status != taskStatus {
			continue
		}

    taskFormat := fmt.Sprintf("ID:%d  %-12s  %s   (%s)   (%s)\n", task.ID, taskStatus, task.Description, task.CreatedAt.Format("2006-01-02 15:04:05"), task.UpdatedAt.Format("2006-01-02 15:04:05"))
		// Format the task and write to the output
		output.WriteString(taskFormat)
		output.WriteString(strings.Repeat("─", len(taskFormat)) + "\n")
	}

	return output.String()
}