package handler

import (
	"pulseops/internal/model"
	"pulseops/internal/service"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	svc *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

func (h *UserHandler) RegisterRoutes(app *fiber.App) {
	v1 := app.Group("/api/v1")
	users := v1.Group("/users")

	users.Get("", h.list)
	users.Post("", h.create)
}

func (h *UserHandler) list(c *fiber.Ctx) error {
	users, err := h.svc.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(users)
}

func (h *UserHandler) create(c *fiber.Ctx) error {
	var u model.User
	if err := c.BodyParser(&u); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.svc.Create(&u); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(u)
}
