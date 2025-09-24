package main

import "github.com/gofiber/fiber/v2"

func main() {
	// Создание инстанса приложения Fiber
	app := fiber.New()
	//Настраиваем порт, который будем слушать
	app.Listen(":3000")
}
