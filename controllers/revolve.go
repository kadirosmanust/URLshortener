package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kadirosmanust/URLshortener/db"
	"github.com/redis/go-redis/v9"
)

func ResolveController(ctx *fiber.Ctx) error {
	url := ctx.Params("url")

	r := db.CreateRedisClient(0)
	defer r.Close()

	value, err := r.Get(db.Ctx, url).Result()
	if err == redis.Nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "shortUrl not found"})
	} else if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal error"})
	}

	return ctx.Redirect(value, 301)
}
