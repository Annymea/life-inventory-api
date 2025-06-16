package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"ToDoInventory/internal/models"
	"ToDoInventory/internal/service"
)

// PostToDo erstellt einen neuen ToDo-Eintrag
// @Summary Neues ToDo anlegen
// @Description Erstellt ein neues ToDo
// @Tags todos
// @Accept json
// @Produce json
// @Param todo body models.ToDoDTO true "Neues ToDo"
// @Success 201 {object} models.ToDoDTO
// @Failure 400
// @Failure 500
// @Router /todos [post]
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
