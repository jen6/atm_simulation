package persistence

import "github.com/jen6/bank_simulation/internal/domain"

type CashbinRepo struct {
	Cashbin *domain.CashBin
}

func NewCashbinRepo(balance uint64) CashbinRepo {
	cashbin := domain.NewCashBin("abcd", balance)
	return CashbinRepo{
		Cashbin: &cashbin,
	}

}

func (repo CashbinRepo) Get(id string) (domain.CashBin, error) {
	return *repo.Cashbin, nil
}

func (repo CashbinRepo) Update(cashbin domain.CashBin) error {
	repo.Cashbin = &cashbin
	return nil
}
