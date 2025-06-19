package handler

import (
	"LifeInventoryApi/internal/service"

	"LifeInventoryApi/internal/storage/datatypes"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary Get all todos (with filters)
// @Description Returns a list of all todos (fitting to the filters)
// @Param done query bool false "Filter by done status"
// @Param date query string false "Filter by planned date (YYYY-MM-DD)"
// @Success 200 {array} models.ToDoDTO
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /todo [get]
func (h *Handler) GetToDoListByParameters(c *gin.Context) {
	toDoList := []datatypes.ToDo{}

	doneFlag := c.Query("done")
	date := c.Query("date")

	q := h.DB.Model(&datatypes.ToDo{})

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

	result := q.Find(&toDoList)

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if len(toDoList) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "no items found"})
		return
	}

	c.IndentedJSON(http.StatusOK, service.ConvListToToDoDTO(toDoList))
}

// @Summary Get all todos
// @Description Returns a list of all todos
// @Success 200 {array} models.ToDoDTO
// @Failure 500
// @Router /list [get]
func (h *Handler) GetToDoList(c *gin.Context) {
	toDoList := []datatypes.ToDo{}

	result := h.DB.Find(&toDoList)

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if len(toDoList) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "no items found"})
		return
	}

	c.IndentedJSON(http.StatusOK, service.ConvListToToDoDTO(toDoList))
}

// @Summary	Get a todo by ID
// @Description Returns exactly one todo with the given ID
// @Param id path string true "ID of the todo"
// @Success 200 {array} models.ToDoDTO
// @Failure 404
// @Failure 500
// @Router /todo/{id} [get]
func (h *Handler) GetListItemById(c *gin.Context) {
	id := c.Param("id")

	toDoListItem := datatypes.ToDo{}

	result := h.DB.First(&toDoListItem, "id = ?", id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found"})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, service.ConvToToDoDTO(toDoListItem))
}
