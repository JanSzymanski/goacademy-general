package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type ToDoStatus int // an alternative to make it string so the values are more descriptive.

const (
	Active ToDoStatus = iota + 1
	OnHold
	InProgress
	Canceled
	Done
)

type ToDo struct {
	Id      int        `json:"id"`
	Message string     `json:"message"`
	Status  ToDoStatus `json:"status"`
}

func CreateInitialToDos(how_many int) []ToDo {
	final_list := make([]ToDo, 0)
	for idx := range how_many {
		todo := ToDo{
			Id:      idx,
			Message: "Go and do something: " + strconv.Itoa(idx),
			Status:  Active,
		}
		final_list = append(final_list, todo)
	}
	return final_list
}

func DecodeStatus(s ToDoStatus) string {
	switch s {
	case 1:
		return "Active"
	case 2:
		return "OnHold"
	case 3:
		return "InProgress"
	case 4:
		return "Canceled"
	case 5:
		return "Done"
	default:
		return "Unknown"
	}
}

func PPrint(todos ...ToDo) {
	for _, todo := range todos {
		fmt.Printf("Id: %d Message: %s Status: %v\n", todo.Id, todo.Message, DecodeStatus(todo.Status))
	}
}

func PPrintJSON(todos ...ToDo) {
	b, err := json.MarshalIndent(todos, "", "  ") //MarshalIndent instead of just Marshal to make output prettier to eye
	if err != nil {
		fmt.Println("error in Marshaling")
	}
	//fmt.Println(string(b))
	os.Stdout.Write(b)
}
func JSONtoFile(filename string, todos ...ToDo) {
	b, err := json.MarshalIndent(todos, "", "  ") //MarshalIndent instead of just Marshal to make output prettier to eye
	if err != nil {
		fmt.Println("error in Marshaling")
	}
	err = os.WriteFile(filename, b, 0777)
	if err != nil {
		fmt.Println("wrror writting file")
	}
}
func ReadToDosFromFile(filename string) []ToDo {
	filebytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Reading data from file error")
	}
	var todos_from_file []ToDo
	err = json.Unmarshal(filebytes, &todos_from_file)
	if err != nil {
		fmt.Println("Decoding file bytes error")
	}
	return todos_from_file
}

func main() {
	todos := CreateInitialToDos(10)
	PPrint(todos...)
	PPrintJSON(todos...)
	filename := "todos.json"
	JSONtoFile(filename, todos...)
	fmt.Printf("Todos read from file: %s\n", filename)
	new_todos := ReadToDosFromFile(filename)
	PPrint(new_todos...)

}
