package home

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

// Структура хэндлера домашней страницы
type HomeHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
}

// Функция конструктор хэндлера
func NewHandler(router fiber.Router, customLogger *zerolog.Logger) {
	h := &HomeHandler{
		router:       router,
		customLogger: customLogger,
	}
	// Роутинг
	h.router.Get("/", h.home) // При GET запросе вызывается функция home
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	return c.SendString("Hello")
}
