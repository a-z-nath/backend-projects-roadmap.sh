package cmd

import (
	"fmt"

	"github.com/a-z-nath/backend-projects/task-tracker/internal/task"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use: "list",
	Short: "List all tasks or filter by status",
	Long: `The 'list' command allows you to view all tasks in your task tracker.
You can optionally filter the tasks based on their status using the '--status' flag.

Available statuses:
  - todo
  - in-progress
  - done

Examples:
  ./task-tracker list               # List all tasks
  ./task-tracker list --status todo # List tasks that are in 'todo' status
  ./task-tracker list -s done       # List tasks that are 'done'\n`,

	Run: listTasks,
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringP("status", "s", "all", "Filter tasks by status (todo, in-progress, done)")
}

func listTasks(cmd *cobra.Command, args []string) {
	status, _ := cmd.Flags().GetString("status")

	if !(status == "all" || status == "todo" || status == "in-progress" || status == "done") {
		fmt.Printf("status flag value must be one of these [todo | in-progress | done] not %s\n", status)
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

	filteredTask := taskList.FilteredTask(status)
	formatedTask := filteredTask.FormatTasks(status)
	
	fmt.Printf("Tasks (%v)", status)

	if formatedTask == "" {
		fmt.Println("\nCurrently there is no tasks with status:", status)
		return
	}
	fmt.Printf("\n%v\n\n", formatedTask)
}
