package home

import "github.com/gofiber/fiber/v2"

// Структура хэндлера домашней страницы
type HomeHandler struct {
	router fiber.Router
}

// Функция конструктор хэндлера
func NewHandler(router fiber.Router) {
	h := &HomeHandler{
		router: router,
	}
	// Роутинг
	h.router.Get("/", h.home) // При GET запросе вызывается функция home
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	return c.SendString("Hello")
}
