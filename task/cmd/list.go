package cmd

import (
	"fmt"
	"os"

	"gophercises/task/tasks"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "lists all of our incomplete tasks",
	Long:  "lists all of our incomplete tasks",
	Run: func(cmd *cobra.Command, args []string) {

		tasks, err := tasks.ListTasks()

		if err != nil {
			fmt.Println("Something went wrong: ", err.Error())
			os.Exit(1)
		}

		if len(tasks) == 0 {
			fmt.Println("No tasks")
			return
		}

		fmt.Println("You have the following tasks")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task.Value)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
