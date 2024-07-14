package usecase

import "github.com/jen6/bank_simulation/internal/constants"

type ShowBalanceCommand struct {
	Account AccountInfo
}
type ShowBalanceResult struct {
	CurrentBalance int64
}

type Recipt struct {
	AccountID       string
	TransactionType constants.TransactionType
	Amount          uint64
	BeforeBalance   int64
	AfterBalance    int64
}

type TransactionCommand struct {
	AtmId           string
	Account         AccountInfo
	TransactionType constants.TransactionType
	Amount          uint64
}

type TransactionCommandResult struct {
	BankRecipt     Recipt
	WithdrawAmount uint64
}
