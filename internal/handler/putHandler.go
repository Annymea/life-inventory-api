package handler

import (
	"LifeInventoryApi/internal/models"
	"LifeInventoryApi/internal/service"
	"LifeInventoryApi/internal/storage/datatypes"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Update a entry
// @Description Updates the entry with the given ID. The existing item will be completely overwritten by the provided data. The item will be identified by the ID of the item.
// @Param Entry body models.ToDoDTO true "Updated Entry object"
// @Success 200 {object} models.ToDoDTO
// @Failure 400
// @Failure 404
// @Router /entry [put]
func (h *Handler) UpdateEntry(c *gin.Context) {
	updatedEntry := models.ToDoDTO{}

	//read json
	err := c.BindJSON(&updatedEntry)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//look if there is an existing item with this id
	existing := datatypes.ToDo{}
	result := h.DB.First(&existing, updatedEntry.ID)
	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "ToDo not found"})
		return
	}

	//update item if existing
	result = h.DB.Model(&existing).Updates(service.ConvToDatabaseToDo(updatedEntry))
	if result.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, updatedEntry)
}
