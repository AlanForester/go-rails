package controllers

import (
	"go-rails/framework/database"
	"go-rails/framework/models"

	"github.com/gin-gonic/gin"
)

// AuthController управляет аутентификацией
type AuthController struct {
	*BaseController
}

// NewAuthController создает новый контроллер аутентификации
func NewAuthController(db *database.Database) *AuthController {
	return &AuthController{
		BaseController: NewBaseController(db),
	}
}

// Login обрабатывает вход пользователя
func (ac *AuthController) Login(c *gin.Context) {
	var loginData struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		ac.ErrorResponse(c, 400, "Invalid login data")
		return
	}

	var user models.User
	if err := ac.DB.Where("email = ?", loginData.Email).First(&user).Error; err != nil {
		ac.Unauthorized(c, "Invalid credentials")
		return
	}

	if !user.CheckPassword(loginData.Password) {
		ac.Unauthorized(c, "Invalid credentials")
		return
	}

	// Генерируем токен
	token, err := user.GenerateToken()
	if err != nil {
		ac.ErrorResponse(c, 500, "Failed to generate token")
		return
	}

	ac.SuccessResponse(c, gin.H{
		"token": token,
		"user":  user,
	})
}

// Register обрабатывает регистрацию пользователя
func (ac *AuthController) Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		ac.ErrorResponse(c, 400, "Invalid registration data")
		return
	}

	// Валидация
	if errors := user.Validate(); len(errors) > 0 {
		ac.ValidationError(c, errors)
		return
	}

	// Проверяем, что email уникален
	var existingUser models.User
	if err := ac.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		ac.ErrorResponse(c, 422, "Email already exists")
		return
	}

	// Хешируем пароль
	if err := user.HashPassword(); err != nil {
		ac.ErrorResponse(c, 500, "Failed to hash password")
		return
	}

	// Создаем пользователя
	if err := ac.DB.Create(&user).Error; err != nil {
		ac.ErrorResponse(c, 500, "Failed to create user")
		return
	}

	// Генерируем токен
	token, err := user.GenerateToken()
	if err != nil {
		ac.ErrorResponse(c, 500, "Failed to generate token")
		return
	}

	ac.SuccessResponse(c, gin.H{
		"token": token,
		"user":  user,
	})
}

// Logout обрабатывает выход пользователя
func (ac *AuthController) Logout(c *gin.Context) {
	// В простой реализации просто возвращаем успех
	// В реальном приложении здесь должна быть логика инвалидации токена
	ac.SuccessResponse(c, gin.H{
		"message": "Successfully logged out",
	})
}
