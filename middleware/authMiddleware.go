package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"strings"
)

func RequireAuth(c *gin.Context) {
	// Pobierz token z nagłówka
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Brak autoryzacji",
		})
		return
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

	// Zweryfikuj token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("nieoczekiwana metoda podpisu: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Nieprawidłowy token",
		})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Ustaw ID użytkownika w kontekście
		c.Set("userId", uint(claims["sub"].(float64)))
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Nieprawidłowy token",
		})
		return
	}
}
