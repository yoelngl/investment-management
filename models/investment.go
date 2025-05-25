package models

type Investment struct {
	Symbol   string
	Quantity int
	BuyPrice int
}

func (i *Investment) Value(currentPrice int) int {
	return i.Quantity * currentPrice
}

func (i *Investment) ProfitLoss(currentPrice int) int {
	currentValue := i.Value(currentPrice)
	buyValue := i.Quantity * i.BuyPrice
	return currentValue - buyValue
}
