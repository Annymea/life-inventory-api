package authhandler

import (
	"LifeInventoryApi/internal/storage/datatypes"

	"github.com/gin-gonic/gin"
)

func (h *AuthHandler) Register(c *gin.Context) {
	input := datatypes.User{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	//erstmal ohne hashed pw
	//hashedPassword, err := HashPassword(input.Password)
	//if err != nil {
	//	c.JSON(500, gin.H{"error": "Passwort konnte nicht gehasht werden"})
	//	return
	//}

	user := datatypes.User{Username: input.Username, Password: input.Password}

	if err := h.DB.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": "Benutzer konnte nicht erstellt werden"})
		return
	}

	c.JSON(200, gin.H{"message": "Benutzer erstellt", "user": user.Username})
}
