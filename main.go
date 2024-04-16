package main

import (
	"github.com/Hanivan/mailer-service/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/send-email", func(c *fiber.Ctx) error {
		var params utils.EmailParams

		if err := c.BodyParser(&params); err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your request", "data": err})
		}

		if params.PostmarkToken == "" {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "\"postmark_token\" and \"account_token\" cannot be empty", "data": nil})
		}

		response, err := utils.SendToGmail(params.PostmarkToken, params.PostmarkToken, params)

		if err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't send email", "data": err})
		}

		return c.JSON(fiber.Map{"status": "success", "message": "Email sent", "data": response})
	})

	app.Listen(":8088")
}
