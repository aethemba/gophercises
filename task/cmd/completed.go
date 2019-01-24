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
	"log"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var completeCmd = &cobra.Command{
	Use:   "completed",
	Short: "lists your completed tasks",
	Long:  "lists your completed tasks",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := bolt.Open("tasks.db", 0600, nil)

		if err != nil {
			log.Fatal(err)
		}

		defer db.Close()

		var world = []byte("world")
		err = db.Update(func(tx *bolt.Tx) error {
			_, err := tx.CreateBucketIfNotExists(world)
			if err != nil {
				return err
			}
			return nil
		})

		err = db.View(func(tx *bolt.Tx) error {
			bucket := tx.Bucket(world)
			if bucket == nil {
				return fmt.Errorf("Bucket %q was not found", world)
			}

			c := bucket.Cursor()

			counter := 1
			fmt.Println("You completed the following tasks:")
			for k, v := c.First(); k != nil; k, v = c.Next() {
				if string(v) == "completed" {
					fmt.Printf("%d. %s\n", counter, k)
					counter++
				}
			}

			if counter == 1 {
				fmt.Println("\nNothing completed...get to work!")
			}

			return nil
		})

		if err != nil {
			log.Fatal(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(completeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
