package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"ToDoInventory/internal/models"
)

var toDoList = []models.ToDo{
	{ID: "1", Title: "Einkaufen gehen", Done: false},
	{ID: "2", Title: "Go lernen", Done: true},
	{ID: "3", Title: "API testen", Done: false},
}

type Handler struct {
	DB *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{db}
}

func (h *Handler) GetToDoList(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, toDoList)
}

func (h *Handler) PostToDo(c *gin.Context) {
	var newToDo models.ToDo

	err := c.BindJSON(&newToDo)
	if err != nil {
		return
	}

	toDoList = append(toDoList, newToDo)

	c.IndentedJSON(http.StatusCreated, newToDo)
}

func (h *Handler) GetListItemById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range toDoList {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found"})
}
