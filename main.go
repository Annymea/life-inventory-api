package main

import "ToDoInventory/router"

func main() {
	router := router.InitRouter()

	router.Run("localhost:8080")
}
