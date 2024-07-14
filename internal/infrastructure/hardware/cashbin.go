package hardware

type MockCashbin struct {
	err error
}

func (mc MockCashbin) CashIn(amount uint64) error {
	return mc.err
}
func (mc MockCashbin) CashOut(amount uint64) error {
	return mc.err
}
