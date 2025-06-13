package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"ToDoInventory/internal/service"
	datatypes "ToDoInventory/internal/storage/databaseTypes"
)

func (h *Handler) PostToDo(c *gin.Context) {
	var newToDo datatypes.ToDo

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

	c.IndentedJSON(http.StatusCreated, service.ConvToToDoDTO(newToDo))
}
