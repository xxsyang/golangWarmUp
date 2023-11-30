package cmd

import "time"

type item struct {
	todoList     string
	isDone       bool
	createTime   time.Time
	finishedTime string
}

type todoList []item

func addToList(content string) []item {
	addItem := new(item)
	addItem.todoList = content
	addItem.isDone = false
	addItem.createTime = time.Now()
	addItem.finishedTime = "Not finished yet! Do it right now!!!!!!!!!!!!!!"

	return todoList.append(addItem)
}

func markDone(index int) []item {

}

func deleteFromList(index int) []item {
	var pointer
	if index < 0 || index > len(todoList) {

	}

}
