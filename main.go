package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/milinddethe15/GoDoIt/cmd"
	"github.com/milinddethe15/GoDoIt/todo"
)

func main() {
	todos := &todo.Todos{}
	flag.Parse()
	// Check which subcommand was invoked
	switch flag.Arg(0) {
	case "":
		cmd.Help()
	case "help":
		cmd.Help()
	case "add":
		cmd.Check(todos)
		cmd.AddTask(todos, os.Args[2:])
	case "list":
		cmd.Check(todos)
		cmd.ListTasks(todos, os.Args[2:])
	case "delete":
		cmd.Check(todos)
		cmd.DeleteTask(todos, os.Args[2:])
	case "update":
		cmd.Check(todos)
		cmd.UpdateTask(todos, os.Args[2:])
	default:
		fmt.Println("Invalid command. Please use 'help' command to see available commands.")
		os.Exit(1)
	}
}
