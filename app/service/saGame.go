package service

import "wynn-agent-api/internal/core/service"

type saGame struct {
	apiKey    string
	apiSecret string
}

func NewSaGame(apiKey, apiSecret string) service.SaGame {
	return &saGame{apiKey: apiKey, apiSecret: apiSecret}
}

func (s saGame) CreateMember() error {
	return nil
}

func (s saGame) GetMemberBalance() error {
	return nil
}

func (s saGame) Deposit() error {
	return nil
}

func (s saGame) Withdraw() error {
	return nil
}
