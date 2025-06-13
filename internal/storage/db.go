package storage

import (
	databasetypes "ToDoInventory/internal/storage/databaseTypes"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func AddNewItem(db *gorm.DB, newItem databasetypes.ToDo) (newId string, err error) {
	res := db.Create(&newItem)

	if res.Error != nil {
		log.Println("Fehler beim Erstellen:", res.Error)
		return "", res.Error
	}

	return newItem.ID, nil
}

func InitDb() *gorm.DB {

	//TODO: Das wirkt noch nicht wirklich gut, hier ist einfach user und pw einfach hard reingecoded...
	dbUrl := "postgres://pg:pass@localhost:5432/toDoList"
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
		return nil
	}
	db.AutoMigrate(&databasetypes.ToDo{})

	testToDo := databasetypes.ToDo{ID: "5", Title: "Test", Done: true, PlannedDate: ""}
	AddNewItem(db, testToDo)

	return db
}
