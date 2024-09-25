package cmd

import (
	"fmt"
	"strconv"

	"github.com/a-z-nath/backend-projects/task-tracker/internal/task"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use: "update [ID] [[Updated Task Name] | [-s | --status] [todo | in-progress | done]]",
	Short: "Update a task by ID",
	Long: `The 'update' command allows you to update a task's information based on its unique ID.
You can update either the task's name or its status (todo, in-progress, done), or both.

Usage:

1. To update a task name:
   ./task-tracker update [ID] "New Task Name"

2. To update the task status:
   ./task-tracker update [ID] -s [todo | in-progress | done]

3. To update both the task name and status:
   ./task-tracker update [ID] "New Task Name" -s [todo | in-progress | done] 
   # or
   ./task-tracker update [ID] -s [todo | in-progress | done] "New Task Name"

Where:
- ID: The unique identifier of the task.
- -s | --status: An optional flag to update the status of the task.\n`,
	Args: cobra.MinimumNArgs(1),

	Run: updateTask,
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringP("status", "s", "", "Update tasks status by [todo | in-progress | done]")
}

func updateTask(cmd *cobra.Command, args []string) {
	taskID, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("Invalid task ID. You've to provide a number not: %v\n", err)
	}
	var updatedTaskName string
	if len(args) > 1 {
		updatedTaskName = args[1]
	}

	status, _ := cmd.Flags().GetString("status")

	if status == "" && updatedTaskName == "" {
		fmt.Printf("Error: You must provide either a new task name or a status to update.\n\nUpdate Command Help: \n\n%v\n", cmd.Long)
		return
	}

	if !(status == "" || status == "todo" || status == "in-progress" || status == "done") {
		fmt.Printf("status flag value must be one of these [todo | in-progress | done] not %v\n", status)
		return
	}

	taskList, err := task.ReadFromFile()

	if err != nil {
		panic(err)
	}

	if len(*taskList) == 0 {
		fmt.Printf(`You have no tasks currently. You can add task using:
	./task-tracker add [Task Name]%v`,"\n")
		return
	}
	descUpdated, statusUpdated := false, false
	var updatedMsg string
	if updatedTaskName != "" {
		descUpdated = taskList.UpdateDescription(taskID, updatedTaskName)
		updatedMsg = fmt.Sprintf("Description of task with [ID: %v] is updated\n", taskID)
	}

	if status != "" {
		statusUpdated = taskList.UpdateStatus(taskID, status)
		updatedMsg = updatedMsg + fmt.Sprintf("Status of task with [ID: %v] is updated\n", taskID)
	}
	if !descUpdated && !statusUpdated {
		fmt.Println("There is no task with ID:", taskID)
		return
	}
	fmt.Printf("Task with [ID: %v] is updated.\n%v", taskID, updatedMsg)
}
