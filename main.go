package main

import (
	"1/controller"
	"1/initializers"
	"1/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	//
	initializers.LoadEnv()
	initializers.ConnectDB()
	initializers.CreateDb()
}

func main() {
	r := gin.Default()

	// Obsługa CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})
	// W funkcji main() dodaj:
	r.Static("/uploads", "./uploads")
	r.Static("/frontend", "./frontend")

	// Publiczne endpointy
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)
	r.GET("/images/public", controller.GetAllPublicImages)

	// Chronione endpointy (wymagają JWT)
	authorized := r.Group("/")
	authorized.Use(middleware.RequireAuth)
	{
		authorized.POST("/images/upload", controller.UploadImage)
	}

	r.Run(":8080")
}
