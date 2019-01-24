// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	tsk "gophercises/task/tasks"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// doCmd represents the do command
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

		// if len(args) == 0 {
		// 	fmt.Println("Please provide a task number")
		// 	return
		// }

		// key, _ := strconv.Atoi(args[0])

		// db, err := bolt.Open("tasks.db", 0600, nil)

		// if err != nil {
		// 	log.Fatal(err)
		// }

		// var world = []byte("world")

		// err = db.Update(func(tx *bolt.Tx) error {
		// 	bucket := tx.Bucket(world)
		// 	if bucket == nil {
		// 		return fmt.Errorf("Bucket %q was not found", world)
		// 	}

		// 	c := bucket.Cursor()

		// 	counter := 1
		// 	var delKey []byte
		// 	for k, v := c.First(); k != nil; k, v = c.Next() {
		// 		if counter == key && string(v) == "open" {
		// 			delKey = k
		// 			break
		// 		}
		// 		if string(v) == "open" {
		// 			counter++
		// 		}

		// 	}
		// 	err := bucket.Put(delKey, []byte("completed"))
		// 	if err != nil {
		// 		return err
		// 	}

		// 	fmt.Printf("You have completed the task \"%s\".\n", delKey)
		// 	return nil
		// })

		// if err != nil {
		// 	fmt.Printf("%s\n", err)
		// }

		// defer db.Close()

	},
}

func init() {
	rootCmd.AddCommand(doCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
