package main

import (
	"fmt"
	"os"

	"github.com/jen6/bank_simulation/internal/controller"
	bankapi "github.com/jen6/bank_simulation/internal/infrastructure/bank_api"
	mbank "github.com/jen6/bank_simulation/internal/infrastructure/bank_api/mock"
	"github.com/jen6/bank_simulation/internal/infrastructure/hardware"
	"github.com/jen6/bank_simulation/internal/infrastructure/persistence"
	"github.com/jen6/bank_simulation/internal/infrastructure/timestamp"
)

func main() {
	atmID := "abcd"

	cashbinRepo := persistence.NewCashbinRepo(200)
	mockBank := mbank.NewMockBank(100, nil, nil)
	banks := bankapi.NewCompositeBank(mockBank)

	atm := controller.NewATM(
		atmID,
		banks,
		hardware.MockCashbin{},
		cashbinRepo,
		timestamp.TimestampGenerator{},
		os.Stdin,
	)
	for {
		var exit string

		atm.Run()
		fmt.Print("exit? y/n")
		fmt.Scan(&exit)

		if exit == "y" {
			fmt.Println("bye")
			return
		}
	}

}
