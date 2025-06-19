package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"LifeInventoryApi/internal/models"
	"LifeInventoryApi/internal/service"
)

// @Summary	Create new entry
// @Description Creates a new entry
// @Param entry body models.EntryDto true "New entry"
// @Success 201 {object} models.EntryDto
// @Failure 400
// @Failure 500
// @Router /entry [post]
func (h *Handler) PostEntry(c *gin.Context) {
	var newEntry models.EntryDto

	err := c.BindJSON(&newEntry)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if newEntry.ID == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	result := h.DB.Create(service.ConvToDatabaseToDo(newEntry))

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newEntry)
}
