package models

import (
	"errors"
	"regexp"
	"time"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User представляет модель пользователя
type User struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Name      string    `json:"name" gorm:"not null"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Password  string    `json:"-" gorm:"not null"` // "-" скрывает поле из JSON
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName возвращает имя таблицы
func (User) TableName() string {
	return "users"
}

// BeforeCreate выполняется перед созданием записи
func (u *User) BeforeCreate(scope *gorm.Scope) error {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate выполняется перед обновлением записи
func (u *User) BeforeUpdate(scope *gorm.Scope) error {
	u.UpdatedAt = time.Now()
	return nil
}

// Validate выполняет валидацию модели
func (u *User) Validate() map[string]string {
	errors := make(map[string]string)

	// Валидация имени
	if u.Name == "" {
		errors["name"] = "Name is required"
	} else if len(u.Name) < 2 {
		errors["name"] = "Name must be at least 2 characters long"
	} else if len(u.Name) > 50 {
		errors["name"] = "Name must be less than 50 characters"
	}

	// Валидация email
	if u.Email == "" {
		errors["email"] = "Email is required"
	} else {
		emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
		if !emailRegex.MatchString(u.Email) {
			errors["email"] = "Invalid email format"
		}
	}

	// Валидация пароля
	if u.Password == "" {
		errors["password"] = "Password is required"
	} else if len(u.Password) < 6 {
		errors["password"] = "Password must be at least 6 characters long"
	}

	return errors
}

// HashPassword хеширует пароль
func (u *User) HashPassword() error {
	if u.Password == "" {
		return errors.New("password is empty")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	return nil
}

// CheckPassword проверяет пароль
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

// GenerateToken генерирует JWT токен
func (u *User) GenerateToken() (string, error) {
	// Простая реализация токена
	// В реальном приложении здесь должна быть JWT логика
	token := "token_" + u.Email + "_" + time.Now().Format("20060102150405")
	return token, nil
}

// ToJSON возвращает пользователя в формате JSON без пароля
func (u *User) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"id":         u.ID,
		"name":       u.Name,
		"email":      u.Email,
		"created_at": u.CreatedAt,
		"updated_at": u.UpdatedAt,
	}
}
