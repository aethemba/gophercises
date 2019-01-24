package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task-cli is a fast and simple tool for maintaning todo items",
	Long:  "Do we really need to make this long?",
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("Running")
	// },
	// RunE: func(cmd *cobra.Command, args []string) error {
	// 	return errors.New("\nPlease provide a command, e.g. add or list\n")
	// },
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
