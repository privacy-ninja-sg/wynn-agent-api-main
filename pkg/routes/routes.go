package routes

import (
	"os"
	"wynn-agent-api/pkg/middleware"

	"entgo.io/ent/examples/fs/ent"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

type Router struct {
	app       *fiber.App
	entCli    *ent.Client
	jwtSecret string
	appEnv    string
}

func NewRouter(app *fiber.App, entCli *ent.Client, jwtSecret string, appEnv string) *Router {
	return &Router{app: app, entCli: entCli, jwtSecret: jwtSecret, appEnv: appEnv}
}

// // Swagger func for describe group of API Docs routes.
// func (r *Router) Swagger() {
// 	if r.appEnv == "develop" {
// 		// Create routes group.
// 		route := r.app.Group("/swagger")
// 		// Route for GET method:
// 		route.Get("*", middleware.BasicAuth(), swagger.Handler) // get one user by ID
// 	}
// }

func (r *Router) Monitor() {
	if r.appEnv == "develop" {
		r.app.Get("/monitor", middleware.BasicAuth(), monitor.New())
	}
}

func (r *Router) Logging() {
	if r.appEnv == "develop" {
		r.app.Use(logger.New(logger.Config{
			Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
		}))
	}
}

func (r *Router) Main() {
	appKey := os.Getenv("APP_KEY")
	appSecret := os.Getenv("APP_SECRET")
	// new service
	// shortSmsService := service.NewShortSms(otpAppKey, otpAppSecret)

	// // new handler
	// hand := handler.NewOTPHandler(shortSmsService)

	// un-auth endpoint
	r.app.Get("", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"s":   "ok",
			"msg": "hello world",
		})
	})

	r.app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			appKey: appSecret,
		},
	}))

	// r.app.Post("/api/request", hand.Request)
	// r.app.Post("/api/verify", hand.Verify)
}
