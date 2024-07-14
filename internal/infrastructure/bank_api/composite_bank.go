package bankapi

import (
	"github.com/jen6/bank_simulation/internal/application/port"
	"github.com/jen6/bank_simulation/internal/constants"
	mbank "github.com/jen6/bank_simulation/internal/infrastructure/bank_api/mock"
)

type CompositeBank struct {
	mockBank mbank.MockBank
}

func (cb CompositeBank) FindBank(bankName string) (port.Bank, port.AccountRepository, error) {
	if bankName == "mock_bank" {
		return &cb.mockBank, &cb.mockBank, nil
	}
	return nil, nil, constants.BankNotFound
}
