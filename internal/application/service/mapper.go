package application

import (
	"github.com/jen6/bank_simulation/internal/application/port"
	"github.com/jen6/bank_simulation/internal/application/usecase"
)

func mapAccountInfo(uac usecase.AccountInfo) port.AccountInfo {
	return port.AccountInfo{
		AccountID: uac.AccountID,
		SessionID: uac.SessionID,
	}
}
