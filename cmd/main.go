package main

import (
	"ToDoInventory/internal/models"
	"ToDoInventory/internal/router"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	//router stuff
	gin.ForceConsoleColor()
	r := gin.Default() //Default Middleware -> Wenn ich hier was eigenes habe, dann muss ich es anpassen

	router.RegisterRoutes(r)

	//database stuff
	dbUrl := "postgres://pg:pass@localhost:5432/toDoList"
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&models.ToDo{})

	testToDo := &models.ToDo{ID: "4", Title: "Test", Done: true, PlannedDate: ""}
	db.Create(&testToDo)

	//start server
	r.Run("localhost:8080")
}
