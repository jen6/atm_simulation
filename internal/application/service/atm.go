package application

import (
	"context"
	"time"

	"github.com/jen6/bank_simulation/internal/application/port"
	"github.com/jen6/bank_simulation/internal/application/usecase"
	"github.com/jen6/bank_simulation/internal/constants"
	"github.com/jen6/bank_simulation/internal/domain"
)

type ATMService struct {
	accountRepository  port.AccountRepository
	cashbinRepository  domain.CashbinRepository
	timestampGenerator TimestampGenerator
}

func (as ATMService) ShowBalance(ctx context.Context, command usecase.ShowBalanceCommand) (usecase.ShowBalanceResult, error) {
	timeoutCtx, cancelFunc := context.WithTimeout(ctx, 30*time.Second)
	defer cancelFunc()

	accountInfo := mapAccountInfo(command.Account)
	balance, err := as.accountRepository.GetBalance(timeoutCtx, accountInfo)
	return usecase.ShowBalanceResult{CurrentBalance: balance}, err
}

func (as ATMService) DoTransaction(ctx context.Context, command usecase.TransactionCommand) (usecase.TransactionCommandResult, error) {
	timeoutCtx, cancelFunc := context.WithTimeout(ctx, 30*time.Second)
	defer cancelFunc()

	cashbin, err := as.cashbinRepository.Get(command.AtmId)
	if err != nil {
		return usecase.TransactionCommandResult{}, err
	}

	if command.TransactionType == constants.DepositeTransaction {
		return as.deposite(timeoutCtx, cashbin, command)
	}
	return as.withdraw(timeoutCtx, cashbin, command)
}

func (as ATMService) withdraw(
	ctx context.Context,
	cashbin domain.CashBin,
	command usecase.TransactionCommand,
) (result usecase.TransactionCommandResult, err error) {
	err = cashbin.Withdraw(command.Amount, as.timestampGenerator.Now())
	if err != nil {
		return result, err
	}

	err = as.cashbinRepository.Update(cashbin)
	if err != nil {
		return result, err
	}
	accountInfo := mapAccountInfo(command.Account)
	transactionReq := port.Transaction{
		Type:   constants.WithdrawTransaction,
		Amount: command.Amount,
	}

	transactionResult, err := as.accountRepository.ApplyTransaction(ctx, accountInfo, transactionReq)
	transaction := transactionResult.UnwrapOr(port.TransactionResult{})

	result.BankRecipt = usecase.Recipt{
		AccountID:       command.Account.AccountID,
		TransactionType: constants.WithdrawTransaction,
		Amount:          command.Amount,
		BeforeBalance:   transaction.BeforeBalance,
		AfterBalance:    transaction.AfterBalance,
	}
	result.WithdrawAmount = command.Amount

	if err != nil {
		result.WithdrawAmount = 0

		cashbin.Rollback()
		err = as.cashbinRepository.Update(cashbin)
	}
	return result, err
}

func (as ATMService) deposite(
	ctx context.Context,
	cashbin domain.CashBin,
	command usecase.TransactionCommand,
) (result usecase.TransactionCommandResult, err error) {
	cashbin.Deposite(command.Amount, as.timestampGenerator.Now())
	err = as.cashbinRepository.Update(cashbin)
	if err != nil {
		return result, err
	}

	accountInfo := mapAccountInfo(command.Account)
	transactionReq := port.Transaction{
		Type:   command.TransactionType,
		Amount: command.Amount,
	}

	transactionResult, err := as.accountRepository.ApplyTransaction(ctx, accountInfo, transactionReq)
	transaction := transactionResult.UnwrapOr(port.TransactionResult{})

	result.BankRecipt = usecase.Recipt{
		AccountID:       command.Account.AccountID,
		TransactionType: command.TransactionType,
		Amount:          command.Amount,
		BeforeBalance:   transaction.BeforeBalance,
		AfterBalance:    transaction.AfterBalance,
	}
	result.WithdrawAmount = 0

	if err != nil {
		cashbin.Rollback()
		err = as.cashbinRepository.Update(cashbin)
	}

	return result, err
}
