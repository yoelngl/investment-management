package service

import (
	"fmt"
	"managementInvestment/models"
	"managementInvestment/utils"
)

// InvestmentService manages investment operations
type InvestmentService struct {
	Investments    []models.Investment
	Count          int
	MaxInvestments int
	StockService   *StockService
	Wallet         *models.CustomerWallet
}

// NewInvestmentService creates new InvestmentService
func NewInvestmentService(maxInvestments int, stockService *StockService, wallet *models.CustomerWallet) *InvestmentService {
	return &InvestmentService{
		Investments:    make([]models.Investment, 0, maxInvestments),
		Count:          0,
		MaxInvestments: maxInvestments,
		StockService:   stockService,
		Wallet:         wallet,
	}
}

// GetInvestments returns current investments slice
func (s *InvestmentService) GetInvestments() []models.Investment {
	return s.Investments
}

// GetCount returns current number of investments
func (s *InvestmentService) GetCount() int {
	return s.Count
}

// GetWallet returns the associated wallet
func (s *InvestmentService) GetWallet() *models.CustomerWallet {
	return s.Wallet
}

// BuyStock buys quantity of stock with symbol, updates wallet and investments
func (s *InvestmentService) BuyStock(symbol string, quantity int) error {
	stockPrice := s.StockService.GetPrice(symbol)
	if stockPrice == 0 {
		return fmt.Errorf("stock not found")
	}

	totalPrice := stockPrice * quantity
	if !s.Wallet.HasSufficientBalance(totalPrice) {
		return fmt.Errorf("insufficient balance")
	}

	invIdx := utils.SequentialSearch(symbol, s.Investments)
	if invIdx == -1 {
		if s.Count >= s.MaxInvestments {
			return fmt.Errorf("investment list full")
		}
		newInvestment := models.Investment{Symbol: symbol, Quantity: quantity, BuyPrice: stockPrice}
		s.Investments = append(s.Investments, newInvestment)
		s.Count++
	} else {
		totalQty := s.Investments[invIdx].Quantity + quantity
		avgPrice := ((s.Investments[invIdx].BuyPrice * s.Investments[invIdx].Quantity) + (stockPrice * quantity)) / totalQty
		s.Investments[invIdx].Quantity = totalQty
		s.Investments[invIdx].BuyPrice = avgPrice
	}

	// Withdraw balance from wallet, cek error
	if err := s.Wallet.Withdraw(totalPrice); err != nil {
		return err
	}

	return nil
}

// SellStock sells quantity of stock with symbol, updates wallet and investments
func (s *InvestmentService) SellStock(symbol string, quantity int) error {
	invIdx := utils.SequentialSearch(symbol, s.Investments)
	if invIdx == -1 {
		return fmt.Errorf("investment not found")
	}

	if quantity > s.Investments[invIdx].Quantity {
		return fmt.Errorf("insufficient shares")
	}

	currentPrice := s.StockService.GetPrice(symbol)
	totalSellPrice := currentPrice * quantity

	// Deposit to wallet, cek error
	if err := s.Wallet.Deposit(totalSellPrice); err != nil {
		return err
	}

	s.Investments[invIdx].Quantity -= quantity

	if s.Investments[invIdx].Quantity == 0 {
		// Remove investment by shifting slice
		s.Investments = append(s.Investments[:invIdx], s.Investments[invIdx+1:]...)
		s.Count--
	}

	return nil
}

// SortInvestments sorts investments by their current value ascending or descending, using selection or insertion sort
func (s *InvestmentService) SortInvestments(ascending bool, useSelection bool) {
	if useSelection {
		utils.SelectionSort(s.Investments, s.Count, s.StockService.GetPrice, ascending)
	} else {
		utils.InsertionSort(s.Investments, s.Count, s.StockService.GetPrice, ascending)
	}
}
