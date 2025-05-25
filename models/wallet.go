package models

import "errors"

type CustomerWallet struct {
	Balance int
}

func (w *CustomerWallet) HasSufficientBalance(amount int) bool {
	return w.Balance >= amount
}

func (w *CustomerWallet) Withdraw(amount int) error {
	if amount < 0 {
		return errors.New("withdraw amount must be positive")
	}
	if w.Balance < amount {
		return errors.New("insufficient balance")
	}
	w.Balance -= amount
	return nil
}

func (w *CustomerWallet) Deposit(amount int) error {
	if amount < 0 {
		return errors.New("deposit amount must be positive")
	}
	w.Balance += amount
	return nil
}
