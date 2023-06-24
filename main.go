package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/kadirosmanust/URLshortener/routes"
)

func GetEnvWithKey(key string) string {
	return os.Getenv(key)
}

func main() {
	godotenv.Load()

	app := fiber.New()
	app.Use(cors.New())
	routes.ShortenerRouter(app)
	routes.ResolverRoute(app)
	log.Fatal(app.Listen(":8080"))
}
