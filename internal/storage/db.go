package storage

import (
	"ToDoInventory/internal/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func AddNewItem(db *gorm.DB, newItem models.ToDo) (newId string, err error) {
	res := db.Create(&newItem)

	if res.Error != nil {
		log.Println("Fehler beim Erstellen:", res.Error)
		return "", res.Error
	}

	return newItem.ID, nil
}

// Hier wird die Datenbank initialisiert
func InitDb() *gorm.DB {
	dbUrl := "postgres://pg:pass@localhost:5432/toDoList"
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
		return nil
	}
	db.AutoMigrate(&models.ToDo{})

	testToDo := models.ToDo{ID: "5", Title: "Test", Done: true, PlannedDate: ""}
	AddNewItem(db, testToDo)

	return db
}
