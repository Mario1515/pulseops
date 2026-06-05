package middleware

import (
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func RequestLogger() fiber.Handler {
	file, err := os.OpenFile("logs/app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic("could not open log file: " + err.Error())
	}

	return logger.New(logger.Config{
		Format: "${time} | ${status} | ${latency} | ${method} | ${path} | ${error}\n",
		Stream: file,
	})
}
