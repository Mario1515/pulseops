package handler

import (
	"pulseops/internal/middleware"
	"pulseops/internal/model"
	"pulseops/internal/service"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v3"
)

type IncidentHandler struct {
	svc *service.IncidentService
}

func NewIncidentHandler(svc *service.IncidentService) *IncidentHandler {
	return &IncidentHandler{svc: svc}
}

func (h *IncidentHandler) RegisterRoutes(app *fiber.App, apiKey string) {
	v1 := app.Group("/api/v1")
	incidents := v1.Group("/incidents", middleware.APIKey(apiKey))

	incidents.Get("", middleware.RateLimiter(10, 1*time.Minute), h.list)
	incidents.Post("", middleware.RateLimiter(10, 1*time.Minute), h.create)
	incidents.Delete("/:id", middleware.RateLimiter(20, 1*time.Minute), h.delete)
}

func (h *IncidentHandler) list(c fiber.Ctx) error {
	incidents, err := h.svc.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(incidents)
}

func (h *IncidentHandler) create(c fiber.Ctx) error {
	var incident model.Incident
	if err := c.Bind().Body(&incident); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if errors := validateStruct(incident); len(errors) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"errors": errors})
	}
	if err := h.svc.Create(&incident); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(incident)
}

func (h *IncidentHandler) delete(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	if err := h.svc.Delete(uint(id)); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
