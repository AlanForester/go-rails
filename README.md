# Go-Rails Framework

🚀 **Фреймворк на Go с функциональностью Ruby on Rails**

Go-Rails - это веб-фреймворк для Go, который предоставляет Rails-подобный опыт разработки с MVC архитектурой, ORM, роутингом и CLI генераторами.

## ✨ Возможности

- **🎯 MVC Архитектура** - Model, View, Controller
- **🗄️ ORM с GORM** - Object-Relational Mapping
- **🛣️ RESTful Роутинг** - автоматические маршруты
- **🔧 Middleware** - логирование, CORS, аутентификация
- **📊 Миграции БД** - управление схемой базы данных
- **⚡ CLI Генераторы** - создание компонентов
- **✅ Валидация** - встроенная валидация моделей
- **🔐 Аутентификация** - базовая система аутентификации
- **📝 Логирование** - структурированное логирование
- **⚙️ Конфигурация** - гибкая система настроек

## 🚀 Быстрый старт

### Установка

```bash
# Клонируйте репозиторий
git clone https://github.com/AlanForester/go-rails
cd go-rails

# Установите зависимости
go mod tidy
```

### Создание нового приложения

```bash
# Создайте новое приложение
go run cmd/gorails/main.go new myapp

# Перейдите в папку приложения
cd myapp

# Установите зависимости
go mod tidy

# Запустите сервер
go run main.go
```

Сервер будет доступен по адресу: http://localhost:3000

## 📖 CLI Команды

### Основные команды

```bash
# Создание нового приложения
go run cmd/gorails/main.go new [app_name]

# Запуск сервера
go run cmd/gorails/main.go server

# Показать справку
go run cmd/gorails/main.go --help
```

### Генерация компонентов

```bash
# Генерация контроллера
go run cmd/gorails/main.go generate controller users

# Генерация модели
go run cmd/gorails/main.go generate model user name:string email:string password:string

# Генерация миграции
go run cmd/gorails/main.go generate migration create_users_table
```

### Работа с базой данных

```bash
# Запуск миграций
go run cmd/gorails/main.go db migrate

# Заполнение тестовыми данными
go run cmd/gorails/main.go db seed
```

## 🛠️ Использование Makefile

Для удобства разработки используйте команды Makefile:

```bash
# Показать все доступные команды
make help

# Установить зависимости
make deps

# Запустить сервер
make run

# Создать новое приложение
make new-app NAME=myapp

# Генерировать контроллер
make generate-controller NAME=posts

# Генерировать модель
make generate-model NAME=post FIELDS="title:string content:text user_id:int"

# Запустить миграции
make db-migrate
```

## 📁 Структура проекта

```
go-rails/
├── cmd/gorails/              # CLI приложение
├── framework/                # Основной фреймворк
│   ├── core/                # Ядро приложения
│   ├── database/            # Работа с БД
│   ├── http/                # HTTP компоненты
│   │   ├── controllers/     # Контроллеры
│   │   └── router/          # Роутинг
│   ├── middleware/          # Middleware
│   ├── models/              # Модели
│   └── generators/          # Генераторы
├── config/                  # Конфигурация
├── examples/                # Примеры приложений
├── docs/                    # Документация
├── templates/               # Шаблоны для генераторов
├── Makefile                 # Команды для разработки
└── README.md               # Документация
```

## 🏗️ Структура приложения

```
myapp/
├── app/
│   ├── controllers/         # Контроллеры
│   ├── models/             # Модели
│   └── views/              # Представления
├── config/
│   └── config.yaml         # Конфигурация
├── db/
│   └── migrate/            # Миграции
├── public/
│   └── assets/             # Статические файлы
├── routes/
│   └── routes.go           # Маршруты
├── main.go                 # Точка входа
└── go.mod                  # Зависимости
```

## 💻 Примеры кода

### Контроллер

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

### Модель

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

### Роутинг

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

## ⚙️ Конфигурация

Отредактируйте `config/config.yaml`:

```yaml
server:
  port: 3000
  host: localhost

database:
  driver: sqlite3      # sqlite3, mysql, postgres
  database: app.db     # имя файла БД или имя БД
  host: localhost      # для MySQL/PostgreSQL
  port: 3306          # для MySQL/PostgreSQL
  username: root      # для MySQL/PostgreSQL
  password: password  # для MySQL/PostgreSQL
```

## 🔌 API Endpoints

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

## 🔧 Middleware

Фреймворк включает встроенные middleware:

- **Logger** - логирование запросов
- **CORS** - поддержка CORS
- **Recovery** - восстановление после паники
- **Auth** - аутентификация (базовая реализация)

## 📊 Поддерживаемые базы данных

- **SQLite3** (по умолчанию)
- **MySQL**
- **PostgreSQL**

## 🧪 Тестирование

```bash
# Запуск тестов
go test ./...

# Запуск тестов с покрытием
go test -cover ./...
```

## 📚 Документация

Подробная документация доступна в папке `docs/`:

- [Руководство по использованию](docs/USAGE.md)

## 🤝 Вклад в проект

1. Форкните репозиторий
2. Создайте ветку для новой функции (`git checkout -b feature/amazing-feature`)
3. Зафиксируйте изменения (`git commit -m 'Add amazing feature'`)
4. Отправьте в ветку (`git push origin feature/amazing-feature`)
5. Откройте Pull Request

## 📄 Лицензия

Этот проект распространяется под лицензией MIT. См. файл `LICENSE` для получения дополнительной информации.

## 🙏 Благодарности

- [Gin](https://github.com/gin-gonic/gin) - HTTP веб-фреймворк
- [GORM](https://gorm.io/) - ORM библиотека
- [Viper](https://github.com/spf13/viper) - конфигурация
- [Cobra](https://github.com/spf13/cobra) - CLI фреймворк

---

**Go-Rails Framework** - Создавайте веб-приложения на Go с легкостью Rails! 🚀 