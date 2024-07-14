package main

import (
	"fmt"
	"os"

	"github.com/jen6/bank_simulation/internal/controller"
	bankapi "github.com/jen6/bank_simulation/internal/infrastructure/bank_api"
	"github.com/jen6/bank_simulation/internal/infrastructure/hardware"
	"github.com/jen6/bank_simulation/internal/infrastructure/persistence"
	"github.com/jen6/bank_simulation/internal/infrastructure/timestamp"
)

func main() {
	atmID := "abcd"

	cashbinRepo := persistence.NewCashbinRepo(200)
	banks := bankapi.NewCompositeBank(100)

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
