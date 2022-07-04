package service

type SaGame interface {
	CreateMember() error
	GetMemberBalance() error
	Deposit() error
	Withdraw() error
}
