package constants

import "errors"

var (
	NotEnoughAccountBalance    error = errors.New("NotEnoughAccountBalance")
	NotEnoughCashBin           error = errors.New("NotEnoughCashBin")
	WrongBankAuthorizationInfo error = errors.New("WrongBankAuthorizationInfo")
)
