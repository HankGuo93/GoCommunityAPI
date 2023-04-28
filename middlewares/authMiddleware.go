package middlewares

import (
	"GoCommunityAPI/helpers"
	"GoCommunityAPI/models"
	"GoCommunityAPI/repositories"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
)

func RequereAuth(c *gin.Context) {
	tokenString, _ := c.Cookie("Authorization")

	if tokenString == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSAPSS); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header)
		}
		publicKey, err := helpers.GetPublicKey()
		if err != nil {
			return nil, fmt.Errorf("Failed to parse the token")
		}
		return publicKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if (float64)(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		var user models.UserModel
		user, _ = repositories.FindOneUserById(int(claims["sub"].(float64)))

		if user.Id == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Set("user", user)

		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
