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

func (ptr *TodoList) InsertItem(content string) {
	*ptr = append(*ptr,
		item{
			content,
			false,
			time.Now(),
			"Not finished yet! Do it right now!!!!!!!!!!!!!!"})

	fmt.Println("Todo List add successfully!")
}

func (ptr *TodoList) MarkDone(index int) error {
	list := *ptr

	if index < 0 || index > len(list) {
		return errors.New("index out of range")
	}

	list[index].isDone = true
	fmt.Println("Todo List marked as done successfully!")
	list[index].finishedTime = time.Now().String()
	fmt.Println("Todo List fish time done successfully!")

	return nil
}

func (ptr *TodoList) DeleteItem(index int) error {
	list := *ptr

	if index < 0 || index > len(list) {
		return errors.New("index out of range")
	}

	list = append(list[:index], list[index+1:]...)

	return nil
}

func (ptr *TodoList) updateItem(index int, content string) error {
	list := *ptr

	if index < 0 || index > len(list) {
		return errors.New("index out of range")
	}

	list[index].goal = content

	return nil
}

//func listAll() TodoList {
//	return listAll()
//}

func (ptr *TodoList) ReadFromFile(fileName string) error {

	content, err := os.ReadFile(fileName)

	if err != nil {
		return errors.New("read file error")
	}

	if fileName == "" {
		return errors.New("file name is empty")
	}

	if len(content) == 0 {
		return errors.New("file is empty")
	}

	// Unmarshal the JSON data into the todoList variable
	err = json.Unmarshal(content, ptr)
	if err != nil {
		return errors.New("parsing file error")
	}

	return nil

}

func (ptr *TodoList) WriteToFile(fileName string) error {
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
	data, err := json.Marshal(ptr)
	if err != nil {
		return errors.New("marshal error")
	}

	return os.WriteFile(fileName, data, 8964)
}

func (ptr *TodoList) PrintTable() {
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

	for idx, item := range *ptr {
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
