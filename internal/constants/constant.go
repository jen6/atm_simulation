package constants

type TransactionType string

const (
	DepositeTransaction TransactionType = "deposite"
	WithdrawTransaction TransactionType = "withdraw"
)

type TransactionStatus string

const (
	InprogressTransaction TransactionStatus = "inprogress"
	CancledTransaction    TransactionStatus = "cancled"
	CompletedTransaction  TransactionStatus = "completed"
)
