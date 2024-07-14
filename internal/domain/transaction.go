package domain

import (
	"errors"
	"time"

	"github.com/jen6/bank_simulation/internal/constants"
)

type Transaction struct {
	Type          constants.TransactionType
	BeforeBalance uint64
	Amount        uint64
	AfterBalance  uint64
	Status        constants.TransactionStatus
	CreatedAt     time.Time
	ExpiredAt     time.Time
}

func NewTransaction(transactionType constants.TransactionType, currentBalance uint64, amount uint64, requestTime time.Time) Transaction {
	afterBalance := currentBalance
	if transactionType == constants.DepositeTransaction {
		afterBalance += amount
	} else if transactionType == constants.WithdrawTransaction {
		afterBalance -= amount
	}

	return Transaction{
		Type:          transactionType,
		BeforeBalance: currentBalance,
		Amount:        amount,
		AfterBalance:  afterBalance,
		Status:        constants.InprogressTransaction,
		CreatedAt:     requestTime,
		ExpiredAt:     requestTime.Add(10 * time.Minute),
	}
}

func (t Transaction) IsValid(now time.Time) error {
	if now.After(t.ExpiredAt) {
		return errors.New("Transaction already expired")
	}

	if t.Type == constants.WithdrawTransaction && t.Amount < t.BeforeBalance {
		return constants.NotEnoughCash
	}

	return nil
}
