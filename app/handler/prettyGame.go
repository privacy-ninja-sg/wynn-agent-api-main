package handler

import (
	"wynn-agent-api/internal/core/handler"
	"wynn-agent-api/internal/core/service"

	"github.com/gofiber/fiber/v2"
)

type prettyGameHandler struct {
	prettyGameService service.PrettyGame
}

func NewPrettyGameHandler(prettyGameService service.PrettyGame) handler.PrettyGameHandler {
	return &prettyGameHandler{prettyGameService: prettyGameService}
}

func (p prettyGameHandler) CreateMember(c *fiber.Ctx) error {
	return nil
}

func (p prettyGameHandler) GetMemberBalance(c *fiber.Ctx) error {
	return nil
}

func (p prettyGameHandler) Deposit(c *fiber.Ctx) error {
	return nil
}

func (p prettyGameHandler) Withdraw(c *fiber.Ctx) error {
	return nil
}
