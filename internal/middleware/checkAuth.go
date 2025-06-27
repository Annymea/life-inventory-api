package middleware

import (
	"LifeInventoryApi/internal/storage/datatypes"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func CheckAuth(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		//Im header muss hier das auth token mit dabei sein
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		//liest den header aus und schaut, ob es das richtige format hat
		authToken := strings.Split(authHeader, " ")
		if len(authToken) != 2 || authToken[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		//liest den token und checkt den auf richtige signatur und auf gültigkeit
		tokenString := authToken[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("SECRET")), nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		//claims sind inhalte, die im token gespeichert werden
		//mit mapClaims werden die daten zugreifbar gemacht
		claims, ok := token.Claims.(jwt.MapClaims)
		//ok ist die typüberprüfung
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		//checkt ob das token noch nicht abgelaufen ist
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token expired"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		//checkt ob der user noch in der datenbank existiert - falls der user gelöscht wurde, das token aber noch nicht abgelaufen ist
		var user datatypes.User
		db.Where("ID=?", claims["id"]).Find(&user)

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		//current user wird im kontext gesetzt, um die nächsten requests abfertigen zu lassen
		c.Set("currentUser", user)

		c.Next()
	}
}
