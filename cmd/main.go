package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/viacheslav-korobeynikov/rally-app/config"
)

func main() {
	// Инициализация чтения файла конфигурации
	config.Init()
	// Получения конфигурации подключения к БД
	config.NewDatabaseConfig()
	// Создание инстанса приложения Fiber
	app := fiber.New()
	//Настраиваем порт, который будем слушать
	app.Listen(":3000")
}
