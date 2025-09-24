package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Инициализация загрузки файла окружения
func Init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file")
		return
	}
	log.Println(".env file loaded")
}

// Работа с дефолтными значениями конфигов
func getString(key, defaultValue string) string {
	// Читаем из переменной окружения
	value := os.Getenv(key)
	// Если значение пустое
	if value == "" {
		// Возвращаем дефолтное значение
		return defaultValue
	}
	// Если значение не пустое - возвращаем значение переменной окружения
	return value
}

// Работа с числовыми значениями
func getInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	// Конвертируем значение в Int
	i, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return i
}

// Работа с булево значениями
func getBool(key string, defaultValue bool) bool {
	value := os.Getenv(key)
	b, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}
	return b
}

// Структура конфигурации БД
type DbConfig struct {
	Url string
}

// Извлечение конфига БД из переменной окружения
func NewDatabaseConfig() *DbConfig {
	return &DbConfig{
		Url: getString("DATABASE_URL", ""),
	}
}
