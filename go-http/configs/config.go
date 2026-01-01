package configs

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	Db   DbConfig
	Auth AuthConfig
}

type DbConfig struct {
	Dsn string
}

type AuthConfig struct {
	Secret string
}

func LoadConfig() *Config {
	// Пытаемся загрузить .env файл из разных возможных мест
	// Пробуем в следующем порядке:
	// 1. Текущая директория (где запущена программа)
	// 2. Родительская директория (корень проекта, где go.mod)
	var err error

	// Пробуем загрузить из текущей директории
	if err = godotenv.Load(); err != nil {
		// Пробуем загрузить из родительской директории (корень проекта)
		// Это нужно когда запускаем из cmd/ или других поддиректорий
		parentEnv := filepath.Join("..", ".env")
		if err = godotenv.Load(parentEnv); err != nil {
			// Это не критичная ошибка - можно использовать переменные окружения системы
			log.Printf("Warning: .env file not found in current or parent directory, using system environment variables: %v", err)
		}
	}

	// Всегда возвращаем валидный Config, используя переменные окружения
	// Если переменные не заданы, будут пустые строки
	return &Config{
		Db: DbConfig{
			Dsn: getEnv("DB_DSN", ""),
		},
		Auth: AuthConfig{
			Secret: getEnv("TOKEN", ""),
		},
	}
}

// getEnv получает переменную окружения или возвращает значение по умолчанию
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
