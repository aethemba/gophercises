package cmd

import (
	"fmt"
	tsk "gophercises/task/tasks"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "marks a task as complete",
	Long:  "marks a task as complete and removes it from the list",
	Run: func(cmd *cobra.Command, args []string) {

		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)

			if err != nil {
				fmt.Println("Failed to parse argument: ", arg)
				os.Exit(1)
			}

			ids = append(ids, id)
		}

		tasks, err := tsk.ListTasks()

		if err != nil {
			fmt.Println("Something went wrong: ", err)
			os.Exit(1)
		}

		for _, id := range ids {
			if id < 0 || id > len(tasks) {
				fmt.Println("Invalid task number: ", id)
				continue
			}

			for i, t := range tasks {
				if i+1 == id {
					err := tsk.DeleteTask(t.Key)
					if err != nil {
						fmt.Printf("Failed to set task \"%d\" as completed. Err: %s\n", id, err)
					} else {
						fmt.Printf("Marked \"%d\" as completed\n", id)
					}
				}
			}

		}
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
