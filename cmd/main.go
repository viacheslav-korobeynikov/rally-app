package main

import (
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/viacheslav-korobeynikov/rally-app/config"
	"github.com/viacheslav-korobeynikov/rally-app/internal/home"
	"github.com/viacheslav-korobeynikov/rally-app/pkg/logger"
)

func main() {
	// Инициализация чтения файла конфигурации
	config.Init()
	// Получения конфигурации подключения к БД
	config.NewDatabaseConfig()
	// Получение конфигурации логирования
	logConfig := config.NewLogConfig()
	customLogger := logger.NewLogger(logConfig)
	// Создание инстанса приложения Fiber
	app := fiber.New()
	// Middleware для логирования запросов
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: customLogger,
	}))
	// Добавлена зависимость с хендлером для главной страницы
	home.NewHandler(app, customLogger)
	//Настраиваем порт, который будем слушать
	app.Listen(":3000")
}
