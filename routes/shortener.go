package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kadirosmanust/URLshortener/controllers"
)

func ShortenerRouter(app fiber.Router) {
	app.Post("/short", controllers.ShortenerController)
}
