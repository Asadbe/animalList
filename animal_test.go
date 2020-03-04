package main

import (
	"testing"
)

var (
	a AnimalListint
)

func TestNewAnimalList(t *testing.T) {
	_, err = NewAnimalList()
	if err != nil {
		t.Error("NewTaskList cannot be created")
	}
}

// func Test_Add(t *testing.T) {
// 	a, err = NewAnimalList()
// 	if err != nil {
// 		t.Error("NewTaskList cannot be created")
// 	}
// 	a.Add(AnimalList{
// 		name:  "Asadbek",
// 		voice: "Javooob",
// 		leg:   5,
// 	}) 
// }

func Test_Update(t *testing.T) {
	a, err = NewAnimalList()
	if err != nil {
		t.Error("NewTaskList cannot be created")
	}

	a.Update(

		"Rustam",
		"Javooob",
		4, 2,
	)
	
}
func Test_Delete(t *testing.T) {
	a, err = NewAnimalList()
	if err != nil {
		t.Error("NewTaskList cannot be created")
	}

	a.Delete(
2,
	)
	
}
func Test_GetTask(t *testing.T) {
	a, err = NewAnimalList()
	if err != nil {
		t.Error("NewTaskList cannot be created")
	}

	a.GetTask(
		2,
	)
}
