package entryhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"LifeInventoryApi/internal/models"
	"LifeInventoryApi/internal/service"
)

// @Summary			Create new entry
// @Description 	Creates a new entry
// @Param 			entry body models.CreateEntryDto true "New entry"
// @Tags			Entry
// @Success 		201 {object} models.EntryDto
// @Failure 		400
// @Failure 		500
// @Security 		BearerAuth
// @Router 			/entry [post]
func (h *EntryHandler) PostEntry(c *gin.Context) {
	newEntry := models.CreateEntryDto{}

	user, ok := service.GetUser(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	err := c.BindJSON(&newEntry)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dbEntry := service.ToDbEntryFromCreate(newEntry)
	dbEntry.UserID = user.ID

	result := h.DB.Create(&dbEntry)

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, service.ToEntryDto(dbEntry))
}
