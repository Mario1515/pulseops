package handler

import (
	"pulseops/internal/middleware"
	"pulseops/internal/model"
	"pulseops/internal/service"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

type UserHandler struct {
	svc *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

func (h *UserHandler) RegisterRoutes(app *fiber.App, apiKey string) {
	v1 := app.Group("/api/v1")
	users := v1.Group("/users", middleware.APIKey(apiKey))

	users.Get("", h.list)
	users.Post("", h.create)
	users.Delete("/:id", h.delete)
}

func (h *UserHandler) list(c fiber.Ctx) error {
	users, err := h.svc.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(users)
}

func (h *UserHandler) create(c fiber.Ctx) error {
	var u model.User
	if err := c.Bind().Body(&u); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if errors := validateStruct(u); len(errors) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"errors": errors})
	}
	if err := h.svc.Create(&u); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(u)
}

func (h *UserHandler) delete(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	if err := h.svc.Delete(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
