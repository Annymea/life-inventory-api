package service

import (
	"ToDoInventory/internal/models"
	databasetypes "ToDoInventory/internal/storage/databaseTypes"
)

//Hier wird das mapping von den entgegengenommenen Types zu Datenbanktypen vorgenommen

func ConvToGetResponse(dbToDo databasetypes.ToDo) models.GetToDoResponse {
	return models.GetToDoResponse{
		PlannedDate: dbToDo.PlannedDate,
		Done:        dbToDo.Done,
		Title:       dbToDo.Title,
		ID:          dbToDo.ID,
	}
}

func ConvListToGetResponse(dbToDoList []databasetypes.ToDo) []models.GetToDoResponse {
	convertedList := []models.GetToDoResponse{}

	for _, toDo := range dbToDoList {
		resp := ConvToGetResponse(toDo)
		convertedList = append(convertedList, resp)
	}

	return convertedList
}
