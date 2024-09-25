/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "task-tracker",
	Short: "Task Tracker is a CLI app to manage your tasks.",
	Long: `Task Tracker is a simple CLI (command line interface) tools to manage your tasks.
You can add new tasks, update curret tasks, delete a tasks. You also can update a task status.


Check out the source code for the Task Tracker project on GitHub!

Explore the code, and give a ⭐ to show your support: github.com/a-z-nath/task-tracker

Every star helps developer to encourage and build something new! 😊

Also thanks for your suppoort!!!
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.task-tracker.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


