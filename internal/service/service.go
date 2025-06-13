package service

import (
	"ToDoInventory/internal/models"
	datatypes "ToDoInventory/internal/storage/databaseTypes"
)

//Hier wird das mapping von den entgegengenommenen Types zu Datenbanktypen vorgenommen

func ConvToToDoDTO(dbToDo datatypes.ToDo) models.ToDoDTO {
	return models.ToDoDTO{
		PlannedDate: dbToDo.PlannedDate,
		Done:        dbToDo.Done,
		Title:       dbToDo.Title,
		ID:          dbToDo.ID,
	}
}

func ConvListToToDoDTO(dbToDoList []datatypes.ToDo) []models.ToDoDTO {
	convertedList := []models.ToDoDTO{}

	for _, toDo := range dbToDoList {
		resp := ConvToToDoDTO(toDo)
		convertedList = append(convertedList, resp)
	}

	return convertedList
}
