package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type toDo struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var toDoList = []toDo{
	{ID: "1", Title: "Einkaufen gehen", Done: false},
	{ID: "2", Title: "Go lernen", Done: true},
	{ID: "3", Title: "API testen", Done: false},
}

func getToDoList(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, toDoList)
}

func main() {
	router := gin.Default()

	router.GET("/list", getToDoList)

	router.Run("localhost:8080")
}
