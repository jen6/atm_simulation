package port

import (
	"context"

	"github.com/jen6/bank_simulation/internal/application/constants"
	"github.com/jen6/bank_simulation/internal/utils/optional"
)

type AccountInfo struct {
	AccountID string // Bank Account ID
	SessionID string // Authorization key for each banking system
}

type Transaction struct {
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
}

type AccountRepository interface {
	GetBalance(ctx context.Context, accountInfo AccountInfo) (int64, error)
	ApplyTransaction(ctx context.Context, accountInfo AccountInfo, transaction Transaction) (optional.Option[TransactionResult], error)
}
