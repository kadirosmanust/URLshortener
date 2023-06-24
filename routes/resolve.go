package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kadirosmanust/URLshortener/controllers"
)

func ResolverRoute(app fiber.Router) {
	app.Get("/:url", controllers.ResolveController)
}
