package persistence

import "github.com/jen6/bank_simulation/internal/domain"

type CashbinRepo struct {
	cashbin *domain.CashBin
}

func (repo CashbinRepo) Get(id string) (domain.CashBin, error) {
	return *repo.cashbin, nil
}

func (repo CashbinRepo) Update(cashbin domain.CashBin) error {
	repo.cashbin = &cashbin
	return nil
}
