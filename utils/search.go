package utils

import (
	"managementInvestment/models"
)

func SequentialSearch(symbol string, investments []models.Investment) int {
	for i, investment := range investments {
		if investment.Symbol == symbol {
			return i
		}
	}
	return -1
}

func BinarySearchStock(symbol string, stockLists []models.Stock) int {
	low := 0
	high := len(stockLists) - 1
	for low <= high {
		mid := (low + high) / 2
		if stockLists[mid].Symbol == symbol {
			return mid
		} else if stockLists[mid].Symbol < symbol {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func BinarySearchInvestments(symbol string, investments []models.Investment) int {
	low := 0
	high := len(investments) - 1
	for low <= high {
		mid := (low + high) / 2
		if investments[mid].Symbol == symbol {
			return mid
		} else if investments[mid].Symbol < symbol {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
