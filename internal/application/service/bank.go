package application

import (
	"context"
	"time"

	"github.com/jen6/bank_simulation/internal/application/port"
	"github.com/jen6/bank_simulation/internal/application/usecase"
	"github.com/jen6/bank_simulation/internal/utils/optional"
)

type BankService struct {
	bank port.Bank
}

func NewBank(bank port.Bank) BankService {
	return BankService{bank: bank}
}

func (bs BankService) ListAccount(ctx context.Context, command usecase.ListAccountCommand) (optional.Option[usecase.ListAccountResult], error) {

	timeOutCtx, cancelFunc := context.WithTimeout(ctx, 30*time.Second)
	defer cancelFunc()

	accounts, err := bs.bank.ListAccounts(timeOutCtx, command.CreditCardNumber, command.PinNumber)
	if err != nil {
		return optional.None[usecase.ListAccountResult](), err
	}

	var accInfos []usecase.AccountInfo = make([]usecase.AccountInfo, len(accounts))
	for i, account := range accounts {
		accInfos[i] = usecase.AccountInfo{
			AccountID: account.AccountID,
			SessionID: account.SessionID,
		}
	}
	result := usecase.ListAccountResult{
		Accounts: accInfos,
	}
	return optional.Some(result), nil

}
