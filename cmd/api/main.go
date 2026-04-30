package main

import (
	"log"

	"pulseops/internal/config"
	"pulseops/internal/database"
	"pulseops/internal/handler"
	"pulseops/internal/middleware"
	"pulseops/internal/repository"
	"pulseops/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gorm.io/gorm"
)

func main() {
	cfg := config.Load()
	db := database.Connect(&cfg.DB)

	app := setupApp(cfg)
	setupRoutes(app, db, cfg)

	log.Fatal(app.Listen(":" + cfg.App.Port))
}

func setupApp(cfg *config.Config) *fiber.App {
	app := fiber.New()
	app.Use(recover.New())

	if cfg.App.Env == "local" {
		app.Use(middleware.RequestLogger())
	}

	return app
}

func setupRoutes(app *fiber.App, db *gorm.DB, cfg *config.Config) {
	setupHealthRoute(app)
	setupUserRoutes(app, db, cfg)
	setupIncidentRoutes(app, db, cfg)
}

// ─── Health ───────────────────────────────────────────────────────────────────
func setupHealthRoute(app *fiber.App) {
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})
}

// ─── Users ────────────────────────────────────────────────────────────────────
func setupUserRoutes(app *fiber.App, db *gorm.DB, cfg *config.Config) {
	repo := repository.NewUserRepository(db)
	svc := service.NewUserService(repo)
	h := handler.NewUserHandler(svc)
	h.RegisterRoutes(app, cfg.App.APIKey)
}

// ─── Incidents ────────────────────────────────────────────────────────────────
func setupIncidentRoutes(app *fiber.App, db *gorm.DB, cfg *config.Config) {
	repo := repository.NewIncidentRepository(db)
	svc := service.NewIncidentService(repo)
	h := handler.NewIncidentHandler(svc)
	h.RegisterRoutes(app, cfg.App.APIKey)
}
