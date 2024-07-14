package constants

import "errors"

var (
	NotEnoughAccountBalance    error = errors.New("NotEnoughAccountBalance")
	NotEnoughCash              error = errors.New("NotEnoughCash")
	WrongBankAuthorizationInfo error = errors.New("WrongBankAuthorizationInfo")
	ATMNotFound                error = errors.New("ATMNotFound")
	BankNotFound               error = errors.New("BankNotFound")
)
