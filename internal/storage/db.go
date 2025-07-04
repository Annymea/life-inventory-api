package storage

import (
	"LifeInventoryApi/internal/storage/datatypes"
	"log"
	"os"

	"github.com/joho/godotenv"
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

func buildDbUrl() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	if user == "" || pass == "" || host == "" || port == "" || name == "" {
		log.Fatalf("Error reading env variables")
		return ""
	}

	return "postgres://" + user + ":" + pass + "@" + host + ":" + port + "/" + name
}

func connectDb(dbUrl string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})

	if err != nil {
		log.Fatalln("Failed to connect to DB:", err)
		return nil
	}

	return db
}

func InitDb() *gorm.DB {
	dbUrl := buildDbUrl()
	if dbUrl == "" {
		return nil
	}

	db := connectDb(dbUrl)
	if db == nil {
		return nil
	}

	db.AutoMigrate(&datatypes.Entry{})
	db.AutoMigrate(&datatypes.User{})

	return db
}
