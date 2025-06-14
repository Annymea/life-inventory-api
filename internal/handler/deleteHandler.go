package handler

import (
	datatypes "ToDoInventory/internal/storage/databaseTypes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) DeleteToDoById(c *gin.Context) {
	id := c.Param("id")

	result := h.DB.Delete(&datatypes.ToDo{}, id)

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, "deleted item id")
}
