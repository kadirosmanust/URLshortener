package controllers

import (
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kadirosmanust/URLshortener/db"
	"github.com/kadirosmanust/URLshortener/utils"
)

type request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"short"`
	Expiry      time.Duration `json:"expiry"`
}

type response struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"short"`
	Expiry      time.Duration `json:"expiry"`
}

func ShortenerController(ctx *fiber.Ctx) error {

	body := &request{}

	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse bodyJson",
		})
	}

	if !utils.IsURL(body.URL) {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "InvalidURL",
		})
	}

	// enforce HTTPS, SSL
	body.URL = utils.EnforceHTTP(body.URL)

	var id string

	if body.CustomShort == "" {
		id = utils.Base62Encode(rand.Uint64())
	} else {
		id = body.CustomShort
	}

	r := db.CreateRedisClient(0)
	defer r.Close()

	val, _ := r.Get(db.Ctx, id).Result()

	if val != "" {

		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Custom short url already in use",
		})
	}

	if body.Expiry == 0 {
		body.Expiry = 24
	}

	err := r.Set(db.Ctx, id, body.URL, body.Expiry*3600*time.Second).Err()

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to connect to server",
		})
	}

	resp := response{
		URL:         body.URL,
		CustomShort: "",
		Expiry:      body.Expiry,
	}

	resp.CustomShort = os.Getenv("DOMAIN") + "/" + id

	return ctx.Status(fiber.StatusOK).JSON(resp)

}
