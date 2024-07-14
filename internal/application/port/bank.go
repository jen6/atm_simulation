package port

import (
	"context"

	"github.com/jen6/bank_simulation/internal/application/constants"
)

type AccountInfo struct {
	AccountID string // Bank Account ID
	SessionID string // Authorization key for each banking system
}

type Transaction struct {
	ID     string
	Type   constants.TransactionType
	Amount uint64
}

type TransactionResult struct {
	ID            string
	BeforeBalance int64
	AfterBalance  int64
}

type Bank interface {
	ListAccounts(ctx context.Context, cardNumber string, pinNumber string) ([]AccountInfo, error)
	GetBalance(ctx context.Context, accountInfo AccountInfo) (int64, error)
	BeginTransaction(ctx context.Context, accountInfo AccountInfo) (string, error)
	ApplyTransaction(ctx context.Context, accountInfo AccountInfo, transaction Transaction) (TransactionResult, error)
}
