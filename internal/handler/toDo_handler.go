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
	toDoList := []models.ToDo{}

	//get all items from the database
	result := h.DB.Find(&toDoList)

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if len(toDoList) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "no items found"})
		return
	}

	c.IndentedJSON(http.StatusOK, toDoList)
}

func (h *Handler) PostToDo(c *gin.Context) {
	var newToDo models.ToDo

	err := c.BindJSON(&newToDo)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := h.DB.Create(&newToDo)

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if newToDo.ID == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

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
