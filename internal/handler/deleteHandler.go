package handler

import (
	"LifeInventoryApi/internal/storage/datatypes"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary	Delete a todo by ID
// @Description Deletes exactly one todo by ID
// @Param id path string true "ID of the todo"
// @Success 200
// @Failure 500
// @Router /todo/{id} [delete]
func (h *Handler) DeleteToDoById(c *gin.Context) {
	id := c.Param("id")

	result := h.DB.Delete(&datatypes.ToDo{}, id)

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, fmt.Sprintf("Deleted item id: %s", id))
}
