package main

import (
	"ToDoInventory/internal/router"
	"ToDoInventory/internal/storage"

	"github.com/gin-gonic/gin"
)

func main() {

	//database stuff
	db := storage.InitDb()

	if db == nil {
		return
	}

	//router stuff
	gin.ForceConsoleColor()
	r := gin.Default() //Default Middleware -> Wenn ich hier was eigenes habe, dann muss ich es anpassen

	router.RegisterRoutes(r)

	//start server
	r.Run("localhost:8080")
}
