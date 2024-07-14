package domain

type CashbinRepository interface {
	Get(id string) (CashBin, error)
	Update(cashbin CashBin) error
}
