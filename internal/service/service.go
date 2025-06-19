package service

import (
	"LifeInventoryApi/internal/models"
	"LifeInventoryApi/internal/storage/datatypes"
)

//Hier wird das mapping von den entgegengenommenen Types zu Datenbanktypen vorgenommen

func ConvToEntryDto(entryDto datatypes.Entry) models.EntryDto {
	return models.EntryDto{
		PlannedDate: entryDto.PlannedDate,
		Done:        entryDto.Done,
		Title:       entryDto.Title,
		ID:          entryDto.ID,
	}
}

func ConvListToEntryDto(dbEntryList []datatypes.Entry) []models.EntryDto {
	convertedList := []models.EntryDto{}

	for _, entry := range dbEntryList {
		resp := ConvToEntryDto(entry)
		convertedList = append(convertedList, resp)
	}

	return convertedList
}

func ConvToDbEntry(entryDto models.EntryDto) datatypes.Entry {
	return datatypes.Entry{
		ID:          entryDto.ID,
		Title:       entryDto.Title,
		Done:        entryDto.Done,
		PlannedDate: entryDto.PlannedDate,
	}
}
