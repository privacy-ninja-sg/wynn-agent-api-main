package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"net/http"
)

func BasicAuth() fiber.Handler {
	return basicauth.New(basicauth.Config{
		Users: map[string]string{
			"admin": "Z2Vlay1wd2QtYWNjZXNz",
		},
		Realm: "Forbidden",
		Authorizer: func(user, pass string) bool {
			if user == "admin" && pass == "Z2Vlay1wd2QtYWNjZXNz" {
				return true
			}
			return false
		},
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"err": true,
				"msg": http.StatusText(fiber.StatusUnauthorized),
			})
		},
		ContextUsername: "_user",
		ContextPassword: "_pass",
	})
}
