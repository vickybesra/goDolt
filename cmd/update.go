// Package cmd provides command-line subcommands for the GoDoIt application.
package cmd

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/milinddethe15/GoDoIt/todo"
)

// Update a task item
func UpdateTask(todos *todo.Todos, args []string) {

	// Define the subcommand to update todo item
	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	updateID := updateCmd.Int("id", 0, "The id of todo to be updated")
	updateCat := updateCmd.String("cat", "", "The to-be-updated category of todo")
	updateTask := updateCmd.String("task", "", "To to-be-updated content of todo")
	updateDone := updateCmd.Int("done", 1, "The to-be-updated status of todo")

	// Parse the arguments for the "update" subcommand
	updateCmd.Parse(args)

	if *updateID == 0 {
		fmt.Println("Error: the --id flag is required for the 'update' subcommand.")
		os.Exit(1)
	}
	err := todos.Update(*updateID, *updateTask, *updateCat, *updateDone)
	if err != nil {
		log.Fatal(err)
	}

	err = todos.Store(GetJsonAddr())
	if err != nil {
		log.Fatal(err)
	}

	todos.Print(2, "")
	fmt.Println("Todo item updated successfully.")
}
