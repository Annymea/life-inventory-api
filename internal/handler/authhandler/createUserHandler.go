package authhandler

import (
	"net/http"

	"LifeInventoryApi/internal/models"
	"LifeInventoryApi/internal/storage/datatypes"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// @Summary			Creates a user
// @Description 	Create a user which can be used to create a token
// @Param 			signup body models.AuthInput true "Username and password"
// @Tags			Auth
// @Success 		200
// @Failure 		500
// @Router 			/auth/signup [post]
func (h *AuthHandler) CreateUser(c *gin.Context) {
	var authInput models.AuthInput

	//liest pw und username aus dem json aus
	if err := c.ShouldBindJSON(&authInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//schaut ob es einen user mit diesem usernamen schon gibt
	var userFound datatypes.User
	h.DB.Where("username=?", authInput.Username).Find(&userFound)

	if userFound.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username already used"})
		return
	}

	//hasht das passwort, damit es "sicher" gespeichert werden kann (keine Ahnung ob das so ausreicht)
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(authInput.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//erstellt einen datenbanktypen mit den Werten
	user := datatypes.User{
		Username: authInput.Username,
		Password: string(passwordHash),
	}

	//Erstellt einen Eintrag in der Datenbank
	h.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}
