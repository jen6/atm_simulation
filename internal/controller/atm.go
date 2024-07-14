package controller

import (
	"context"
	"fmt"

	application "github.com/jen6/bank_simulation/internal/application/service"
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

func NewItem(
	id string,
	cashbinHardware port.Cashbin,
) ATM {
	return ATM{id: id, cashbinHardware: cashbinHardware}
}

func (atm ATM) Run() {
	ctx := context.Background()

	var cardNumber, pin, bankName string
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

}
