package mbank

import (
	"context"

	"github.com/google/uuid"
	"github.com/jen6/bank_simulation/internal/application/port"
	"github.com/jen6/bank_simulation/internal/constants"
	"github.com/jen6/bank_simulation/internal/utils/optional"
)

type MockBank struct {
	balance *int64
}

func NewMockBank(balance int64) MockBank {
	return MockBank{balance: &balance}
}

func (bank MockBank) ListAccounts(ctx context.Context, cardNumber string, pinNumber string) ([]port.AccountInfo, error) {
	if !(cardNumber == "1234" && pinNumber == "1234") {
		return nil, constants.WrongBankAuthorizationInfo
	}

	return []port.AccountInfo{
		{AccountID: "test account1", SessionID: "test sessionid1"},
		{AccountID: "test account2", SessionID: "test sessionid2"},
	}, nil
}
func (bank MockBank) GetBalance(ctx context.Context, accountInfo port.AccountInfo) (int64, error) {
	return *bank.balance, nil
}

func (bank *MockBank) ApplyTransaction(ctx context.Context, accountInfo port.AccountInfo, transaction port.Transaction) (optional.Option[port.TransactionResult], error) {
	var err error = nil

	beforeBalance := *bank.balance
	afterBalance := *bank.balance
	if transaction.Type == constants.DepositeTransaction {
		afterBalance += int64(transaction.Amount)
	} else if transaction.Type == constants.WithdrawTransaction {
		afterBalance -= int64(transaction.Amount)
	}

	if afterBalance < -100 {
		afterBalance = beforeBalance
		err = constants.NotEnoughAccountBalance
	}
	*bank.balance = afterBalance
	result := port.TransactionResult{
		ID:            uuid.NewString(),
		BeforeBalance: beforeBalance,
		AfterBalance:  afterBalance,
	}
	return optional.Some(result), err
}
