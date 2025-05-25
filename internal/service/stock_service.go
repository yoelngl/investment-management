package service

import (
	"fmt"
	"managementInvestment/models"
	"managementInvestment/utils"
)

// StockService manages stock data and prices
type StockService struct {
	Stocks []models.Stock
}

// NewStockService creates a new StockService instance
func NewStockService(stocks []models.Stock) *StockService {
	return &StockService{
		Stocks: stocks,
	}
}

// GetPrice returns the price of a stock by symbol, or 0 if not found
func (s *StockService) GetPrice(symbol string) int {
	idx := utils.BinarySearchStock(symbol, s.Stocks)
	if idx == -1 {
		return 0
	}
	return s.Stocks[idx].Price
}

// UpdatePrice updates the price of a stock by symbol
// Returns true if updated, false if stock not found
func (s *StockService) UpdatePrice(symbol string, newPrice int) bool {
	fmt.Printf("Updating price for %s\n", symbol)
	idx := utils.BinarySearchStock(symbol, s.Stocks)
	if idx == -1 {
		return false
	}
	s.Stocks[idx].Price = newPrice
	return true
}
