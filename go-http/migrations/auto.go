package main

import (
	"app/adv-http/internal/link"
	"app/adv-http/internal/stat"
	"app/adv-http/internal/user"
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Пытаемся загрузить .env файл из разных возможных мест
	// Пробуем в следующем порядке:
	// 1. Текущая директория (где запущена программа)
	// 2. Родительская директория (корень проекта, где go.mod)
	var err error
	if err = godotenv.Load(); err != nil {
		// Пробуем загрузить из родительской директории (корень проекта)
		parentEnv := filepath.Join("..", ".env")
		if err = godotenv.Load(parentEnv); err != nil {
			fmt.Printf("Warning: .env file not found, trying to use system environment variables: %v\n", err)
		}
	}

	// Получаем DSN из переменных окружения
	dsn := os.Getenv("DSN")
	if dsn == "" {
		panic("DSN environment variable is not set. Please set it in .env file or environment variables")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	// Выполняем миграции
	err = db.AutoMigrate(&link.Link{}, &user.User{}, &stat.Stat{})
	if err != nil {
		panic(fmt.Sprintf("Failed to run migrations: %v", err))
	}

	fmt.Println("Migrations completed successfully!")
}
