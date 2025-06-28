package service

import (
	"LifeInventoryApi/internal/models"
	"LifeInventoryApi/internal/storage/datatypes"

	"github.com/gin-gonic/gin"
)

//Hier wird das mapping von den entgegengenommenen Types zu Datenbanktypen vorgenommen

func ToEntryDto(entryDto datatypes.Entry) models.EntryDto {
	return models.EntryDto{
		PlannedDate: entryDto.PlannedDate,
		Done:        entryDto.Done,
		Title:       entryDto.Title,
		ID:          entryDto.ID,
	}
}

func ToEntryDtoList(dbEntryList []datatypes.Entry) []models.EntryDto {
	convertedList := []models.EntryDto{}

	for _, entry := range dbEntryList {
		resp := ToEntryDto(entry)
		convertedList = append(convertedList, resp)
	}

	return convertedList
}

func ToDbEntry(entryDto models.EntryDto) datatypes.Entry {
	return datatypes.Entry{
		ID:          entryDto.ID,
		Title:       entryDto.Title,
		Done:        entryDto.Done,
		PlannedDate: entryDto.PlannedDate,
	}
}

func ToDbEntryFromCreate(entryDto models.CreateEntryDto) datatypes.Entry {
	return datatypes.Entry{
		Title:       entryDto.Title,
		Done:        entryDto.Done,
		PlannedDate: entryDto.PlannedDate,
	}
}

func GetUser(c *gin.Context) (datatypes.User, bool) {
	userRaw, ok := c.Get("currentUser")
	if !ok {
		return datatypes.User{}, false
	}
	user, ok := userRaw.(datatypes.User)
	if !ok {
		return datatypes.User{}, false
	}
	return user, true
}
