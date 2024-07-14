package port

type Cashbin interface {
	CashIn(amount uint64) error
	CashOut(amount uint64) error
}
