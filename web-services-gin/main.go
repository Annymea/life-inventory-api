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

func postToDo(c *gin.Context) {
	var newToDo toDo

	err := c.BindJSON(&newToDo)
	if err != nil {
		return
	}

	toDoList = append(toDoList, newToDo)

	c.IndentedJSON(http.StatusCreated, newToDo)
}

func main() {
	router := gin.Default()

	router.GET("/list", getToDoList)
	router.POST("/todo", postToDo)

	router.Run("localhost:8080")
}
