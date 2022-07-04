package handler

import (
	"wynn-agent-api/internal/core/handler"
	"wynn-agent-api/internal/core/service"

	"github.com/gofiber/fiber/v2"
)

type saGameHandler struct {
	saGameService service.SaGame
}

func NewSaGameHandler(saGameService service.SaGame) handler.SaGameHandler {
	return &saGameHandler{saGameService: saGameService}
}

func (s saGameHandler) CreateMember(c *fiber.Ctx) error {
	return nil
}

func (s saGameHandler) GetMemberBalance(c *fiber.Ctx) error {
	return nil
}

func (s saGameHandler) Deposit(c *fiber.Ctx) error {
	return nil
}

func (s saGameHandler) Withdraw(c *fiber.Ctx) error {
	return nil
}
