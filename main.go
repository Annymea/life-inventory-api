// @title           ToDo Inventory API
// @version         1.0
// @description     API zur Verwaltung von ToDos.
// @host            localhost:8080
// @BasePath        /
package main

import (
	_ "LifeInventoryApi/docs"
	"LifeInventoryApi/internal/router"
	"LifeInventoryApi/internal/storage"

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
	//r.Use(cors.Default())
	router.RegisterRoutes(r, db)

	//start server
	r.Run("localhost:8080")
}
