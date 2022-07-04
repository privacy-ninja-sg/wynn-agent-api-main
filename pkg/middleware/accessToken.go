package middleware

import (
	"net/http"
	"time"
	"wynn-agent-api/internal/core/model"
	"wynn-agent-api/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func AcccessToken(jwtSecret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get now time.
		now := time.Now().Unix()

		// Get claims from JWT.
		claims, err := utils.ExtractTokenMetadata(c, jwtSecret)
		if err != nil {
			// Return status 500 and JWT parse error.
			return c.Status(http.StatusUnauthorized).JSON(model.DefaultResponse{
				Status: "error",
				Code:   http.StatusUnauthorized,
				ErrMsg: err.Error(),
			})
		}

		// Set expiration time from JWT data of current book.
		expires := claims.Expires

		// Checking, if now time greather than expiration from JWT.
		if now > expires {
			// Return status 401 and unauthorized error message.
			return c.Status(fiber.StatusUnauthorized).JSON(model.DefaultResponse{
				Status: "error",
				Code:   http.StatusUnauthorized,
				ErrMsg: "unauthorized, check expiration time of your token",
			})
		}

		// set tokenPayload for handler
		c.Locals("tokenPayload", claims)

		return c.Next()
	}
}
