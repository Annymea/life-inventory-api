// @title           Meine API
// @version         1.0
// @description     Dokumentation der REST-API
// @BasePath        /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter your token with the 'Bearer ' prefix. Example: 'Bearer eyJhbGciOi...'

// @tag.name Entry
// @tag.description All entry-related endpoints

// @tag.name Auth
// @tag.description Login & registration
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
	router.RegisterRoutes(r, db)

	//start server
	r.Run("localhost:8080")
}
