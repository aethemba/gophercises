package cmd

import (
	"fmt"
	"os"
	"strings"

	"gophercises/task/tasks"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "adds a new task to our list",
	Long:  "adds a new task to our list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := tasks.CreateTask(task)

		if err != nil {
			fmt.Println("Error creating task: ", err.Error())
			os.Exit(1)
		}

		fmt.Printf("Added \"%s\" to task list\n", task)

	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
