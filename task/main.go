package main

import (
	"gophercises/task/cmd"
	"gophercises/task/tasks"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()

	dbPath := filepath.Join(home, "tasks.db")
	err := tasks.Init(dbPath)
	if err != nil {
		panic(err)
	}
	cmd.Execute()
}
