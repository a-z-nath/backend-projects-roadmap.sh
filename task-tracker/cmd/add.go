package cmd

import (
	"fmt"

	"github.com/a-z-nath/backend-projects/task-tracker/internal/task"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use: "add [Task Name]",
	Short: "Add a new task",
	Long: `The 'add' command allows you to add a new task to your task tracker.

You can specify the task as a single argument in quotes if it contains spaces. 
For example:

  ./task-tracker add "Finish the project report"

This will add the task to your list, and you can view it later or mark it as complete.\n`,

	Args: cobra.MinimumNArgs(1),

	Run: addTask,
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func addTask(cmd *cobra.Command, args []string) {
	taskName := args[0]

	taskList, err := task.ReadFromFile()

	if err != nil {
		panic(err)
	}
	var id int
	if len(*taskList) == 0 {
		id = 1
		newTask := task.NewTask(id, taskName)
		taskList.Add(newTask)
	} else {
		lastTask := (*taskList)[len(*taskList)-1]
		id = lastTask.ID + 1
		newTask := task.NewTask(id, taskName)
		taskList.Add(newTask)
	}

	fmt.Printf("Task is successfully added. Task ID: %d\n", id)
}
