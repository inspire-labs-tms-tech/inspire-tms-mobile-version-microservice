package routes

import "github.com/gofiber/fiber/v2"

func GetPing(c *fiber.Ctx) error {
	return c.SendString("pong")
}
