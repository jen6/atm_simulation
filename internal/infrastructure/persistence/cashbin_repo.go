package persistence

import "github.com/jen6/bank_simulation/internal/domain"

type CashbinRepo struct {
	cashbin *domain.CashBin
}

func NewCashbinRepo(balance uint64) CashbinRepo {
	cashbin := domain.NewCashBin("abcd", 200)
	return CashbinRepo{
		cashbin: &cashbin,
	}

}

func (repo CashbinRepo) Get(id string) (domain.CashBin, error) {
	return *repo.cashbin, nil
}

func (repo CashbinRepo) Update(cashbin domain.CashBin) error {
	repo.cashbin = &cashbin
	return nil
}
