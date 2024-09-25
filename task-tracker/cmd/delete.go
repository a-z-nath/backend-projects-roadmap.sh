package cmd

import (
	"fmt"
	"strconv"

	"github.com/a-z-nath/backend-projects/task-tracker/internal/task"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use: "delete [ID]",
	Short: "Delete a task from from your tasks",
	Long: `The 'delete' command allows you to delete a task with [ID] from your task tracker.

You can specify the task you want to delete by its ID
For example:

  ./task-tracker delete 1

This will delete the task from your tasks with [ID: 1]

`,

	Args: cobra.MinimumNArgs(1),

	Run: deleteTask,
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func deleteTask(cmd *cobra.Command, args []string) {
	taskID, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("Invalid task ID. You've to provide a number not: %v\n", err)
	}

	taskList, err := task.ReadFromFile()

	if err != nil {
		panic(err)
	}

	isDeleted := false
	if len(*taskList) == 0 {
		fmt.Printf(`You have tasks currently. You can add task using:
	./task-tracker add [Task Name]%v`,"\n")
		return
	} 
	isDeleted = taskList.DeleteTask(taskID)
	if !isDeleted {
		fmt.Printf("There is no task with task ID: %v\n", taskID)
		return
	}

	fmt.Printf("Task with ID: %v is deleted successfully.\n", taskID)
}