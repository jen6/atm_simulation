package controller

import (
	"context"
	"fmt"

	application "github.com/jen6/bank_simulation/internal/application/service"
	"github.com/jen6/bank_simulation/internal/application/usecase"
	"github.com/jen6/bank_simulation/internal/constants"
	"github.com/jen6/bank_simulation/internal/controller/port"
	bankapi "github.com/jen6/bank_simulation/internal/infrastructure/bank_api"
	"github.com/jen6/bank_simulation/internal/infrastructure/persistence"
	"github.com/jen6/bank_simulation/internal/infrastructure/timestamp"
)

type ATM struct {
	id                 string
	banks              bankapi.CompositeBank
	timestampGenerator timestamp.TimestampGenerator
	cashbinHardware    port.Cashbin
	cashbinRepo        persistence.CashbinRepo
}

func NewATM(
	id string,
	banks bankapi.CompositeBank,
	cashbinHardware port.Cashbin,
	cashbinRepo persistence.CashbinRepo,
	timestampGenerator timestamp.TimestampGenerator,
) ATM {
	return ATM{
		id:                 id,
		banks:              banks,
		cashbinHardware:    cashbinHardware,
		cashbinRepo:        cashbinRepo,
		timestampGenerator: timestampGenerator,
	}
}

func (atm ATM) Run() {
	ctx := context.Background()

	var cardNumber, pin, bankName string
	fmt.Println("-----------------------------")
	fmt.Print("Input Card Number : ")
	fmt.Scan(&cardNumber)

	fmt.Print("Input Pin Number : ")
	fmt.Scan(&pin)

	fmt.Print(atm.banks.ListBank(), " Choose bank : ")
	fmt.Scan(&bankName)

	bank, accountRepo, err := atm.banks.FindBank(bankName)
	if err != nil {
		fmt.Println(err)
		return
	}

	bankService := application.NewBank(bank)
	atmService := application.NewATMService(accountRepo, atm.cashbinRepo, atm.timestampGenerator)

	accountsResult, err := bankService.ListAccount(ctx, usecase.ListAccountCommand{
		CreditCardNumber: cardNumber,
		PinNumber:        pin,
	})
	accounts := accountsResult.UnwrapOr(usecase.ListAccountResult{Accounts: nil})
	if err != nil {
		fmt.Println(err)
		return
	}

	if accounts.Accounts == nil {
		fmt.Println("No Accounts Exists")
		return
	}

	accountIdx := -1
	fmt.Println("Choose Account\n")
	for i, account := range accounts.Accounts {
		fmt.Println(i+1, " ", account.AccountID)
	}
	fmt.Scan(&accountIdx)

	if !(1 <= accountIdx && accountIdx <= len(accounts.Accounts)) {
		fmt.Println("There are no account ", accountIdx)
		return
	}

	selectedAccount := accounts.Accounts[accountIdx-1]

	if err = atm.ShowBalance(ctx, atmService, selectedAccount); err != nil {
		fmt.Println(err)
		return
	}

	if err = atm.chooseAction(ctx, atmService, selectedAccount); err != nil {
		fmt.Println(err)
		return
	}
}

func (atm ATM) ShowBalance(ctx context.Context, atmService application.ATMService, account usecase.AccountInfo) error {
	command := usecase.ShowBalanceCommand{
		Account: account,
	}
	result, err := atmService.ShowBalance(ctx, command)
	if err != nil {
		return err
	}
	fmt.Println("Current Balance is ", result.CurrentBalance)
	return nil
}

func (atm ATM) chooseAction(ctx context.Context, atmService application.ATMService, account usecase.AccountInfo) error {
	var option, amount uint64
	fmt.Println("Choose Action")
	fmt.Println("1. Deposite")
	fmt.Println("2. Withdraw")

	fmt.Scan(&option)

	if option == 1 {
		fmt.Println("Put Money : ")
		fmt.Scan(&amount)

		command := usecase.TransactionCommand{
			AtmId:           atm.id,
			Account:         account,
			TransactionType: constants.DepositeTransaction,
			Amount:          amount,
		}

		result, err := atmService.DoTransaction(ctx, command)
		if err != nil {
			return err
		}

		err = atm.cashbinHardware.CashIn(amount)
		if err != nil {
			return err
		}

		fmt.Println("Transaction Completed")
		fmt.Println("Current Balance is ", result.BankRecipt.AfterBalance)
		return nil
	} else if option == 2 {
		fmt.Println("How much money do you want?: ")
		fmt.Scan(&amount)

		command := usecase.TransactionCommand{
			AtmId:           atm.id,
			Account:         account,
			TransactionType: constants.WithdrawTransaction,
			Amount:          amount,
		}

		result, err := atmService.DoTransaction(ctx, command)
		if err != nil {
			return err
		}

		err = atm.cashbinHardware.CashOut(result.WithdrawAmount)
		if err != nil {
			return err
		}

		fmt.Println("Transaction Completed")
		fmt.Println("Current Balance is ", result.BankRecipt.AfterBalance)
		return nil
	}

	return nil
}
