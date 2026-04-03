package main

import (
	"log"
	"pulseops/internal/config"
	"pulseops/internal/database"
	"pulseops/internal/handler"
	"pulseops/internal/repository"
	"pulseops/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gorm.io/gorm"
)

func main() {
	cfg := config.Load()
	db := database.Connect(cfg)

	app := setupApp()
	setupRoutes(app, db)

	log.Fatal(app.Listen(":" + cfg.AppPort))
}

func setupApp() *fiber.App {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())
	return app
}

func setupRoutes(app *fiber.App, db *gorm.DB) {
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	userHandler.RegisterRoutes(app)
}
