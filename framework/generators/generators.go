package generators

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"go-rails/framework/database"
)

// CreateNewApp создает новое приложение
func CreateNewApp(appName string) error {
	// Создаем структуру папок
	dirs := []string{
		appName,
		filepath.Join(appName, "app"),
		filepath.Join(appName, "app", "controllers"),
		filepath.Join(appName, "app", "models"),
		filepath.Join(appName, "app", "views"),
		filepath.Join(appName, "config"),
		filepath.Join(appName, "db"),
		filepath.Join(appName, "db", "migrate"),
		filepath.Join(appName, "public"),
		filepath.Join(appName, "public", "assets"),
		filepath.Join(appName, "routes"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	// Создаем основные файлы
	files := map[string]string{
		filepath.Join(appName, "go.mod"):                                          generateGoMod(appName),
		filepath.Join(appName, "main.go"):                                         generateMainGo(appName),
		filepath.Join(appName, "config", "config.yaml"):                           generateConfig(),
		filepath.Join(appName, "routes", "routes.go"):                             generateRoutes(),
		filepath.Join(appName, "README.md"):                                       generateReadme(appName),
		filepath.Join(appName, ".gitignore"):                                      generateGitignore(),
		filepath.Join(appName, "app", "controllers", "application_controller.go"): generateApplicationController(),
	}

	for path, content := range files {
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			return err
		}
	}

	return nil
}

// GenerateController генерирует новый контроллер
func GenerateController(controllerName string) error {
	controllerContent := generateControllerContent(controllerName)

	// Создаем папку если её нет
	controllerDir := filepath.Join("app", "controllers")
	if err := os.MkdirAll(controllerDir, 0755); err != nil {
		return err
	}

	controllerPath := filepath.Join(controllerDir, strings.ToLower(controllerName)+"_controller.go")

	return os.WriteFile(controllerPath, []byte(controllerContent), 0644)
}

// GenerateModel генерирует новую модель
func GenerateModel(modelName string, fields []string) error {
	modelContent := generateModelContent(modelName, fields)

	// Создаем папку если её нет
	modelDir := filepath.Join("app", "models")
	if err := os.MkdirAll(modelDir, 0755); err != nil {
		return err
	}

	modelPath := filepath.Join(modelDir, strings.ToLower(modelName)+".go")

	return os.WriteFile(modelPath, []byte(modelContent), 0644)
}

// GenerateMigration генерирует новую миграцию
func GenerateMigration(migrationName string) error {
	timestamp := time.Now().Format("20060102150405")
	migrationContent := generateMigrationContent(migrationName)

	// Создаем папку если её нет
	migrationDir := filepath.Join("db", "migrate")
	if err := os.MkdirAll(migrationDir, 0755); err != nil {
		return err
	}

	migrationPath := filepath.Join(migrationDir, timestamp+"_"+strings.ToLower(migrationName)+".go")

	return os.WriteFile(migrationPath, []byte(migrationContent), 0644)
}

// SeedDatabase заполняет базу данных тестовыми данными
func SeedDatabase(db *database.Database) error {
	// Здесь должна быть логика заполнения БД
	return nil
}

// Вспомогательные функции для генерации содержимого файлов

func generateGoMod(appName string) string {
	return fmt.Sprintf(`module %s

go 1.21

require (
	github.com/gin-gonic/gin v1.9.1
	github.com/jinzhu/gorm v1.9.16
	github.com/mattn/go-sqlite3 v1.14.17
	github.com/spf13/viper v1.16.0
	golang.org/x/crypto v0.14.0
)`, appName)
}

func generateMainGo(appName string) string {
	return fmt.Sprintf(`package main

import (
	"log"
	"%s/framework/core"
)

func main() {
	app := core.NewApplication()
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}`, appName)
}

func generateConfig() string {
	return `server:
  port: 3000
  host: localhost

database:
  driver: sqlite3
  database: app.db
`
}

func generateRoutes() string {
	return `package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Go-Rails!",
		})
	})
}
`
}

func generateReadme(appName string) string {
	return fmt.Sprintf(`# %s

A Go-Rails application.

## Getting Started

1. Install dependencies:
   `+"`"+`bash
   go mod tidy
   `+"`"+`

2. Run the server:
   `+"`"+`bash
   go run main.go
   `+"`"+`

3. Visit http://localhost:3000
`, appName)
}

func generateGitignore() string {
	return `# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with 'go test -c'
*.test

# Output of the go coverage tool, specifically when used with LiteIDE
*.out

# Dependency directories (remove the comment below to include it)
# vendor/

# Go workspace file
go.work

# Database
*.db
*.sqlite

# Environment variables
.env

# IDE
.vscode/
.idea/
`
}

func generateApplicationController() string {
	return `package controllers

import (
	"github.com/gin-gonic/gin"
)

// ApplicationController базовый контроллер
type ApplicationController struct{}

// Index главная страница
func (ac *ApplicationController) Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to Go-Rails!",
	})
}
`
}

func generateControllerContent(controllerName string) string {
	return fmt.Sprintf(`package controllers

import (
	"github.com/gin-gonic/gin"
)

// %sController контроллер для %s
type %sController struct{}

// Index возвращает список всех %s
func (cc *%sController) Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Index %s",
	})
}

// Show возвращает конкретный %s
func (cc *%sController) Show(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"message": "Show %s with ID: " + id,
	})
}

// Create создает новый %s
func (cc *%sController) Create(c *gin.Context) {
	c.JSON(201, gin.H{
		"message": "Create %s",
	})
}

// Update обновляет %s
func (cc *%sController) Update(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"message": "Update %s with ID: " + id,
	})
}

// Destroy удаляет %s
func (cc *%sController) Destroy(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"message": "Destroy %s with ID: " + id,
	})
}
`,
		strings.Title(controllerName),
		controllerName,
		strings.Title(controllerName),
		controllerName,
		strings.Title(controllerName),
		controllerName,
		controllerName,
		strings.Title(controllerName),
		controllerName,
		strings.Title(controllerName),
		controllerName,
		strings.Title(controllerName),
		controllerName,
		strings.Title(controllerName),
		controllerName,
	)
}

func generateModelContent(modelName string, fields []string) string {
	fieldDeclarations := ""
	for _, field := range fields {
		parts := strings.Split(field, ":")
		if len(parts) == 2 {
			fieldName := parts[0]
			fieldType := parts[1]
			fieldDeclarations += fmt.Sprintf("\t%s %s `json:\"%s\" gorm:\"not null\"`\n",
				strings.Title(fieldName),
				getGoType(fieldType),
				fieldName)
		}
	}

	return fmt.Sprintf(`package models

import (
	"time"
	"github.com/jinzhu/gorm"
)

// %s представляет модель %s
type %s struct {
	ID        uint      `+"`json:\"id\" gorm:\"primary_key\"`"+`
%s	CreatedAt time.Time `+"`json:\"created_at\"`"+`
	UpdatedAt time.Time `+"`json:\"updated_at\"`"+`
}

// TableName возвращает имя таблицы
func (%s) TableName() string {
	return "%s"
}

// BeforeCreate выполняется перед созданием записи
func (m *%s) BeforeCreate(scope *gorm.Scope) error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate выполняется перед обновлением записи
func (m *%s) BeforeUpdate(scope *gorm.Scope) error {
	m.UpdatedAt = time.Now()
	return nil
}
`,
		strings.Title(modelName),
		modelName,
		strings.Title(modelName),
		fieldDeclarations,
		strings.Title(modelName),
		strings.ToLower(modelName)+"s",
		strings.Title(modelName),
		strings.Title(modelName),
	)
}

func generateMigrationContent(migrationName string) string {
	return fmt.Sprintf(`package migrate

import (
	"github.com/jinzhu/gorm"
)

// %s миграция для %s
func %s(db *gorm.DB) error {
	// TODO: Implement migration logic
	return nil
}
`,
		strings.Title(migrationName),
		migrationName,
		strings.Title(migrationName),
	)
}

func getGoType(dbType string) string {
	switch dbType {
	case "string":
		return "string"
	case "text":
		return "string"
	case "integer":
		return "int"
	case "int":
		return "int"
	case "bigint":
		return "int64"
	case "float":
		return "float64"
	case "decimal":
		return "float64"
	case "boolean":
		return "bool"
	case "bool":
		return "bool"
	case "datetime":
		return "time.Time"
	case "timestamp":
		return "time.Time"
	case "date":
		return "time.Time"
	default:
		return "string"
	}
}
