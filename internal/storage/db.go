package storage

import (
	"LifeInventoryApi/internal/storage/datatypes"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func AddNewItem(db *gorm.DB, newItem datatypes.Entry) (newId string, err error) {
	res := db.Create(&newItem)

	if res.Error != nil {
		log.Println("Fehler beim Erstellen:", res.Error)
		return "", res.Error
	}

	return newItem.ID, nil
}

func InitDb() *gorm.DB {

	//TODO: Das wirkt noch nicht wirklich gut, hier ist einfach user und pw einfach hard reingecoded...
	dbUrl := "postgres://pg:pass@localhost:5432/entryInventory"
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
		return nil
	}
	db.AutoMigrate(&datatypes.Entry{})

	return db
}
