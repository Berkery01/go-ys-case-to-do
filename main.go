package main

import (
	"flag"
	"fmt"
	"time"
)

type ToDo struct {
	taskList []Task
}

type Task struct {
	id        int
	item      string
	date      string
	completed bool
}

var (
	version        = flag.Bool("v", false, "version")
	listItems      = flag.Bool("l", false, "list all items(un-completed)")
	completedItems = flag.Bool("c", false, "list completed items")
	addItems       = flag.String("a", "No value given", "add new item")
	markAsComplete = flag.Int("m", -1, "mark as complete")
	delete         = flag.Int("d", -1, "deleteItem")
	allTasks       ToDo
	allTasksPtr        = &allTasks
	itemCount      int = 0
	spaceText          = "------------------------------------------------------------------------------------"
	path               = "main.txt"
)

func main() {
	flag.Parse()
	checkInput(version, listItems, completedItems, addItems, markAsComplete, delete)
	importDataFromText()
	uploadDataToText()
}

func checkInput(v *bool, l *bool, c *bool, a *string, m *int, d *int) {
	if *v == true {
		fmt.Println("Current program version is v0.1")
	}
	if *l == true {
		allTasksPtr.listAllItems()
	}
	if *c == true {
		allTasksPtr.listCompletedItems()
	}
	if *a != "No value given" {
		allTasksPtr.addItemsToList(itemCount, "Buy Milk")
		allTasksPtr.addItemsToList(itemCount, "New Task")
		allTasksPtr.addItemsToList(itemCount, *a)
		allTasksPtr.addItemsToList(itemCount, "New Task2")
		allTasksPtr.addItemsToList(itemCount, "New Task3")
		fmt.Println(itemCount)
		allTasksPtr.markAsCompleted(2)
		allTasksPtr.deleteItemFromList(4)
		allTasksPtr.listCompletedItems()
		allTasksPtr.listAllItems()
	}
	if *m != -1 {
		allTasksPtr.markAsCompleted(*m)
	}
	if *d != -1 {
		allTasksPtr.deleteItemFromList(*d)
	}
}
func importDataFromText() {

}

func uploadDataToText() {

}

func (t *ToDo) listAllItems() {
	fmt.Println(spaceText)
	fmt.Printf("| %*s| %*s| %*s| %*s|\n",
		-5, "ID",
		-40, "Item",
		-20, "Date Created",
		-10, "Completed")
	fmt.Println(spaceText)
	for _, key := range t.taskList {
		fmt.Printf("| %*v| %*v| %*v| %*v| \n",
			-5, key.id,
			-40, key.item,
			-20, key.date,
			-10, key.completed,
		)
	}
	fmt.Println(spaceText)
}

func (t *ToDo) listCompletedItems() {
	printedBool := false
	fmt.Println(spaceText)
	fmt.Printf("| %*s| %*s| %*s| %*s|\n",
		-5, "ID",
		-40, "Item",
		-20, "Date Created",
		-10, "Completed")
	fmt.Println(spaceText)
	for index, key := range t.taskList {
		if t.taskList[index].completed == true {
			fmt.Printf("| %*v| %*v| %*v| %*v| \n",
				-5, key.id,
				-40, key.item,
				-20, key.date,
				-10, key.completed,
			)
			printedBool = true
		}
	}
	if printedBool == false {
		fmt.Println("No item found.")
	}
	fmt.Println(spaceText)
}

func (t *ToDo) addItemsToList(itemID int, itemName string) {
	newTask, itemID := createNewTask(itemID, itemName)
	t.taskList = append(t.taskList, newTask)
}

func (t *ToDo) markAsCompleted(itemID int) {
	itemIndex := itemID - 1
	t.taskList[itemIndex].completeTask()
}

func (t *ToDo) deleteItemFromList(itemID int) {
	itemIndex := itemID - 1
	if itemID > itemCount {
		fmt.Println("There is no item matched with this ID.")
	} else if itemID == itemCount {
		t.taskList = t.taskList[:len(t.taskList)-1]
	} else {
		for index, key := range t.taskList {
			if index < itemIndex {
				t.taskList[index] = key
			} else if index > itemIndex {
				t.taskList[index-1] = key
			}
		}
		t.taskList = t.taskList[:len(t.taskList)-1]
	}

}

func createNewTask(itemID int, itemName string) (Task, int) {
	itemID = itemID + 1
	itemCount = itemID
	itemDate := time.Now().Local().Format("02-01-2006 15:04:05")
	t := Task{
		id:        itemID,
		item:      itemName,
		date:      itemDate,
		completed: false, //as default
	}
	return t, itemCount
}

func (t *Task) completeTask() {
	t.completed = true
}
