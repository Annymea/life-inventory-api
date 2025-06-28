package entryhandler

import (
	"LifeInventoryApi/internal/service"
	"LifeInventoryApi/internal/storage/datatypes"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
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
	user, ok := service.GetUser(c)
	if !ok {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	entryIdInput := c.Param("id")
	id, err := uuid.Parse(entryIdInput)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}

	entry := datatypes.Entry{}
	result := h.DB.First(&entry, "id = ? AND user_id = ?", id, user.ID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found"})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	result = h.DB.Delete(&entry)
	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete entry"})
		return
	}

	c.IndentedJSON(http.StatusOK, fmt.Sprintf("Deleted item id: %s", id))
}
