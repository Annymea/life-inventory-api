package authhandler

import (
	"net/http"
	"os"
	"time"

	"LifeInventoryApi/internal/models"
	"LifeInventoryApi/internal/storage/datatypes"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func (h *AuthHandler) Login(c *gin.Context) {
	var authInput models.AuthInput

	//liest pw und username aus dem json aus
	if err := c.ShouldBindJSON(&authInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//schaut ob es einen user mit diesem usernamen schon gibt - fail fast
	var userFound datatypes.User
	h.DB.Where("username=?", authInput.Username).Find(&userFound)

	if userFound.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	//pw hashen und dann mit dem gespeicherten Pw vergleichen
	if err := bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(authInput.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
		return
	}

	//generiert ein token für den user
	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userFound.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(), //dauer für token validität hier 24h
	})

	//signiert das token mit dem secret, welches von mir vorher festgelegt wurde in getenv
	token, err := generateToken.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to generate token"})
	}

	//gibt den token wieder zurück
	c.JSON(200, gin.H{
		"token": token,
	})
}
