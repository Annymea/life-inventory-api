package handler

import (
	"ToDoInventory/internal/models"
	"ToDoInventory/internal/service"
	"ToDoInventory/internal/storage/datatypes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) UpdateToDo(c *gin.Context) {
	updatedToDo := models.ToDoDTO{}

	//read json
	err := c.BindJSON(&updatedToDo)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//look if there is an existing item with this id
	existing := datatypes.ToDo{}
	result := h.DB.First(&existing, updatedToDo.ID)
	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "ToDo not found"})
		return
	}

	//update item if existing
	result = h.DB.Model(&existing).Updates(service.ConvToDatabaseToDo(updatedToDo))
	if result.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, updatedToDo)
}
