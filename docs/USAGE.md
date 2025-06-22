# Go-Rails Framework - Руководство по использованию

## Установка и настройка

### 1. Установка зависимостей

```bash
go mod tidy
```

### 2. Создание нового приложения

```bash
go run cmd/gorails/main.go new myapp
cd myapp
go mod tidy
```

### 3. Запуск сервера

```bash
go run main.go
```

Сервер будет доступен по адресу http://localhost:3000

## CLI Команды

### Создание нового приложения
```bash
go run cmd/gorails/main.go new [app_name]
```

### Запуск сервера
```bash
go run cmd/gorails/main.go server
```

### Генерация компонентов

#### Контроллер
```bash
go run cmd/gorails/main.go generate controller users
```

#### Модель
```bash
go run cmd/gorails/main.go generate model user name:string email:string password:string
```

#### Миграция
```bash
go run cmd/gorails/main.go generate migration create_users_table
```

### Работа с базой данных

#### Миграции
```bash
go run cmd/gorails/main.go db migrate
```

#### Заполнение тестовыми данными
```bash
go run cmd/gorails/main.go db seed
```

## Структура приложения

```
myapp/
├── app/
│   ├── controllers/     # Контроллеры
│   ├── models/         # Модели
│   └── views/          # Представления
├── config/
│   └── config.yaml     # Конфигурация
├── db/
│   └── migrate/        # Миграции
├── public/
│   └── assets/         # Статические файлы
├── routes/
│   └── routes.go       # Маршруты
├── main.go             # Точка входа
└── go.mod              # Зависимости
```

## Создание контроллера

```go
package controllers

import (
    "github.com/gin-gonic/gin"
)

type UsersController struct{}

func (uc *UsersController) Index(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "List of users",
    })
}

func (uc *UsersController) Show(c *gin.Context) {
    id := c.Param("id")
    c.JSON(200, gin.H{
        "message": "User with ID: " + id,
    })
}
```

## Создание модели

```go
package models

import (
    "time"
    "github.com/jinzhu/gorm"
)

type User struct {
    ID        uint      `json:"id" gorm:"primary_key"`
    Name      string    `json:"name" gorm:"not null"`
    Email     string    `json:"email" gorm:"unique;not null"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

func (User) TableName() string {
    return "users"
}
```

## Настройка маршрутов

```go
package routes

import (
    "github.com/gin-gonic/gin"
    "myapp/app/controllers"
)

func SetupRoutes(r *gin.Engine) {
    usersController := &controllers.UsersController{}
    
    r.GET("/users", usersController.Index)
    r.GET("/users/:id", usersController.Show)
    r.POST("/users", usersController.Create)
    r.PUT("/users/:id", usersController.Update)
    r.DELETE("/users/:id", usersController.Destroy)
}
```

## Конфигурация базы данных

Отредактируйте файл `config/config.yaml`:

```yaml
database:
  driver: sqlite3      # sqlite3, mysql, postgres
  database: app.db     # имя файла БД или имя БД
  host: localhost      # для MySQL/PostgreSQL
  port: 3306          # для MySQL/PostgreSQL
  username: root      # для MySQL/PostgreSQL
  password: password  # для MySQL/PostgreSQL
```

## API Endpoints

### Пользователи
- `GET /api/v1/users` - список пользователей
- `GET /api/v1/users/:id` - получить пользователя
- `POST /api/v1/users` - создать пользователя
- `PUT /api/v1/users/:id` - обновить пользователя
- `DELETE /api/v1/users/:id` - удалить пользователя

### Аутентификация
- `POST /api/v1/login` - вход
- `POST /api/v1/register` - регистрация
- `POST /api/v1/logout` - выход

## Middleware

Фреймворк включает встроенные middleware:

- **Logger** - логирование запросов
- **CORS** - поддержка CORS
- **Recovery** - восстановление после паники
- **Auth** - аутентификация (базовая реализация)

## Валидация

Модели поддерживают встроенную валидацию:

```go
func (u *User) Validate() map[string]string {
    errors := make(map[string]string)
    
    if u.Name == "" {
        errors["name"] = "Name is required"
    }
    
    if u.Email == "" {
        errors["email"] = "Email is required"
    }
    
    return errors
}
```

## Примеры

Смотрите папку `examples/` для примеров использования фреймворка. 