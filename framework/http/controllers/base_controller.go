package controllers

import (
	"go-rails/framework/database"

	"github.com/gin-gonic/gin"
)

// BaseController содержит общие методы для всех контроллеров
type BaseController struct {
	DB *database.Database
}

// NewBaseController создает новый базовый контроллер
func NewBaseController(db *database.Database) *BaseController {
	return &BaseController{DB: db}
}

// SuccessResponse возвращает успешный ответ
func (bc *BaseController) SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"success": true,
		"data":    data,
	})
}

// ErrorResponse возвращает ответ с ошибкой
func (bc *BaseController) ErrorResponse(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{
		"success": false,
		"error":   message,
	})
}

// ValidationError возвращает ошибку валидации
func (bc *BaseController) ValidationError(c *gin.Context, errors map[string]string) {
	c.JSON(422, gin.H{
		"success": false,
		"errors":  errors,
	})
}

// NotFound возвращает ошибку 404
func (bc *BaseController) NotFound(c *gin.Context, message string) {
	if message == "" {
		message = "Resource not found"
	}
	bc.ErrorResponse(c, 404, message)
}

// Unauthorized возвращает ошибку 401
func (bc *BaseController) Unauthorized(c *gin.Context, message string) {
	if message == "" {
		message = "Unauthorized"
	}
	bc.ErrorResponse(c, 401, message)
}

// Forbidden возвращает ошибку 403
func (bc *BaseController) Forbidden(c *gin.Context, message string) {
	if message == "" {
		message = "Forbidden"
	}
	bc.ErrorResponse(c, 403, message)
}
