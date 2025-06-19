package service

import (
	"LifeInventoryApi/internal/models"
	"LifeInventoryApi/internal/storage/datatypes"
)

//Hier wird das mapping von den entgegengenommenen Types zu Datenbanktypen vorgenommen

func ConvToToDoDTO(dbToDo datatypes.ToDo) models.EntryDto {
	return models.EntryDto{
		PlannedDate: dbToDo.PlannedDate,
		Done:        dbToDo.Done,
		Title:       dbToDo.Title,
		ID:          dbToDo.ID,
	}
}

func ConvListToToDoDTO(dbToDoList []datatypes.ToDo) []models.EntryDto {
	convertedList := []models.EntryDto{}

	for _, toDo := range dbToDoList {
		resp := ConvToToDoDTO(toDo)
		convertedList = append(convertedList, resp)
	}

	return convertedList
}

func ConvToDatabaseToDo(dtoToDo models.EntryDto) datatypes.ToDo {
	return datatypes.ToDo{
		ID:          dtoToDo.ID,
		Title:       dtoToDo.Title,
		Done:        dtoToDo.Done,
		PlannedDate: dtoToDo.PlannedDate,
	}
}
