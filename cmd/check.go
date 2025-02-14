// Package cmd provides command-line subcommands for the GoDoIt application.
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/milinddethe15/GoDoIt/todo"
)

// Check Json file and load it
func Check(todos *todo.Todos) {
	Json := CheckJson()
	if !Json {
		fmt.Println("JSON file not found in Home dir to store your todo items.")
		ok := GetUserApproval()
		if !ok {
			fmt.Print("You've declined to create the JSON file.\n")
			os.Exit(0)
		}
		CreateJson()
	}
	if err := todos.Load(GetJsonAddr()); err != nil {
		log.Fatal(err)
	}
}
