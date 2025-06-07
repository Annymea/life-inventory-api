package main

import (
	"ToDoInventory/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.ForceConsoleColor()
	r := gin.Default() //Default Middleware -> Wenn ich hier was eigenes habe, dann muss ich es anpassen

	router.RegisterRoutes(r)

	r.Run("localhost:8080")
}
