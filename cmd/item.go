package cmd

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

type item struct {
	goal         string
	isDone       bool
	createTime   time.Time
	finishedTime string
}

var todoList []item

func insertItem(content string) []item {
	addItem := new(item)
	addItem.goal = content
	addItem.isDone = false
	addItem.createTime = time.Now()
	addItem.finishedTime = "Not finished yet! Do it right now!!!!!!!!!!!!!!"

	return append(todoList, *addItem)
}

func markDone(index int) ([]item, error) {
	if index < 0 || index > len(todoList) {
		return todoList, errors.New("index out of range")
	}

	todoList[index].isDone = true
	todoList[index].finishedTime = time.Now().String()

	return todoList, nil
}

func deleteItem(index int) ([]item, error) {
	if index < 0 || index > len(todoList) {
		return todoList, errors.New("index out of range")
	}

	todoList = append(todoList[:index], todoList[index+1:]...)

	return todoList, nil
}

func updateItem(index int, content string) ([]item, error) {
	if index < 0 || index > len(todoList) {
		return todoList, errors.New("index out of range")
	}

	todoList[index].goal = content

	return todoList, nil
}

func listAll() []item {
	return todoList
}

func readFromFile(fileName string) ([]item, error) {

	content, err := os.ReadFile(fileName)

	if err != nil {
		return todoList, errors.New("read file error")
	}

	if fileName == "" {
		return nil, errors.New("file name is empty")
	}

	if len(content) == 0 {
		return nil, errors.New("file is empty")
	}

	// Unmarshal the JSON data into the todoList variable
	err = json.Unmarshal(content, &todoList)
	if err != nil {
		return nil, errors.New("parsing file error")
	}

	return todoList, nil

}

func writeToFile(fileName string) error {
	if fileName == "" {
		return errors.New("file name is empty")
	}
	//// Create the JSON file
	//file, err := os.Create(fileName)
	//if err != nil {
	//	fmt.Println("Error creating JSON file:", err)
	//	return nil
	//}
	//defer file.Close()

	// Marshal the todoList into JSON format
	data, err := json.Marshal(todoList)
	if err != nil {
		return errors.New("marshal error")
	}

	return os.WriteFile(fileName, data, 8964)
}

//func printTable(fileName string) {
//	items, err := readFromFile(fileName)
//
//
//	if err != nil {
//		return
//	}
//
//	if len(todoList) == 0 {
//		println("No item in the list")
//		return
//}
