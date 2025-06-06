package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"ToDoInventory/models"
)

var toDoList = []models.ToDo{
	{ID: "1", Title: "Einkaufen gehen", Done: false},
	{ID: "2", Title: "Go lernen", Done: true},
	{ID: "3", Title: "API testen", Done: false},
}

func getToDoList(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, toDoList)
}

func postToDo(c *gin.Context) {
	var newToDo models.ToDo

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

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery()) //ToDo: Was ist das?

	router.GET("/list", getToDoList)
	router.POST("/todo", postToDo)
	router.GET("/todo/:id", getListItemById)

	return router
}
