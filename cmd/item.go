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

func (ptr *TodoList) InsertItem(content string) {
	list := *ptr

	*ptr = append(list,
		item{
			goal:         content,
			isDone:       false,
			createTime:   time.Now(),
			finishedTime: "Not finished yet! Do it right now!!!!!!!!!!!!!!"})
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

	*ptr = append(list[:index-1], list[index:]...)

	return nil
}

func (ptr *TodoList) UpdateItem(index int, content string) error {
	list := *ptr

	if index < 0 || index > len(list) {
		return errors.New("index out of range")
	}

	list[index].goal = content

	return nil
}

func (ptr *TodoList) PrintListFirst() string {
	list := *ptr

	return list[0].goal

}

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

	fmt.Println("Todo List after insertion:", ptr)
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
	data, err := json.MarshalIndent(ptr, "", " ")
	if err != nil {
		return errors.New("marshal error")
	}

	err = os.WriteFile(fileName, data, 0666)

	if err != nil {
		return errors.New("write file error")
	}

	return nil

}

func (ptr *TodoList) CountPending() int {
	total := 0
	for _, item := range *ptr {
		if !item.isDone {
			total++
		}
	}
	return total
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
		task := cyan(item.goal)
		done := red("no")
		if item.isDone {
			task = green(fmt.Sprintf("\u2705 %s", item.goal))
			done = green("yes")
		}
		cells = append(cells, *&[]*simpletable.Cell{
			{Text: fmt.Sprintf("%d", idx)},
			{Text: task},
			{Text: done},
			{Text: fmt.Sprintf(red(item.createTime.String()))},
			{Text: fmt.Sprintf(item.finishedTime)},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}

	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Span: 5, Text: red(fmt.Sprintf("You have %d pending todos", ptr.CountPending()))},
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
