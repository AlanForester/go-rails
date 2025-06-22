package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

// Config содержит конфигурацию базы данных
type Config struct {
	Driver   string
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

// Database представляет соединение с базой данных
type Database struct {
	*gorm.DB
}

// NewDatabase создает новое соединение с базой данных
func NewDatabase(config Config) (*Database, error) {
	var dsn string

	switch config.Driver {
	case "sqlite3":
		dsn = config.Database
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			config.Username, config.Password, config.Host, config.Port, config.Database)
	case "postgres":
		dsn = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			config.Host, config.Port, config.Username, config.Database, config.Password)
	default:
		return nil, fmt.Errorf("unsupported database driver: %s", config.Driver)
	}

	db, err := gorm.Open(config.Driver, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// Настройка GORM
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	return &Database{DB: db}, nil
}

// AutoMigrate выполняет автоматическую миграцию моделей
func (db *Database) AutoMigrate(models ...interface{}) error {
	return db.DB.AutoMigrate(models...).Error
}

// CreateTable создает таблицу для модели
func (db *Database) CreateTable(model interface{}) error {
	return db.DB.CreateTable(model).Error
}

// DropTable удаляет таблицу
func (db *Database) DropTable(model interface{}) error {
	return db.DB.DropTable(model).Error
}

// HasTable проверяет существование таблицы
func (db *Database) HasTable(model interface{}) bool {
	return db.DB.HasTable(model)
}

// Методы для работы с GORM (делегирование к встроенному DB)

// Find находит записи
func (db *Database) Find(out interface{}, where ...interface{}) *gorm.DB {
	return db.DB.Find(out, where...)
}

// First находит первую запись
func (db *Database) First(out interface{}, where ...interface{}) *gorm.DB {
	return db.DB.First(out, where...)
}

// Create создает запись
func (db *Database) Create(value interface{}) *gorm.DB {
	return db.DB.Create(value)
}

// Save сохраняет запись
func (db *Database) Save(value interface{}) *gorm.DB {
	return db.DB.Save(value)
}

// Delete удаляет запись
func (db *Database) Delete(value interface{}, where ...interface{}) *gorm.DB {
	return db.DB.Delete(value, where...)
}

// Where добавляет условие WHERE
func (db *Database) Where(query interface{}, args ...interface{}) *gorm.DB {
	return db.DB.Where(query, args...)
}
