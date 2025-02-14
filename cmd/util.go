// Package cmd provides command-line subcommands for the GoDoIt application.
package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/common-nighthawk/go-figure"
)

// GetJsonAdd will grab the .todos.json file located at user home directory
func GetJsonAddr() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(homeDir, ".todos.json")
}

// getUserApproval will get the user's approval when creating an empty json file
func GetUserApproval() bool {
	confirmMessage := "Need to create an empty \".todos.json\" file in your home directory, continue? (y/n): "
	var s string
	for {
		fmt.Print(confirmMessage)
		fmt.Scanln(&s)

		if s == "y" || s == "yes" {
			return true
		}
		if s == "n" || s == "no" {
			return false
		}

		fmt.Println("Invalid input. Please enter 'y' or 'n'.")
	}
}

func CheckJson() bool {
	// check if .todos.json already exists in user home directory
	_, err := os.Stat(GetJsonAddr())
	return err == nil
}

func CreateJson() {
	filepath := GetJsonAddr()
	file, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Println("Successfully created a \".todos.json\" file in your home directory.")
}

func WelcomePrint() {
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("                          Welcome to")
	myFigure := figure.NewColorFigure("Go Do It", "banner3-D", "blue", true)
	myFigure.Print()
	fmt.Println("               a cli app for your productivity")
	fmt.Println("----------------------------------------------------------------")
}
