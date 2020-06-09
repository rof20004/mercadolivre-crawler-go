package main

import (
	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("[main]", err)
	}

	app := fiber.New()
	initRoutes(app)
	initServer(app)
}
