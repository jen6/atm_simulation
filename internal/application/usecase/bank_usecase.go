package usecase

type ListAccountCommand struct {
	CreditCardNumber string
	PinNumber        string
}

type AccountInfo struct {
	AccountID string // Bank Account ID
	SessionID string // Authorization key for each banking system
}

type ListAccountResult struct {
	Accounts []AccountInfo
}
