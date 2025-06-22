package router

import (
	"go-rails/framework/database"
	"go-rails/framework/http/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes настраивает все маршруты приложения
func SetupRoutes(r *gin.Engine, db *database.Database) {
	// Главная страница
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Go-Rails Framework!",
			"version": "1.0.0",
		})
	})

	// API маршруты
	api := r.Group("/api/v1")
	{
		// Пользователи
		usersController := controllers.NewUsersController(db)
		api.GET("/users", usersController.Index)
		api.GET("/users/:id", usersController.Show)
		api.POST("/users", usersController.Create)
		api.PUT("/users/:id", usersController.Update)
		api.DELETE("/users/:id", usersController.Destroy)

		// Аутентификация
		authController := controllers.NewAuthController(db)
		api.POST("/login", authController.Login)
		api.POST("/register", authController.Register)
		api.POST("/logout", authController.Logout)
	}

	// Статические файлы
	r.Static("/assets", "./public/assets")
	r.StaticFile("/favicon.ico", "./public/favicon.ico")
}
