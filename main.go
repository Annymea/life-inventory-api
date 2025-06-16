// @title           ToDo Inventory API
// @version         1.0
// @description     API zur Verwaltung von ToDos.
// @host            localhost:8080
// @BasePath        /
package main

import (
	_ "ToDoInventory/docs"
	"ToDoInventory/internal/router"
	"ToDoInventory/internal/storage"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	router.RegisterRoutes(r, db)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//start server
	r.Run("localhost:8080")
}
