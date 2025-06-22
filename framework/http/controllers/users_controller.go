package controllers

import (
	"net/http"
	"strconv"

	"go-rails/framework/database"
	"go-rails/framework/models"

	"github.com/gin-gonic/gin"
)

// UsersController управляет пользователями
type UsersController struct {
	*BaseController
}

// NewUsersController создает новый контроллер пользователей
func NewUsersController(db *database.Database) *UsersController {
	return &UsersController{
		BaseController: NewBaseController(db),
	}
}

// Index возвращает список всех пользователей
func (uc *UsersController) Index(c *gin.Context) {
	var users []models.User

	if err := uc.DB.Find(&users).Error; err != nil {
		uc.ErrorResponse(c, 500, "Failed to fetch users")
		return
	}

	uc.SuccessResponse(c, users)
}

// Show возвращает конкретного пользователя
func (uc *UsersController) Show(c *gin.Context) {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		uc.ErrorResponse(c, 400, "Invalid user ID")
		return
	}

	var user models.User
	if err := uc.DB.First(&user, userID).Error; err != nil {
		uc.NotFound(c, "User not found")
		return
	}

	uc.SuccessResponse(c, user)
}

// Create создает нового пользователя
func (uc *UsersController) Create(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		uc.ErrorResponse(c, 400, "Invalid request data")
		return
	}

	// Валидация
	if errors := user.Validate(); len(errors) > 0 {
		uc.ValidationError(c, errors)
		return
	}

	// Хеширование пароля
	if err := user.HashPassword(); err != nil {
		uc.ErrorResponse(c, 500, "Failed to hash password")
		return
	}

	if err := uc.DB.Create(&user).Error; err != nil {
		uc.ErrorResponse(c, 500, "Failed to create user")
		return
	}

	uc.SuccessResponse(c, user)
}

// Update обновляет пользователя
func (uc *UsersController) Update(c *gin.Context) {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		uc.ErrorResponse(c, 400, "Invalid user ID")
		return
	}

	var user models.User
	if err := uc.DB.First(&user, userID).Error; err != nil {
		uc.NotFound(c, "User not found")
		return
	}

	var updateData models.User
	if err := c.ShouldBindJSON(&updateData); err != nil {
		uc.ErrorResponse(c, 400, "Invalid request data")
		return
	}

	// Обновляем только разрешенные поля
	if updateData.Name != "" {
		user.Name = updateData.Name
	}
	if updateData.Email != "" {
		user.Email = updateData.Email
	}

	if err := uc.DB.Save(&user).Error; err != nil {
		uc.ErrorResponse(c, 500, "Failed to update user")
		return
	}

	uc.SuccessResponse(c, user)
}

// Destroy удаляет пользователя
func (uc *UsersController) Destroy(c *gin.Context) {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		uc.ErrorResponse(c, 400, "Invalid user ID")
		return
	}

	var user models.User
	if err := uc.DB.First(&user, userID).Error; err != nil {
		uc.NotFound(c, "User not found")
		return
	}

	if err := uc.DB.Delete(&user).Error; err != nil {
		uc.ErrorResponse(c, 500, "Failed to delete user")
		return
	}

	c.Status(http.StatusNoContent)
}
