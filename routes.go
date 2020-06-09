package main

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber"
)

type Search struct {
	Text  string `json:"search"`
	Limit int    `json:"limit"`
}

func (s Search) isValid() bool {
	if strings.TrimSpace(s.Text) == "" {
		return false
	}

	return true
}

func initRoutes(app *fiber.App) {
	app.Post("/api/v1/products", func(ctx *fiber.Ctx) {
		search := &Search{
			Text:  "",
			Limit: 5,
		}

		if err := ctx.BodyParser(search); err != nil {
			ctx.Status(http.StatusBadRequest).SendString(err.Error())
			return
		}

		if !search.isValid() {
			ctx.Status(http.StatusBadRequest).SendString("Please, inform a text to search")
			return
		}

		products := searchProduct(search)

		ctx.JSON(products)
	})
}
