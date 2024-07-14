package usecase

import "github.com/jen6/bank_simulation/internal/application/constants"

type ShowBalanceCommand AccountInfo
type ShowBalanceResult struct {
	CurrentBalance int64
}

type Recipt struct {
	AccountID       string
	TransactionType constants.TransactionType
	Amount          uint64
	BeforeBalance   int64
	AfterBalance    int64
	Error           error
}

type TransactionCommand struct {
	Account         AccountInfo
	TransactionType constants.TransactionType
	Amount          uint64
}
