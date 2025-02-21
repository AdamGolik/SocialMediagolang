package controller

import (
	"1/initializers"
	"1/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

func GetAllPublicImages(c *gin.Context) {
	var images []models.Image
	initializers.DB.Where("public = ?", true).Find(&images)

	c.JSON(http.StatusOK, images)
}

func UploadImage(c *gin.Context) {
	// Pobierz ID użytkownika z kontekstu (ustawione przez middleware)
	userId := c.GetUint("userId")

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Nie znaleziono pliku",
		})
		return
	}

	// Generuj unikalną nazwę pliku
	filename := filepath.Base(file.Filename)

	// Zapisz plik
	if err := c.SaveUploadedFile(file, "uploads/"+filename); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Błąd podczas zapisywania pliku",
		})
		return
	}

	// Zapisz informacje o obrazie w bazie danych
	image := models.Image{
		Name:   filename,
		Public: c.DefaultPostForm("public", "false") == "true",
	}

	var user models.User
	initializers.DB.First(&user, userId)

	// Przypisz obraz do użytkownika
	user.Images = append(user.Images, image)
	initializers.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": "Plik został przesłany pomyślnie",
		"image":   image,
	})
}
