package entryhandler

import (
	"LifeInventoryApi/internal/storage/datatypes"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary			Delete a entry by ID
// @Description 	Deletes exactly one entry by ID
// @Param 			id path string true "ID of the entry"
// @Tags			Entry
// @Success 		200
// @Failure 		500
// @Security 		BearerAuth
// @Router 			/entry/{id} [delete]
func (h *EntryHandler) DeleteEntryById(c *gin.Context) {
	idStr := c.Param("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}

	result := h.DB.Delete(&datatypes.Entry{}, "id = ?", id)

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, fmt.Sprintf("Deleted item id: %s", id))
}
