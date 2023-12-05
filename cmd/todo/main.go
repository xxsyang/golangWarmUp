package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"github.com/xxsyang/golangWarmUp"
	"io"
	"os"
	"strings"
)

const (
	storeFile = "todoList.json"
)

func main() {
	add := flag.Bool("add", false, "add a new todo")
	complete := flag.Int("complete", 0, "mark a todo as completed")
	del := flag.Int("del", 0, "delete a todo")
	list := flag.Bool("list", false, "list all todo")

	flag.Parse()

	todoList := &cmd.TodoList{}

	if err := todoList.ReadFromFile(storeFile); err != nil {
		fmt.Println(err)
		return
	}

	switch {
	case *add:
		task, err := getInput(os.Stdin, flag.Args()...)
		if err != nil {
			_, err := fmt.Fprintln(os.Stderr, err.Error())
			if err != nil {
				return
			}
			os.Exit(1)
		}

		todoList.InsertItem(task)
		fmt.Println("Todo List after insertion:", todoList)
		err = todoList.WriteToFile(storeFile)
		if err != nil {
			_, err := fmt.Fprintln(os.Stderr, err.Error())
			if err != nil {
				return
			}
			os.Exit(1)
		}

	case *complete > 0:
		var err = todoList.MarkDone(*complete)
		if err != nil {
			_, err := fmt.Fprintln(os.Stderr, err.Error())
			if err != nil {
				return
			}
			os.Exit(1)
		}
		err = todoList.WriteToFile(storeFile)
		if err != nil {
			_, err := fmt.Fprintln(os.Stderr, err.Error())
			if err != nil {
				return
			}
			os.Exit(1)
		}

	case *del > 0:
		var err = todoList.DeleteItem(*del)
		if err != nil {
			_, err := fmt.Fprintln(os.Stderr, err.Error())
			if err != nil {
				return
			}
			os.Exit(1)
		}
		err = todoList.WriteToFile(storeFile)
		if err != nil {
			_, err := fmt.Fprintln(os.Stderr, err.Error())
			if err != nil {
				return
			}
			os.Exit(1)
		}

	case *list:
		todoList.PrintTable()
	default:
		_, err := fmt.Fprintln(os.Stdout, "invalid command")
		if err != nil {
			return
		}
		os.Exit(0)
	}

}

func getInput(r io.Reader, args ...string) (string, error) {

	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	text := scanner.Text()

	if len(text) == 0 {
		return "", errors.New("empty todo is not allowed")
	}

	return text, nil

}
