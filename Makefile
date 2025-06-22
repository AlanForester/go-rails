.PHONY: help build run test clean deps generate new-app

# Переменные
BINARY_NAME=gorails
MAIN_PATH=cmd/gorails/main.go

# Цвета для вывода
GREEN=\033[0;32m
NC=\033[0m # No Color

help: ## Показать справку
	@echo "Доступные команды:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "$(GREEN)%-20s$(NC) %s\n", $$1, $$2}'

deps: ## Установить зависимости
	go mod tidy
	go mod download

build: ## Собрать приложение
	go build -o bin/$(BINARY_NAME) $(MAIN_PATH)

run: ## Запустить сервер разработки
	go run $(MAIN_PATH) server

test: ## Запустить тесты
	go test ./...

clean: ## Очистить собранные файлы
	rm -rf bin/
	go clean

generate-controller: ## Генерировать контроллер (использование: make generate-controller NAME=users)
	@if [ -z "$(NAME)" ]; then echo "Укажите имя контроллера: make generate-controller NAME=users"; exit 1; fi
	go run $(MAIN_PATH) generate controller $(NAME)

generate-model: ## Генерировать модель (использование: make generate-model NAME=user FIELDS="name:string email:string")
	@if [ -z "$(NAME)" ]; then echo "Укажите имя модели: make generate-model NAME=user FIELDS='name:string email:string'"; exit 1; fi
	go run $(MAIN_PATH) generate model $(NAME) $(FIELDS)

generate-migration: ## Генерировать миграцию (использование: make generate-migration NAME=create_users)
	@if [ -z "$(NAME)" ]; then echo "Укажите имя миграции: make generate-migration NAME=create_users"; exit 1; fi
	go run $(MAIN_PATH) generate migration $(NAME)

new-app: ## Создать новое приложение (использование: make new-app NAME=myapp)
	@if [ -z "$(NAME)" ]; then echo "Укажите имя приложения: make new-app NAME=myapp"; exit 1; fi
	go run $(MAIN_PATH) new $(NAME)

db-migrate: ## Запустить миграции базы данных
	go run $(MAIN_PATH) db migrate

db-seed: ## Заполнить базу данных тестовыми данными
	go run $(MAIN_PATH) db seed

dev: ## Запустить в режиме разработки с автоперезагрузкой
	@echo "Запуск в режиме разработки..."
	@echo "Для автоперезагрузки установите air: go install github.com/cosmtrek/air@latest"
	air

install-air: ## Установить air для автоперезагрузки
	go install github.com/cosmtrek/air@latest

lint: ## Запустить линтер
	golangci-lint run

fmt: ## Форматировать код
	go fmt ./...

vet: ## Проверить код
	go vet ./...

# Команды для демонстрации
demo: ## Запустить демо приложение
	@echo "Запуск демо приложения..."
	cd examples/blog_app && go run main.go

demo-simple: ## Запустить простое демо
	@echo "Запуск простого демо..."
	cd examples/simple_app && go run main.go 