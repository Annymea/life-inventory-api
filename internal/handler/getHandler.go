package handler

import (
	"LifeInventoryApi/internal/service"

	"LifeInventoryApi/internal/storage/datatypes"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary Get all entries (with filters)
// @Description Returns a list of all entries (fitting to the filters)
// @Param done query bool false "Filter by done status"
// @Param date query string false "Filter by planned date (YYYY-MM-DD)"
// @Success 200 {array} models.EntryDto
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /entry [get]
func (h *Handler) GetEntryListByParameters(c *gin.Context) {
	entryList := []datatypes.Entry{}

	doneFlag := c.Query("done")
	date := c.Query("date")

	q := h.DB.Model(&datatypes.Entry{})

	if doneFlag != "" {
		boolValue, err := strconv.ParseBool(doneFlag)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid value for done"})
			return
		}

		q = q.Where("done = ?", boolValue)
	}

	if date != "" {
		q = q.Where("planned_date = ?", date)
	}

	result := q.Find(&entryList)

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if len(entryList) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "no items found"})
		return
	}

	c.IndentedJSON(http.StatusOK, service.ToEntryDtoList(entryList))
}

// @Summary Get all entries
// @Description Returns a list of all entries
// @Success 200 {array} models.EntryDto
// @Failure 500
// @Router /list [get]
func (h *Handler) GetEntryList(c *gin.Context) {
	entryList := []datatypes.Entry{}

	result := h.DB.Find(&entryList)

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if len(entryList) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "no items found"})
		return
	}

	c.IndentedJSON(http.StatusOK, service.ToEntryDtoList(entryList))
}

// @Summary	Get a entry by ID
// @Description Returns exactly one entry with the given ID
// @Param id path string true "ID of the entry"
// @Success 200 {array} models.EntryDto
// @Failure 404
// @Failure 500
// @Router /entry/{id} [get]
func (h *Handler) GetListItemById(c *gin.Context) {
	id := c.Param("id")

	entryListItem := datatypes.Entry{}

	result := h.DB.First(&entryListItem, "id = ?", id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found"})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, service.ToEntryDto(entryListItem))
}
