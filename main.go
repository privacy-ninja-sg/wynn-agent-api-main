package main

import (
	"os"
	"wynn-agent-api/pkg/configs"
	"wynn-agent-api/pkg/routes"
	"wynn-agent-api/pkg/utils"

	"entgo.io/ent/examples/fs/ent"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func init() {
	_ = godotenv.Load(".env.example")
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	// ========== Initial fiber app ==========
	app := fiber.New(configs.FiberConfig())

	// init router for routing application server...
	appEnv := os.Getenv("APP_ENV")
	router := routes.NewRouter(app, &ent.Client{}, configs.JwtConfig(), appEnv)
	// run on develop mode
	// router.Swagger()
	router.Monitor()
	router.Logging()
	// run on all mode
	router.Main()

	// setup 404
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("Sorry can't find that!")
	})

	// start application server with graceful shutdown...
	utils.StartServerWithGracefulShutdown(app)
}
