package service

import "wynn-agent-api/internal/core/service"

type prettyGame struct {
	apiKey    string
	apiSecret string
}

func NewPrettyGame(apiKey, apiSecret string) service.PrettyGame {
	return &prettyGame{apiKey: apiKey, apiSecret: apiSecret}
}

func (p prettyGame) CreateMember() error {
	return nil
}

func (p prettyGame) GetMemberBalance() error {
	return nil
}

func (p prettyGame) Deposit() error {
	return nil
}

func (p prettyGame) Withdraw() error {
	return nil
}
