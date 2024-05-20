package config

import (
	"github.com/gofiber/fiber/v2"
	errorhandler "system_employee/error_handler"
)

func NewFiberConfig() fiber.Config {

	return fiber.Config{
		ErrorHandler: errorhandler.ErrorHandler,
	}
}
