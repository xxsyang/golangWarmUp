package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/alexeyco/simpletable"
	"os"
	"time"
)

type item struct {
	goal         string
	isDone       bool
	createTime   time.Time
	finishedTime string
}

type TodoList []item

//var TodoList []item

func InsertItem(list TodoList, content string) TodoList {
	addItem := new(item)
	addItem.goal = content
	addItem.isDone = false
	addItem.createTime = time.Now()
	addItem.finishedTime = "Not finished yet! Do it right now!!!!!!!!!!!!!!"

	return append(list, *addItem)
}

func MarkDone(list TodoList, index int) (TodoList, error) {
	if index < 0 || index > len(list) {
		return list, errors.New("index out of range")
	}

	list[index].isDone = true
	list[index].finishedTime = time.Now().String()

	return list, nil
}

func DeleteItem(list TodoList, index int) (TodoList, error) {
	if index < 0 || index > len(list) {
		return list, errors.New("index out of range")
	}

	list = append(list[:index], list[index+1:]...)

	return list, nil
}

func updateItem(list TodoList, index int, content string) (TodoList, error) {
	if index < 0 || index > len(list) {
		return list, errors.New("index out of range")
	}

	list[index].goal = content

	return list, nil
}

//func listAll() TodoList {
//	return listAll()
//}

func ReadFromFile(list TodoList, fileName string) (TodoList, error) {

	content, err := os.ReadFile(fileName)

	if err != nil {
		return list, errors.New("read file error")
	}

	if fileName == "" {
		return nil, errors.New("file name is empty")
	}

	if len(content) == 0 {
		return nil, errors.New("file is empty")
	}

	// Unmarshal the JSON data into the todoList variable
	err = json.Unmarshal(content, &list)
	if err != nil {
		return nil, errors.New("parsing file error")
	}

	return list, nil

}

func WriteToFile(list TodoList, fileName string) error {
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
	data, err := json.Marshal(list)
	if err != nil {
		return errors.New("marshal error")
	}

	return os.WriteFile(fileName, data, 8964)
}

func PrintTable(list TodoList) {
	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "Index"},
			{Align: simpletable.AlignCenter, Text: "Goal"},
			{Align: simpletable.AlignCenter, Text: "IsDone?"},
			{Align: simpletable.AlignCenter, Text: "CreateTime"},
			{Align: simpletable.AlignCenter, Text: "FinishedTime"},
		},
	}

	var cells [][]*simpletable.Cell

	for idx, item := range list {
		idx++
		task := blue(item.goal)
		done := blue("no")
		if item.isDone {
			task = green(fmt.Sprintf("\u2705 %s", item.goal))
			done = green("yes")
		}
		cells = append(cells, *&[]*simpletable.Cell{
			{Text: fmt.Sprintf("%d", idx)},
			{Text: task},
			{Text: done},
			{Text: item.createTime.Format(time.RFC822)},
			{Text: item.finishedTime},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}

	table.SetStyle(simpletable.StyleUnicode)

	table.Println()
}

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
