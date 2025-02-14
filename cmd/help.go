// Package cmd provides command-line subcommands for the GoDoIt application.
package cmd

import (
	"fmt"
)

// Prints all the available commands
func Help() {
	fmt.Println("Usage: GoDoIt <command> [arguments]")
	fmt.Println("")
	fmt.Println("Available commands:")
	fmt.Println("  add <task> <cat>    		Add a new task")
	fmt.Println("  list <done>         		List all tasks")
	fmt.Println("  update <id> <task> <cat>	Update an existing task")
	fmt.Println("  delete <id>         		Delete an existing task")
	fmt.Println("  help                		Show this help message")
	fmt.Println("")
}
