package entryhandler

import (
	"LifeInventoryApi/internal/models"
	"LifeInventoryApi/internal/service"
	"LifeInventoryApi/internal/storage/datatypes"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary Update a entry
// @Description Updates the entry with the given ID. The existing item will be completely overwritten by the provided data. The item will be identified by the ID of the item.
// @Param Entry body models.EntryDto true "Updated Entry object"
// @Success 200 {object} models.EntryDto
// @Failure 400
// @Failure 404
// @Router /entry [put]
func (h *EntryHandler) UpdateEntry(c *gin.Context) {
	updatedEntry := models.EntryDto{}

	//read json
	err := c.BindJSON(&updatedEntry)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//parse id and check if it is correct
	id, err := uuid.Parse(updatedEntry.ID)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}

	//look if there is an existing item with this id
	existing := datatypes.Entry{}
	result := h.DB.First(&existing, "id = ?", id)
	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Entry not found"})
		return
	}

	//update item if existing
	result = h.DB.Model(&existing).Updates(service.ToDbEntry(updatedEntry))
	if result.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, updatedEntry)
}
