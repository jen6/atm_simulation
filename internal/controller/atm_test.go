package controller

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/jen6/bank_simulation/internal/constants"
	bankapi "github.com/jen6/bank_simulation/internal/infrastructure/bank_api"
	mbank "github.com/jen6/bank_simulation/internal/infrastructure/bank_api/mock"
	"github.com/jen6/bank_simulation/internal/infrastructure/hardware"
	"github.com/jen6/bank_simulation/internal/infrastructure/persistence"
	"github.com/jen6/bank_simulation/internal/infrastructure/timestamp"
	"github.com/stretchr/testify/assert"
)

const (
	CorrectCardNumber = "1234\n"
	CorrectPinNumber  = "1234\n"
	CorrectBankName   = "mock_bank\n"
	AccountSelect     = "1\n"

	WrongCardNumber = "wrong_Card\n"
	WrongPinNumber  = "wrong_Pin\n"
	WrongBankName   = "wrong_bank\n"

	BaseOutput = `-----------------------------
Input Card Number : Input Pin Number : [mock_bank] Choose bank : `

	AccountSelectOutput = `Choose Account
1   test account1
2   test account2
`
	ChooseActionOutput = `
Choose Action
1. Deposite
2. Withdraw
`
	depositeOutput = "Put Money : \n"
	withdrawOutput = "How much money do you want?: \n"
	completeOutput = "Transaction Completed\n"
)

func currentBalanceOutput(balance int64) string {
	return fmt.Sprint("Current Balance is ", balance)
}

func deposite(amount int64) string {
	return fmt.Sprint("1\n", amount, "\n")
}

func withdraw(amount int64) string {
	return fmt.Sprint("2\n", amount, "\n")
}

func TestATMMachine(t *testing.T) {
	scenarios := []struct {
		name                string
		givenCashbinBalance uint64
		givenAccountBalance int64
		whenInput           []string
		expectedOutput      []string
	}{
		{
			name:                "base case(deposite)",
			givenCashbinBalance: 200,
			givenAccountBalance: 0,
			whenInput: []string{
				CorrectCardNumber,
				CorrectPinNumber,
				CorrectBankName,
				AccountSelect,
				deposite(0),
			},
			expectedOutput: []string{
				BaseOutput,
				AccountSelectOutput,
				currentBalanceOutput(0),
				ChooseActionOutput,
				depositeOutput,
				completeOutput,
				currentBalanceOutput(0),
				"\n",
			},
		},
		{
			name:                "base case(withdraw)",
			givenCashbinBalance: 200,
			givenAccountBalance: 20,
			whenInput: []string{
				CorrectCardNumber,
				CorrectPinNumber,
				CorrectBankName,
				AccountSelect,
				withdraw(10),
			},
			expectedOutput: []string{
				BaseOutput,
				AccountSelectOutput,
				currentBalanceOutput(20),
				ChooseActionOutput,
				withdrawOutput,
				completeOutput,
				currentBalanceOutput(10),
				"\n",
			},
		},
		{
			name:                "wrong info(card)",
			givenCashbinBalance: 200,
			givenAccountBalance: 20,
			whenInput: []string{
				WrongCardNumber,
				CorrectPinNumber,
				CorrectBankName,
			},
			expectedOutput: []string{
				BaseOutput,
				constants.WrongBankAuthorizationInfo.Error(),
				"\n",
			},
		},
		{
			name:                "wrong info(pin)",
			givenCashbinBalance: 200,
			givenAccountBalance: 20,
			whenInput: []string{
				CorrectCardNumber,
				WrongPinNumber,
				CorrectBankName,
			},
			expectedOutput: []string{
				BaseOutput,
				constants.WrongBankAuthorizationInfo.Error(),
				"\n",
			},
		},
		{
			name:                "wrong info(bank)",
			givenCashbinBalance: 200,
			givenAccountBalance: 20,
			whenInput: []string{
				CorrectCardNumber,
				CorrectPinNumber,
				WrongBankName,
			},
			expectedOutput: []string{
				BaseOutput,
				constants.BankNotFound.Error(),
				"\n",
			},
		},
		{
			name:                "decline withdraw by bank",
			givenCashbinBalance: 200,
			givenAccountBalance: -100,
			whenInput: []string{
				CorrectCardNumber,
				CorrectPinNumber,
				CorrectBankName,
				AccountSelect,
				withdraw(10),
			},
			expectedOutput: []string{
				BaseOutput,
				AccountSelectOutput,
				currentBalanceOutput(-100),
				ChooseActionOutput,
				withdrawOutput,
				constants.NotEnoughAccountBalance.Error(),
				"\n",
			},
		},
		{
			name:                "decline withdraw by cashbin",
			givenCashbinBalance: 0,
			givenAccountBalance: 200,
			whenInput: []string{
				CorrectCardNumber,
				CorrectPinNumber,
				CorrectBankName,
				AccountSelect,
				withdraw(10),
			},
			expectedOutput: []string{
				BaseOutput,
				AccountSelectOutput,
				currentBalanceOutput(200),
				ChooseActionOutput,
				withdrawOutput,
				constants.NotEnoughCash.Error(),
				"\n",
			},
		},
	}

	for _, scenario := range scenarios {
		cashbinRepo := persistence.NewCashbinRepo(scenario.givenCashbinBalance)
		mockBank := mbank.NewMockBank(scenario.givenAccountBalance, nil, nil)
		banks := bankapi.NewCompositeBank(mockBank)
		fixedTime := time.Now()

		inputFile, _ := os.CreateTemp("", "")
		for _, in := range scenario.whenInput {
			io.WriteString(inputFile, in)
		}
		inputFile.Seek(0, io.SeekStart)

		atm := NewATM(
			"abcd",
			banks,
			hardware.MockCashbin{},
			cashbinRepo,
			timestamp.TimestampGenerator{FixedTimestamp: &fixedTime},
			inputFile,
		)

		t.Run(scenario.name, func(t *testing.T) {
			originStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w
			defer func() {
				os.Stdout = originStdout
			}()

			expected := strings.Builder{}
			for _, outputStr := range scenario.expectedOutput {
				expected.WriteString(outputStr)
			}

			atm.Run()

			w.Close()
			out, _ := io.ReadAll(r)
			realOut := string(out)

			assert.Equal(t, realOut, expected.String())

		})
	}
}
