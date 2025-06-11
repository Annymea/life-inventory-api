package handler

import (
	"ToDoInventory/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h *Handler) GetToDoList(c *gin.Context) {
	toDoList := []models.ToDo{}

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

func (h *Handler) GetListItemById(c *gin.Context) {
	id := c.Param("id")

	toDoListItem := models.ToDo{}

	result := h.DB.First(&toDoListItem, "id = ?", id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found"})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, toDoListItem)
}
