package authhandler

import (
	"net/http"

	"LifeInventoryApi/internal/models"
	"LifeInventoryApi/internal/storage/datatypes"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (h *AuthHandler) CreateUser(c *gin.Context) {
	var authInput models.AuthInput

	if err := c.ShouldBindJSON(&authInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFound datatypes.User
	h.DB.Where("username=?", authInput.Username).Find(&userFound)

	if userFound.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username already used"})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(authInput.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := datatypes.User{
		Username: authInput.Username,
		Password: string(passwordHash),
	}

	h.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}
