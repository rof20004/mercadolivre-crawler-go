package main

import "github.com/gofiber/fiber"

func initServer(app *fiber.App) {
	app.Listen(3000)
}