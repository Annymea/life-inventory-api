package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"ToDoInventory/internal/models"
	"ToDoInventory/internal/service"
)

func (h *Handler) PostToDo(c *gin.Context) {
	var newToDo models.ToDoDTO

	err := c.BindJSON(&newToDo)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if newToDo.ID == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	result := h.DB.Create(service.ConvToDatabaseToDo(newToDo))

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newToDo)
}
