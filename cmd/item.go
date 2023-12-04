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

var TodoList []item

func InsertItem(content string) []item {
	addItem := new(item)
	addItem.goal = content
	addItem.isDone = false
	addItem.createTime = time.Now()
	addItem.finishedTime = "Not finished yet! Do it right now!!!!!!!!!!!!!!"

	return append(TodoList, *addItem)
}

func MarkDone(index int) ([]item, error) {
	if index < 0 || index > len(TodoList) {
		return TodoList, errors.New("index out of range")
	}

	TodoList[index].isDone = true
	TodoList[index].finishedTime = time.Now().String()

	return TodoList, nil
}

func DeleteItem(index int) ([]item, error) {
	if index < 0 || index > len(TodoList) {
		return TodoList, errors.New("index out of range")
	}

	TodoList = append(TodoList[:index], TodoList[index+1:]...)

	return TodoList, nil
}

func updateItem(index int, content string) ([]item, error) {
	if index < 0 || index > len(TodoList) {
		return TodoList, errors.New("index out of range")
	}

	TodoList[index].goal = content

	return TodoList, nil
}

func listAll() []item {
	return TodoList
}

func ReadFromFile(fileName string) ([]item, error) {

	content, err := os.ReadFile(fileName)

	if err != nil {
		return TodoList, errors.New("read file error")
	}

	if fileName == "" {
		return nil, errors.New("file name is empty")
	}

	if len(content) == 0 {
		return nil, errors.New("file is empty")
	}

	// Unmarshal the JSON data into the todoList variable
	err = json.Unmarshal(content, &TodoList)
	if err != nil {
		return nil, errors.New("parsing file error")
	}

	return TodoList, nil

}

func WriteToFile(fileName string) error {
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
	data, err := json.Marshal(TodoList)
	if err != nil {
		return errors.New("marshal error")
	}

	return os.WriteFile(fileName, data, 8964)
}

func PrintTable() {
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

	for idx, item := range *t {
		idx++
		task := blue(item.Task)
		done := blue("no")
		if item.Done {
			task = green(fmt.Sprintf("\u2705 %s", item.Task))
			done = green("yes")
		}
		cells = append(cells, *&[]*simpletable.Cell{
			{Text: fmt.Sprintf("%d", idx)},
			{Text: task},
			{Text: done},
			{Text: item.CreatedAt.Format(time.RFC822)},
			{Text: item.CompletedAt.Format(time.RFC822)},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}

	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Span: 5, Text: red(fmt.Sprintf("You have %d pending todos", t.CountPending()))},
	}}

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
