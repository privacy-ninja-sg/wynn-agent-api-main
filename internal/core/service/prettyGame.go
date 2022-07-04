package service

type PrettyGame interface {
	CreateMember() error
	GetMemberBalance() error
	Deposit() error
	Withdraw() error
}
