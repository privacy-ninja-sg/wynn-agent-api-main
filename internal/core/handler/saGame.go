package handler

import (
	"github.com/gofiber/fiber/v2"
)

type SaGameHandler interface {
	CreateMember(c *fiber.Ctx) error
	GetMemberBalance(c *fiber.Ctx) error
	Deposit(c *fiber.Ctx) error
	Withdraw(c *fiber.Ctx) error
}
