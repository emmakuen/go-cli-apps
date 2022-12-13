package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/emmakuen/go-cli-apps/todo"
)

const todoFilename = ".todo.json"

func main() {
	todoList := todo.List{}

	// get the saved list and save the value to todoList variable
	if err := todo.Get(&todoList, todoFilename); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	// if there's no extra argument, print out todo list
	case len(os.Args) == 1:
		for _, item := range todoList {
			fmt.Println(item.Task)
		}
	// else, add the new item to the list and save it as a file
	default:
		taskName := strings.Join(os.Args[1:], " ")
		todoList = todo.Add(todoList, taskName)

		if err := todo.Save(todoList, todoFilename); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
