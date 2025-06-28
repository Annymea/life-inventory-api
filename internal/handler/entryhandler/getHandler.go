package entryhandler

import (
	"LifeInventoryApi/internal/service"

	"LifeInventoryApi/internal/storage/datatypes"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// @Summary 		Get all entries (with filters)
// @Description 	Returns a list of all entries (fitting to the filters)
// @Param 			done query bool false "Filter by done status"
// @Param 			date query string false "Filter by planned date (YYYY-MM-DD)"
// @Tags			Entry
// @Success 		200 {array} models.EntryDto
// @Failure 		400
// @Failure 		404
// @Failure 		500
// @Security 		BearerAuth
// @Router 			/entry [get]
func (h *EntryHandler) GetEntryListByParameters(c *gin.Context) {
	user, ok := service.GetUser(c)
	if !ok {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	entryList := []datatypes.Entry{}

	doneFlag := c.Query("done")
	date := c.Query("date")

	query := h.DB.Model(&datatypes.Entry{})
	query = query.Where("user_id = ?", user.ID)

	if doneFlag != "" {
		boolValue, err := strconv.ParseBool(doneFlag)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid value for done"})
			return
		}

		query = query.Where("done = ?", boolValue)
	}

	if date != "" {
		query = query.Where("planned_date = ?", date)
	}

	result := query.Find(&entryList)
	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, service.ToEntryDtoList(entryList))
}

// @Summary			Get a entry by ID
// @Description 	Returns exactly one entry with the given ID
// @Param 			id path string true "ID of the entry"
// @Tags			Entry
// @Success 		200 {array} models.EntryDto
// @Failure 		404
// @Failure 		500
// @Security 		BearerAuth
// @Router 			/entry/{id} [get]
func (h *EntryHandler) GetListItemById(c *gin.Context) {
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

	c.IndentedJSON(http.StatusOK, service.ToEntryDto(entry))
}
