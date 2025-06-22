package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger возвращает middleware для логирования
func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	})
}

// CORS возвращает middleware для CORS
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// Recovery возвращает middleware для восстановления после паники
func Recovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			c.String(500, "Internal Server Error: "+err)
		}
		c.AbortWithStatus(500)
	})
}

// Auth возвращает middleware для аутентификации
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(401, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		// Здесь должна быть логика проверки токена
		// Пока что просто проверяем наличие заголовка
		c.Next()
	}
}

// RateLimit возвращает middleware для ограничения запросов
func RateLimit(limit int) gin.HandlerFunc {
	// Простая реализация rate limiting
	return func(c *gin.Context) {
		// Здесь должна быть логика rate limiting
		c.Next()
	}
}
