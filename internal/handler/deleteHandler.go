package handler

import (
	"LifeInventoryApi/internal/storage/datatypes"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary	Delete a entry by ID
// @Description Deletes exactly one entry by ID
// @Param id path string true "ID of the entry"
// @Success 200
// @Failure 500
// @Router /entry/{id} [delete]
func (h *Handler) DeleteEntryById(c *gin.Context) {
	id := c.Param("id")

	result := h.DB.Delete(&datatypes.Entry{}, id)

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, fmt.Sprintf("Deleted item id: %s", id))
}
