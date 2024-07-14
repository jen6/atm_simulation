package domain

import (
	"time"

	"github.com/jen6/bank_simulation/internal/constants"
	"github.com/jen6/bank_simulation/internal/utils/optional"
)

type CashBin struct {
	ID                 string
	balance            uint64
	currentTransaction optional.Option[Transaction]
}

func NewCashBin(id string, balance uint64) CashBin {
	return CashBin{
		ID:      id,
		balance: balance,
	}
}

func (atm *CashBin) Deposite(amount uint64, requestTime time.Time) {
	atm.balance += amount
	transaction := NewTransaction(
		constants.DepositeTransaction,
		atm.balance,
		amount,
		requestTime,
	)

	atm.currentTransaction = optional.Some(transaction)
}

func (atm *CashBin) Withdraw(amount uint64, requestTime time.Time) error {
	if amount > atm.balance {
		return constants.NotEnoughCash
	}
	atm.balance -= amount

	transaction := NewTransaction(
		constants.WithdrawTransaction,
		atm.balance,
		amount,
		requestTime,
	)

	atm.currentTransaction = optional.Some(transaction)
	return nil
}

func (atm *CashBin) Rollback() {
	transaction, err := atm.currentTransaction.Take()
	if err != nil {
		return
	}

	atm.balance = transaction.BeforeBalance
}
