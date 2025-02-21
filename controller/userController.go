package controller

import (
	"1/initializers"
	"1/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

func Register(c *gin.Context) {
	// Struktura dla danych rejestracji
	var body struct {
		Nickname string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Nieprawidłowe dane",
		})
		return
	}

	// Hashowanie hasła
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Błąd podczas hashowania hasła",
		})
		return
	}

	// Tworzenie użytkownika
	user := models.User{
		Nickname: body.Nickname,
		Password: string(hash),
	}

	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Błąd podczas tworzenia użytkownika",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Użytkownik został zarejestrowany",
	})
}

func Login(c *gin.Context) {
	var body struct {
		Nickname string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Nieprawidłowe dane",
		})
		return
	}

	// Znajdź użytkownika
	var user models.User
	initializers.DB.First(&user, "nickname = ?", body.Nickname)

	if user.Id == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Nieprawidłowe dane logowania",
		})
		return
	}

	// Porównaj hasła
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Nieprawidłowe dane logowania",
		})
		return
	}

	// Generuj JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Błąd podczas generowania tokenu",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}
