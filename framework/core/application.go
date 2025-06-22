package core

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"go-rails/framework/database"
	"go-rails/framework/http/router"
	"go-rails/framework/middleware"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Application представляет основное приложение
type Application struct {
	Router   *gin.Engine
	DB       *database.Database
	Config   *viper.Viper
	RootPath string
	Env      string
}

// NewApplication создает новое приложение
func NewApplication() *Application {
	app := &Application{
		Router:   gin.Default(),
		Config:   viper.New(),
		RootPath: getRootPath(),
		Env:      getEnvironment(),
	}

	app.setupConfig()
	app.setupDatabase()
	app.setupMiddleware()
	app.setupRoutes()

	return app
}

// setupConfig настраивает конфигурацию
func (app *Application) setupConfig() {
	app.Config.SetConfigName("config")
	app.Config.SetConfigType("yaml")
	app.Config.AddConfigPath(app.RootPath)
	app.Config.AddConfigPath(filepath.Join(app.RootPath, "config"))

	// Значения по умолчанию
	app.Config.SetDefault("server.port", 3000)
	app.Config.SetDefault("server.host", "localhost")
	app.Config.SetDefault("database.driver", "sqlite3")
	app.Config.SetDefault("database.database", "app.db")

	if err := app.Config.ReadInConfig(); err != nil {
		log.Printf("Warning: Could not read config file: %v", err)
	}
}

// setupDatabase настраивает базу данных
func (app *Application) setupDatabase() {
	dbConfig := database.Config{
		Driver:   app.Config.GetString("database.driver"),
		Host:     app.Config.GetString("database.host"),
		Port:     app.Config.GetString("database.port"),
		Database: app.Config.GetString("database.database"),
		Username: app.Config.GetString("database.username"),
		Password: app.Config.GetString("database.password"),
	}

	var err error
	app.DB, err = database.NewDatabase(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
}

// setupMiddleware настраивает middleware
func (app *Application) setupMiddleware() {
	// Логирование
	app.Router.Use(middleware.Logger())

	// CORS
	app.Router.Use(middleware.CORS())

	// Recovery
	app.Router.Use(middleware.Recovery())

	// Статические файлы
	app.Router.Static("/assets", filepath.Join(app.RootPath, "public", "assets"))
	app.Router.StaticFile("/favicon.ico", filepath.Join(app.RootPath, "public", "favicon.ico"))
}

// setupRoutes настраивает маршруты
func (app *Application) setupRoutes() {
	router.SetupRoutes(app.Router, app.DB)
}

// Run запускает приложение
func (app *Application) Run() error {
	port := app.Config.GetInt("server.port")
	host := app.Config.GetString("server.host")

	addr := fmt.Sprintf("%s:%d", host, port)
	log.Printf("Starting server on %s", addr)

	return app.Router.Run(addr)
}

// getRootPath возвращает корневую папку приложения
func getRootPath() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

// getEnvironment возвращает текущее окружение
func getEnvironment() string {
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "development"
	}
	return env
}
