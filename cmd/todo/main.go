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
	list := flag.Bool("list", false, "list all todos")

	flag.Parse()

	todos := cmd.TodoList{}

	if err, _ := cmd.ReadFromFile(todos, storeFile); err != nil {
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

		cmd.InsertItem(todos, task)
		fmt.Println("Todo List after insertion:", todos)
		err = cmd.WriteToFile(todos, storeFile)
		if err != nil {
			_, err := fmt.Fprintln(os.Stderr, err.Error())
			if err != nil {
				return
			}
			os.Exit(1)
		}

	case *complete > 0:
		var list, err = cmd.MarkDone(todos, *complete)
		_ = list
		if err != nil {
			//fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		err = cmd.WriteToFile(todos, storeFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

	case *del > 0:
		var list, err = cmd.DeleteItem(todos, *del)
		_ = list
		if err != nil {
			//fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		err = cmd.WriteToFile(todos, storeFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

	case *list:
		cmd.PrintTable(todos)
	default:
		fmt.Fprintln(os.Stdout, "invalid command")
		os.Exit(0)
	}

	//if *add {
	//	content, err := getInput(os.Stdin, flag.Args()...)
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//
	//	todoList = cmd.InsertItem(content)
	//	fmt.Println("Todo List after insertion:", todoList)
	//}
	//
	//if *complete != 0 {
	//	todoList, err := cmd.MarkDone(*complete)
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//
	//	fmt.Println("Todo List after completion:", todoList)
	//}
	//
	//if *del != 0 {
	//	todoList, err := cmd.DeleteItem(*del)
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//
	//	fmt.Println("Todo List after deletion:", todoList)
	//}
	//
	//if *list {
	//	cmd.PrintTable()
	//}

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
