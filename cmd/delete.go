// Package cmd provides command-line subcommands for the GoDoIt application.
package cmd

import (
	"flag"
	"fmt"
	"log"

	"github.com/milinddethe15/GoDoIt/todo"
)

// Delete a task item
func DeleteTask(todos *todo.Todos, args []string) {
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	// If no --id=1 flag defined, deleteID will default to 0
	// but if --id is present but didn't set any value, an error will be thrown
	deleteID := deleteCmd.Int("id", 0, "The id of todo to be deleted")

	// Parse the arguments for the "delete" subcommand
	deleteCmd.Parse(args)

	err := todos.Delete(*deleteID)
	if err != nil {
		log.Fatal(err)
	}

	err = todos.Store(GetJsonAddr())
	if err != nil {
		log.Fatal(err)
	}

	todos.Print(2, "")
	fmt.Println("Todo item deleted successfully.")
}
