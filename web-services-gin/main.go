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

func getListItemById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range toDoList {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found"})
}

func main() {
	router := gin.Default()

	router.GET("/list", getToDoList)
	router.POST("/todo", postToDo)
	router.GET("/todo/:id", getListItemById)

	router.Run("localhost:8080")
}
